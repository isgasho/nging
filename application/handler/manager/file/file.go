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

// Package file 上传文件管理
package file

import (
	"time"

	"github.com/webx-top/echo"

	"github.com/admpub/nging/application/handler"
	"github.com/admpub/nging/application/model/file"
)

func FileList(ctx echo.Context) error {
	err := List(ctx, ``, 0)
	dialog := ctx.Formx(`dialog`).Bool()
	var suffix string
	if dialog {
		suffix = `_dialog`
	}

	multiple := ctx.Formx(`multiple`).Bool()
	ctx.Set(`dialog`, dialog)
	ctx.Set(`multiple`, multiple)
	return ctx.Render(`manager/file/list`+suffix, err)
}

func FileDelete(ctx echo.Context) (err error) {
	user := handler.User(ctx)
	id := ctx.Paramx("id").Uint64()
	fileM := file.NewFile(ctx)
	ownerID := uint64(user.Id)
	if id == 0 {
		ids := ctx.FormxValues(`id`).Uint64()
		for _, id := range ids {
			err = fileM.DeleteByID(id, `user`, ownerID)
			if err != nil {
				return err
			}
		}
		goto END
	}
	err = fileM.DeleteByID(id, `user`, ownerID)
	if err != nil {
		return err
	}

END:
	return ctx.Redirect(handler.URLFor(`/manager/file/list`))
}

// FileClean 删除未使用文件
func FileClean(ctx echo.Context) (err error) {
	fileM := file.NewFile(ctx)
	ago := ctx.Form(`ago`)
	var seconds int64 = 86400 * 365
	if len(ago) > 0 {
		t, e := time.ParseDuration(ago)
		if e != nil {
			return e
		}
		seconds = int64(t.Seconds())
	}
	err = fileM.RemoveUnused(seconds, ``, 0)
	if err != nil {
		return err
	}

	return ctx.Redirect(handler.URLFor(`/manager/file/list`))
}