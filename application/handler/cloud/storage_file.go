/*
   Nging is a toolbox for webmasters
   Copyright (C) 2018-present  Wenhui Shen <swh@admpub.com>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package cloud

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	minio "github.com/minio/minio-go"
	"github.com/webx-top/com"
	"github.com/webx-top/echo"

	"github.com/admpub/nging/application/dbschema"
	"github.com/admpub/nging/application/handler"
	"github.com/admpub/nging/application/library/config"
	"github.com/admpub/nging/application/library/filemanager"
	"github.com/admpub/nging/application/library/notice"
	"github.com/admpub/nging/application/library/s3manager"
	"github.com/admpub/nging/application/model"
)

func storageConnect(m *dbschema.CloudStorage) (client *minio.Client, err error) {
	isSecure := m.Secure == `Y`
	if len(m.Region) == 0 {
		client, err = minio.New(m.Endpoint, m.Key, m.Secret, isSecure)
	} else {
		client, err = minio.NewWithRegion(m.Endpoint, m.Key, m.Secret, isSecure, m.Region)
	}
	if err != nil {
		return client, err
	}
	if m.Secure != `Y` {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client.SetCustomTransport(tr)
	}
	return client, nil
}

func StorageFile(ctx echo.Context) error {
	ctx.Set(`activeURL`, `/cloud/storage`)
	id := ctx.Formx(`id`).Uint()
	m := model.NewCloudStorage(ctx)
	err := m.Get(nil, `id`, id)
	if err != nil {
		return err
	}
	client, err := storageConnect(m.CloudStorage)
	if err != nil {
		return err
	}

	mgr := s3manager.New(client, m.Bucket, config.DefaultConfig.Sys.EditableFileMaxBytes, ctx)

	ppath := ctx.Form(`path`)
	do := ctx.Form(`do`)
	var parentPath string
	if len(ppath) == 0 {
		ppath = `/`
	}
	if ppath != `/` {
		parentPath = strings.TrimSuffix(ppath, `/`)
		parentPath = path.Dir(parentPath)
	}
	if len(parentPath) > 0 && parentPath != `/` {
		parentPath += `/`
	}
	switch do {
	case `edit`:
		data := ctx.Data()
		if _, ok := config.DefaultConfig.Sys.Editable(ppath); !ok {
			data.SetInfo(ctx.T(`此文件不能在线编辑`), 0)
		} else {
			content := ctx.Form(`content`)
			encoding := ctx.Form(`encoding`)
			dat, err := mgr.Edit(ppath, content, encoding)
			if err != nil {
				data.SetInfo(err.Error(), 0)
			} else {
				if ctx.IsPost() {
					data.SetInfo(ctx.T(`保存成功`), 1)
				}
				data.SetData(dat, 1)
			}
		}
		return ctx.JSON(data)
	case `mkdir`:
		data := ctx.Data()
		newName := ctx.Form(`name`)
		if len(newName) == 0 {
			data.SetInfo(ctx.T(`请输入文件夹名`), 0)
		} else {
			err = mgr.Mkdir(ppath, newName)
			if err != nil {
				data.SetError(err)
			}
			if data.GetCode() == 1 {
				data.SetInfo(ctx.T(`创建成功`))
			}
		}
		return ctx.JSON(data)
	case `rename`:
		data := ctx.Data()
		newName := ctx.Form(`name`)
		err = mgr.Rename(ppath, newName)
		if err != nil {
			data.SetInfo(err.Error(), 0)
		} else {
			data.SetInfo(ctx.T(`重命名成功`), 1)
		}
		return ctx.JSON(data)
	case `search`:
		prefix := ctx.Form(`query`)
		num := ctx.Formx(`size`, `10`).Int()
		if num <= 0 {
			num = 10
		}
		paths := mgr.Search(ppath, prefix, num)
		data := ctx.Data().SetData(paths)
		return ctx.JSON(data)
	case `delete`:
		err = mgr.Remove(ppath)
		if err != nil {
			handler.SendFail(ctx, err.Error())
		}
		return ctx.Redirect(ctx.Referer())
	case `upload`:
		err = mgr.Upload(ppath)
		if err != nil {
			user := handler.User(ctx)
			if user != nil {
				notice.OpenMessage(user.Username, `upload`)
				notice.Send(user.Username, notice.NewMessageWithValue(`upload`, ctx.T(`文件上传出错`), err.Error()))
			}
			return ctx.JSON(echo.H{`error`: err.Error()}, 500)
		}
		return ctx.String(`OK`)
	case `download`:
		return mgr.Download(ppath)
	default:
		var dirs []os.FileInfo
		var exit bool
		err, exit, dirs = mgr.List(ppath)
		if exit {
			return err
		}
		ctx.Set(`dirs`, dirs)
	}
	ctx.Set(`parentPath`, parentPath)
	ctx.Set(`path`, ppath)
	var pathPrefix string
	pathPrefix = ppath
	pathSlice := strings.Split(strings.Trim(pathPrefix, `/`), `/`)
	pathLinks := make(echo.KVList, len(pathSlice))
	encodedSep := filemanager.EncodedSep
	urlPrefix := ctx.Request().URL().Path() + fmt.Sprintf(`?id=%d&path=`, id) + encodedSep
	for k, v := range pathSlice {
		urlPrefix += com.URLEncode(v)
		pathLinks[k] = &echo.KV{K: v, V: urlPrefix + `/`}
		urlPrefix += encodedSep
	}
	if !strings.HasSuffix(pathPrefix, `/`) {
		pathPrefix += `/`
	}
	ctx.Set(`pathLinks`, pathLinks)
	ctx.Set(`pathPrefix`, pathPrefix)
	ctx.SetFunc(`Editable`, func(fileName string) bool {
		_, ok := config.DefaultConfig.Sys.Editable(fileName)
		return ok
	})
	ctx.SetFunc(`Playable`, func(fileName string) string {
		mime, _ := config.DefaultConfig.Sys.Playable(fileName)
		return mime
	})
	ctx.Set(`data`, m.CloudStorage)
	return ctx.Render(`cloud/storage_file`, handler.Err(ctx, err))
}
