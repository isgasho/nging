// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"time"

	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/param"
)

type Slice_TaskGroup []*TaskGroup

func (s Slice_TaskGroup) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_TaskGroup) RangeRaw(fn func(m *TaskGroup) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_TaskGroup) GroupBy(keyField string) map[string][]*TaskGroup {
	r := map[string][]*TaskGroup{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*TaskGroup{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (s Slice_TaskGroup) KeyBy(keyField string) map[string]*TaskGroup {
	r := map[string]*TaskGroup{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (s Slice_TaskGroup) AsKV(keyField string, valueField string) param.Store {
	r := param.Store{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (s Slice_TaskGroup) Transform(transfers map[string]param.Transfer) []param.Store {
	r := make([]param.Store, len(s))
	for idx, row := range s {
		r[idx] = row.AsMap().Transform(transfers)
	}
	return r
}

// TaskGroup 任务组
type TaskGroup struct {
	base    factory.Base
	objects []*TaskGroup

	Id          uint   `db:"id,omitempty,pk" bson:"id,omitempty" comment:"" json:"id" xml:"id"`
	Uid         uint   `db:"uid" bson:"uid" comment:"用户ID" json:"uid" xml:"uid"`
	Name        string `db:"name" bson:"name" comment:"组名" json:"name" xml:"name"`
	Description string `db:"description" bson:"description" comment:"说明" json:"description" xml:"description"`
	Created     uint   `db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Updated     uint   `db:"updated" bson:"updated" comment:"修改时间" json:"updated" xml:"updated"`
	CmdPrefix   string `db:"cmd_prefix" bson:"cmd_prefix" comment:"命令前缀" json:"cmd_prefix" xml:"cmd_prefix"`
	CmdSuffix   string `db:"cmd_suffix" bson:"cmd_suffix" comment:"命令后缀" json:"cmd_suffix" xml:"cmd_suffix"`
}

// - base function

func (a *TaskGroup) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *TaskGroup) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *TaskGroup) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *TaskGroup) EventON(on ...bool) factory.Model {
	a.base.EventON(on...)
	return a
}

func (a *TaskGroup) EventOFF(off ...bool) factory.Model {
	a.base.EventOFF(off...)
	return a
}

func (a *TaskGroup) Context() echo.Context {
	return a.base.Context()
}

func (a *TaskGroup) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *TaskGroup) SetNamer(namer func(string) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *TaskGroup) Namer() func(string) string {
	return a.base.Namer()
}

func (a *TaskGroup) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *TaskGroup) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

// - current function

func (a *TaskGroup) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName, connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName, a.base.ConnID()).Use(a.base.Trans())
}

func (a *TaskGroup) Objects() []*TaskGroup {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *TaskGroup) XObjects() Slice_TaskGroup {
	return Slice_TaskGroup(a.Objects())
}

func (a *TaskGroup) NewObjects() factory.Ranger {
	return &Slice_TaskGroup{}
}

func (a *TaskGroup) InitObjects() *[]*TaskGroup {
	a.objects = []*TaskGroup{}
	return &a.objects
}

func (a *TaskGroup) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *TaskGroup) Short_() string {
	return "task_group"
}

func (a *TaskGroup) Struct_() string {
	return "TaskGroup"
}

func (a *TaskGroup) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *TaskGroup) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *TaskGroup) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	base := a.base
	err := a.Param(mw, args...).SetRecv(a).One()
	a.base = base
	return err
}

func (a *TaskGroup) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	return a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv).List()
}

func (a *TaskGroup) GroupBy(keyField string, inputRows ...[]*TaskGroup) map[string][]*TaskGroup {
	var rows Slice_TaskGroup
	if len(inputRows) > 0 {
		rows = Slice_TaskGroup(inputRows[0])
	} else {
		rows = Slice_TaskGroup(a.Objects())
	}
	return rows.GroupBy(keyField)
}

func (a *TaskGroup) KeyBy(keyField string, inputRows ...[]*TaskGroup) map[string]*TaskGroup {
	var rows Slice_TaskGroup
	if len(inputRows) > 0 {
		rows = Slice_TaskGroup(inputRows[0])
	} else {
		rows = Slice_TaskGroup(a.Objects())
	}
	return rows.KeyBy(keyField)
}

func (a *TaskGroup) AsKV(keyField string, valueField string, inputRows ...[]*TaskGroup) param.Store {
	var rows Slice_TaskGroup
	if len(inputRows) > 0 {
		rows = Slice_TaskGroup(inputRows[0])
	} else {
		rows = Slice_TaskGroup(a.Objects())
	}
	return rows.AsKV(keyField, valueField)
}

func (a *TaskGroup) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	return a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv).List()
}

func (a *TaskGroup) Add() (pk interface{}, err error) {
	a.Created = uint(time.Now().Unix())
	a.Id = 0
	if a.base.Eventable() {
		err = DBI.Fire("creating", a, nil)
		if err != nil {
			return
		}
	}
	pk, err = a.Param(nil).SetSend(a).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint(v)
		}
	}
	if err == nil && a.base.Eventable() {
		err = DBI.Fire("created", a, nil)
	}
	return
}

func (a *TaskGroup) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	a.Updated = uint(time.Now().Unix())
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(a).Update()
	}
	if err = DBI.Fire("updating", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(a).Update(); err != nil {
		return
	}
	return DBI.Fire("updated", a, mw, args...)
}

func (a *TaskGroup) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *TaskGroup) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {
	kvset["updated"] = uint(time.Now().Unix())
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(kvset).Update()
	}
	m := *a
	m.FromRow(kvset)
	var editColumns []string
	for column := range kvset {
		editColumns = append(editColumns, column)
	}
	if err = DBI.FireUpdate("updating", &m, editColumns, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(kvset).Update(); err != nil {
		return
	}
	return DBI.FireUpdate("updated", &m, editColumns, mw, args...)
}

func (a *TaskGroup) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
		a.Updated = uint(time.Now().Unix())
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		a.Created = uint(time.Now().Unix())
		a.Id = 0
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("creating", a, nil)
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint(v)
		}
	}
	if err == nil && a.base.Eventable() {
		if pk == nil {
			err = DBI.Fire("updated", a, mw, args...)
		} else {
			err = DBI.Fire("created", a, nil)
		}
	}
	return
}

func (a *TaskGroup) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if !a.base.Eventable() {
		return a.Param(mw, args...).Delete()
	}
	if err = DBI.Fire("deleting", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).Delete(); err != nil {
		return
	}
	return DBI.Fire("deleted", a, mw, args...)
}

func (a *TaskGroup) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *TaskGroup) Reset() *TaskGroup {
	a.Id = 0
	a.Uid = 0
	a.Name = ``
	a.Description = ``
	a.Created = 0
	a.Updated = 0
	a.CmdPrefix = ``
	a.CmdSuffix = ``
	return a
}

func (a *TaskGroup) AsMap() param.Store {
	r := param.Store{}
	r["Id"] = a.Id
	r["Uid"] = a.Uid
	r["Name"] = a.Name
	r["Description"] = a.Description
	r["Created"] = a.Created
	r["Updated"] = a.Updated
	r["CmdPrefix"] = a.CmdPrefix
	r["CmdSuffix"] = a.CmdSuffix
	return r
}

func (a *TaskGroup) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint(value)
		case "uid":
			a.Uid = param.AsUint(value)
		case "name":
			a.Name = param.AsString(value)
		case "description":
			a.Description = param.AsString(value)
		case "created":
			a.Created = param.AsUint(value)
		case "updated":
			a.Updated = param.AsUint(value)
		case "cmd_prefix":
			a.CmdPrefix = param.AsString(value)
		case "cmd_suffix":
			a.CmdSuffix = param.AsString(value)
		}
	}
}

func (a *TaskGroup) Set(key interface{}, value ...interface{}) {
	switch k := key.(type) {
	case map[string]interface{}:
		for kk, vv := range k {
			a.Set(kk, vv)
		}
	default:
		var (
			kk string
			vv interface{}
		)
		if k, y := key.(string); y {
			kk = k
		} else {
			kk = fmt.Sprint(key)
		}
		if len(value) > 0 {
			vv = value[0]
		}
		switch kk {
		case "Id":
			a.Id = param.AsUint(vv)
		case "Uid":
			a.Uid = param.AsUint(vv)
		case "Name":
			a.Name = param.AsString(vv)
		case "Description":
			a.Description = param.AsString(vv)
		case "Created":
			a.Created = param.AsUint(vv)
		case "Updated":
			a.Updated = param.AsUint(vv)
		case "CmdPrefix":
			a.CmdPrefix = param.AsString(vv)
		case "CmdSuffix":
			a.CmdSuffix = param.AsString(vv)
		}
	}
}

func (a *TaskGroup) AsRow() param.Store {
	r := param.Store{}
	r["id"] = a.Id
	r["uid"] = a.Uid
	r["name"] = a.Name
	r["description"] = a.Description
	r["created"] = a.Created
	r["updated"] = a.Updated
	r["cmd_prefix"] = a.CmdPrefix
	r["cmd_suffix"] = a.CmdSuffix
	return r
}

func (a *TaskGroup) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return factory.BatchValidate(a.Short_(), kvset)
}

func (a *TaskGroup) Validate(field string, value interface{}) error {
	return factory.Validate(a.Short_(), field, value)
}
