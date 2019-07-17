// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	
	"time"
)

type TaskGroup struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*TaskGroup
	namer   func(string) string
	connID  int
	
	Id         	uint    	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"" json:"id" xml:"id"`
	Uid        	uint    	`db:"uid" bson:"uid" comment:"用户ID" json:"uid" xml:"uid"`
	Name       	string  	`db:"name" bson:"name" comment:"组名" json:"name" xml:"name"`
	Description	string  	`db:"description" bson:"description" comment:"说明" json:"description" xml:"description"`
	Created    	uint    	`db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Updated    	uint    	`db:"updated" bson:"updated" comment:"修改时间" json:"updated" xml:"updated"`
	CmdPrefix  	string  	`db:"cmd_prefix" bson:"cmd_prefix" comment:"命令前缀" json:"cmd_prefix" xml:"cmd_prefix"`
	CmdSuffix  	string  	`db:"cmd_suffix" bson:"cmd_suffix" comment:"命令后缀" json:"cmd_suffix" xml:"cmd_suffix"`
}

func (this *TaskGroup) Trans() *factory.Transaction {
	return this.trans
}

func (this *TaskGroup) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *TaskGroup) SetConnID(connID int) factory.Model {
	this.connID = connID
	return this
}

func (this *TaskGroup) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.connID).Use(this.trans)
}

func (this *TaskGroup) Objects() []*TaskGroup {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *TaskGroup) NewObjects() *[]*TaskGroup {
	this.objects = []*TaskGroup{}
	return &this.objects
}

func (this *TaskGroup) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *TaskGroup) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *TaskGroup) Short_() string {
	return "task_group"
}

func (this *TaskGroup) Struct_() string {
	return "TaskGroup"
}

func (this *TaskGroup) Name_() string {
	if this.namer != nil {
		return WithPrefix(this.namer(this.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(this.Short_())(this))
}

func (this *TaskGroup) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *TaskGroup) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *TaskGroup) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *TaskGroup) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *TaskGroup) GroupBy(keyField string, inputRows ...[]*TaskGroup) map[string][]*TaskGroup {
	var rows []*TaskGroup
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string][]*TaskGroup{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*TaskGroup{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (this *TaskGroup) KeyBy(keyField string, inputRows ...[]*TaskGroup) map[string]*TaskGroup {
	var rows []*TaskGroup
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]*TaskGroup{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (this *TaskGroup) AsKV(keyField string, valueField string, inputRows ...[]*TaskGroup) map[string]interface{} {
	var rows []*TaskGroup
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

func (this *TaskGroup) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *TaskGroup) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
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

func (this *TaskGroup) Edit(mw func(db.Result) db.Result, args ...interface{}) error {
	this.Updated = uint(time.Now().Unix())
	return this.Setter(mw, args...).SetSend(this).Update()
}

func (this *TaskGroup) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *TaskGroup) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) error {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *TaskGroup) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) error {
	kvset["updated"] = uint(time.Now().Unix())
	return this.Setter(mw, args...).SetSend(kvset).Update()
}

func (this *TaskGroup) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func(){
		this.Updated = uint(time.Now().Unix())
	},func(){
		this.Created = uint(time.Now().Unix())
	this.Id = 0
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

func (this *TaskGroup) Delete(mw func(db.Result) db.Result, args ...interface{}) error {
	
	return this.Param().SetArgs(args...).SetMiddleware(mw).Delete()
}

func (this *TaskGroup) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *TaskGroup) Reset() *TaskGroup {
	this.Id = 0
	this.Uid = 0
	this.Name = ``
	this.Description = ``
	this.Created = 0
	this.Updated = 0
	this.CmdPrefix = ``
	this.CmdSuffix = ``
	return this
}

func (this *TaskGroup) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["Uid"] = this.Uid
	r["Name"] = this.Name
	r["Description"] = this.Description
	r["Created"] = this.Created
	r["Updated"] = this.Updated
	r["CmdPrefix"] = this.CmdPrefix
	r["CmdSuffix"] = this.CmdSuffix
	return r
}

func (this *TaskGroup) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = this.Id
	r["uid"] = this.Uid
	r["name"] = this.Name
	r["description"] = this.Description
	r["created"] = this.Created
	r["updated"] = this.Updated
	r["cmd_prefix"] = this.CmdPrefix
	r["cmd_suffix"] = this.CmdSuffix
	return r
}

func (this *TaskGroup) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = this.AsRow()
	}
	return factory.BatchValidate(this.Short_(), kvset)
}

func (this *TaskGroup) Validate(field string, value interface{}) error {
	return factory.Validate(this.Short_(), field, value)
}

