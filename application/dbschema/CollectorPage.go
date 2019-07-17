// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	
	"time"
)

type CollectorPage struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*CollectorPage
	namer   func(string) string
	connID  int
	
	Id            	uint    	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"" json:"id" xml:"id"`
	ParentId      	uint    	`db:"parent_id" bson:"parent_id" comment:"父级规则" json:"parent_id" xml:"parent_id"`
	RootId        	uint    	`db:"root_id" bson:"root_id" comment:"根页面ID" json:"root_id" xml:"root_id"`
	HasChild      	string  	`db:"has_child" bson:"has_child" comment:"是否有子级" json:"has_child" xml:"has_child"`
	Uid           	uint    	`db:"uid" bson:"uid" comment:"用户ID" json:"uid" xml:"uid"`
	GroupId       	uint    	`db:"group_id" bson:"group_id" comment:"规则组" json:"group_id" xml:"group_id"`
	Name          	string  	`db:"name" bson:"name" comment:"规则名" json:"name" xml:"name"`
	Description   	string  	`db:"description" bson:"description" comment:"说明" json:"description" xml:"description"`
	EnterUrl      	string  	`db:"enter_url" bson:"enter_url" comment:"入口网址模板(网址一行一个)" json:"enter_url" xml:"enter_url"`
	Sort          	int     	`db:"sort" bson:"sort" comment:"排序" json:"sort" xml:"sort"`
	Created       	uint    	`db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Browser       	string  	`db:"browser" bson:"browser" comment:"浏览器" json:"browser" xml:"browser"`
	Type          	string  	`db:"type" bson:"type" comment:"页面类型" json:"type" xml:"type"`
	ScopeRule     	string  	`db:"scope_rule" bson:"scope_rule" comment:"页面区域规则" json:"scope_rule" xml:"scope_rule"`
	DuplicateRule 	string  	`db:"duplicate_rule" bson:"duplicate_rule" comment:"去重规则(url-判断网址;rule-判断规则是否改过;content-判断网页内容是否改过;none-不去重)" json:"duplicate_rule" xml:"duplicate_rule"`
	ContentType   	string  	`db:"content_type" bson:"content_type" comment:"内容类型" json:"content_type" xml:"content_type"`
	Charset       	string  	`db:"charset" bson:"charset" comment:"字符集" json:"charset" xml:"charset"`
	Timeout       	uint    	`db:"timeout" bson:"timeout" comment:"超时时间(秒)" json:"timeout" xml:"timeout"`
	Waits         	string  	`db:"waits" bson:"waits" comment:"等待时间范围(秒),例如2-8" json:"waits" xml:"waits"`
	Proxy         	string  	`db:"proxy" bson:"proxy" comment:"代理地址" json:"proxy" xml:"proxy"`
}

func (this *CollectorPage) Trans() *factory.Transaction {
	return this.trans
}

func (this *CollectorPage) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *CollectorPage) SetConnID(connID int) factory.Model {
	this.connID = connID
	return this
}

func (this *CollectorPage) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.connID).Use(this.trans)
}

func (this *CollectorPage) Objects() []*CollectorPage {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *CollectorPage) NewObjects() *[]*CollectorPage {
	this.objects = []*CollectorPage{}
	return &this.objects
}

func (this *CollectorPage) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *CollectorPage) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *CollectorPage) Short_() string {
	return "collector_page"
}

func (this *CollectorPage) Struct_() string {
	return "CollectorPage"
}

func (this *CollectorPage) Name_() string {
	if this.namer != nil {
		return WithPrefix(this.namer(this.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(this.Short_())(this))
}

func (this *CollectorPage) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *CollectorPage) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *CollectorPage) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *CollectorPage) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *CollectorPage) GroupBy(keyField string, inputRows ...[]*CollectorPage) map[string][]*CollectorPage {
	var rows []*CollectorPage
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string][]*CollectorPage{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*CollectorPage{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (this *CollectorPage) KeyBy(keyField string, inputRows ...[]*CollectorPage) map[string]*CollectorPage {
	var rows []*CollectorPage
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]*CollectorPage{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (this *CollectorPage) AsKV(keyField string, valueField string, inputRows ...[]*CollectorPage) map[string]interface{} {
	var rows []*CollectorPage
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]interface{}{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (this *CollectorPage) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *CollectorPage) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.ContentType) == 0 { this.ContentType = "html" }
	if len(this.DuplicateRule) == 0 { this.DuplicateRule = "none" }
	if len(this.HasChild) == 0 { this.HasChild = "N" }
	if len(this.Type) == 0 { this.Type = "content" }
	pk, err = this.Param().SetSend(this).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint(v)
		}
	}
	return
}

func (this *CollectorPage) Edit(mw func(db.Result) db.Result, args ...interface{}) error {
	
	if len(this.ContentType) == 0 { this.ContentType = "html" }
	if len(this.DuplicateRule) == 0 { this.DuplicateRule = "none" }
	if len(this.HasChild) == 0 { this.HasChild = "N" }
	if len(this.Type) == 0 { this.Type = "content" }
	return this.Setter(mw, args...).SetSend(this).Update()
}

func (this *CollectorPage) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *CollectorPage) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) error {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *CollectorPage) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) error {
	
	if val, ok := kvset["content_type"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["content_type"] = "html" } }
	if val, ok := kvset["duplicate_rule"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["duplicate_rule"] = "none" } }
	if val, ok := kvset["has_child"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["has_child"] = "N" } }
	if val, ok := kvset["type"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["type"] = "content" } }
	return this.Setter(mw, args...).SetSend(kvset).Update()
}

func (this *CollectorPage) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func(){
		
	if len(this.ContentType) == 0 { this.ContentType = "html" }
	if len(this.DuplicateRule) == 0 { this.DuplicateRule = "none" }
	if len(this.HasChild) == 0 { this.HasChild = "N" }
	if len(this.Type) == 0 { this.Type = "content" }
	},func(){
		this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.ContentType) == 0 { this.ContentType = "html" }
	if len(this.DuplicateRule) == 0 { this.DuplicateRule = "none" }
	if len(this.HasChild) == 0 { this.HasChild = "N" }
	if len(this.Type) == 0 { this.Type = "content" }
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint(v)
		}
	}
	return 
}

func (this *CollectorPage) Delete(mw func(db.Result) db.Result, args ...interface{}) error {
	
	return this.Param().SetArgs(args...).SetMiddleware(mw).Delete()
}

func (this *CollectorPage) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *CollectorPage) Reset() *CollectorPage {
	this.Id = 0
	this.ParentId = 0
	this.RootId = 0
	this.HasChild = ``
	this.Uid = 0
	this.GroupId = 0
	this.Name = ``
	this.Description = ``
	this.EnterUrl = ``
	this.Sort = 0
	this.Created = 0
	this.Browser = ``
	this.Type = ``
	this.ScopeRule = ``
	this.DuplicateRule = ``
	this.ContentType = ``
	this.Charset = ``
	this.Timeout = 0
	this.Waits = ``
	this.Proxy = ``
	return this
}

func (this *CollectorPage) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["ParentId"] = this.ParentId
	r["RootId"] = this.RootId
	r["HasChild"] = this.HasChild
	r["Uid"] = this.Uid
	r["GroupId"] = this.GroupId
	r["Name"] = this.Name
	r["Description"] = this.Description
	r["EnterUrl"] = this.EnterUrl
	r["Sort"] = this.Sort
	r["Created"] = this.Created
	r["Browser"] = this.Browser
	r["Type"] = this.Type
	r["ScopeRule"] = this.ScopeRule
	r["DuplicateRule"] = this.DuplicateRule
	r["ContentType"] = this.ContentType
	r["Charset"] = this.Charset
	r["Timeout"] = this.Timeout
	r["Waits"] = this.Waits
	r["Proxy"] = this.Proxy
	return r
}

func (this *CollectorPage) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = this.Id
	r["parent_id"] = this.ParentId
	r["root_id"] = this.RootId
	r["has_child"] = this.HasChild
	r["uid"] = this.Uid
	r["group_id"] = this.GroupId
	r["name"] = this.Name
	r["description"] = this.Description
	r["enter_url"] = this.EnterUrl
	r["sort"] = this.Sort
	r["created"] = this.Created
	r["browser"] = this.Browser
	r["type"] = this.Type
	r["scope_rule"] = this.ScopeRule
	r["duplicate_rule"] = this.DuplicateRule
	r["content_type"] = this.ContentType
	r["charset"] = this.Charset
	r["timeout"] = this.Timeout
	r["waits"] = this.Waits
	r["proxy"] = this.Proxy
	return r
}

func (this *CollectorPage) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = this.AsRow()
	}
	return factory.BatchValidate(this.Short_(), kvset)
}

func (this *CollectorPage) Validate(field string, value interface{}) error {
	return factory.Validate(this.Short_(), field, value)
}

