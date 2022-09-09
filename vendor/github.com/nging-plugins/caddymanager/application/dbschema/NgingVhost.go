// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"time"

	"github.com/webx-top/com"
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/db/lib/factory/pagination"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/param"
)

type Slice_NgingVhost []*NgingVhost

func (s Slice_NgingVhost) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingVhost) RangeRaw(fn func(m *NgingVhost) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingVhost) GroupBy(keyField string) map[string][]*NgingVhost {
	r := map[string][]*NgingVhost{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*NgingVhost{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (s Slice_NgingVhost) KeyBy(keyField string) map[string]*NgingVhost {
	r := map[string]*NgingVhost{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (s Slice_NgingVhost) AsKV(keyField string, valueField string) param.Store {
	r := param.Store{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (s Slice_NgingVhost) Transform(transfers map[string]param.Transfer) []param.Store {
	r := make([]param.Store, len(s))
	for idx, row := range s {
		r[idx] = row.AsMap().Transform(transfers)
	}
	return r
}

func (s Slice_NgingVhost) FromList(data interface{}) Slice_NgingVhost {
	values, ok := data.([]*NgingVhost)
	if !ok {
		for _, value := range data.([]interface{}) {
			row := &NgingVhost{}
			row.FromRow(value.(map[string]interface{}))
			s = append(s, row)
		}
		return s
	}
	s = append(s, values...)

	return s
}

func NewNgingVhost(ctx echo.Context) *NgingVhost {
	m := &NgingVhost{}
	m.SetContext(ctx)
	return m
}

// NgingVhost 虚拟主机
type NgingVhost struct {
	base    factory.Base
	objects []*NgingVhost

	Id       uint   `db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Name     string `db:"name" bson:"name" comment:"网站名称" json:"name" xml:"name"`
	GroupId  uint   `db:"group_id" bson:"group_id" comment:"组" json:"group_id" xml:"group_id"`
	Domain   string `db:"domain" bson:"domain" comment:"域名" json:"domain" xml:"domain"`
	Root     string `db:"root" bson:"root" comment:"网站物理路径" json:"root" xml:"root"`
	Created  uint   `db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Updated  uint   `db:"updated" bson:"updated" comment:"更新时间" json:"updated" xml:"updated"`
	Setting  string `db:"setting" bson:"setting" comment:"设置" json:"setting" xml:"setting"`
	Disabled string `db:"disabled" bson:"disabled" comment:"是否停用" json:"disabled" xml:"disabled"`
}

// - base function

func (a *NgingVhost) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *NgingVhost) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *NgingVhost) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *NgingVhost) EventON(on ...bool) factory.Model {
	a.base.EventON(on...)
	return a
}

func (a *NgingVhost) EventOFF(off ...bool) factory.Model {
	a.base.EventOFF(off...)
	return a
}

func (a *NgingVhost) Context() echo.Context {
	return a.base.Context()
}

func (a *NgingVhost) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *NgingVhost) SetNamer(namer func(factory.Model) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *NgingVhost) Namer() func(factory.Model) string {
	return a.base.Namer()
}

func (a *NgingVhost) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *NgingVhost) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

// - current function

func (a *NgingVhost) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName, connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName, a.base.ConnID()).Use(a.base.Trans())
}

func (a *NgingVhost) Objects() []*NgingVhost {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *NgingVhost) XObjects() Slice_NgingVhost {
	return Slice_NgingVhost(a.Objects())
}

func (a *NgingVhost) NewObjects() factory.Ranger {
	return &Slice_NgingVhost{}
}

func (a *NgingVhost) InitObjects() *[]*NgingVhost {
	a.objects = []*NgingVhost{}
	return &a.objects
}

func (a *NgingVhost) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *NgingVhost) Short_() string {
	return "nging_vhost"
}

func (a *NgingVhost) Struct_() string {
	return "NgingVhost"
}

func (a *NgingVhost) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *NgingVhost) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *NgingVhost) Get(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	base := a.base
	if !a.base.Eventable() {
		err = a.Param(mw, args...).SetRecv(a).One()
		a.base = base
		return
	}
	queryParam := a.Param(mw, args...).SetRecv(a)
	if err = DBI.FireReading(a, queryParam); err != nil {
		return
	}
	err = queryParam.One()
	a.base = base
	if err == nil {
		err = DBI.FireReaded(a, queryParam)
	}
	return
}

func (a *NgingVhost) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv).List()
	}
	queryParam := a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv)
	if err := DBI.FireReading(a, queryParam); err != nil {
		return nil, err
	}
	cnt, err := queryParam.List()
	if err == nil {
		switch v := recv.(type) {
		case *[]*NgingVhost:
			err = DBI.FireReaded(a, queryParam, Slice_NgingVhost(*v))
		case []*NgingVhost:
			err = DBI.FireReaded(a, queryParam, Slice_NgingVhost(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingVhost) GroupBy(keyField string, inputRows ...[]*NgingVhost) map[string][]*NgingVhost {
	var rows Slice_NgingVhost
	if len(inputRows) > 0 {
		rows = Slice_NgingVhost(inputRows[0])
	} else {
		rows = Slice_NgingVhost(a.Objects())
	}
	return rows.GroupBy(keyField)
}

func (a *NgingVhost) KeyBy(keyField string, inputRows ...[]*NgingVhost) map[string]*NgingVhost {
	var rows Slice_NgingVhost
	if len(inputRows) > 0 {
		rows = Slice_NgingVhost(inputRows[0])
	} else {
		rows = Slice_NgingVhost(a.Objects())
	}
	return rows.KeyBy(keyField)
}

func (a *NgingVhost) AsKV(keyField string, valueField string, inputRows ...[]*NgingVhost) param.Store {
	var rows Slice_NgingVhost
	if len(inputRows) > 0 {
		rows = Slice_NgingVhost(inputRows[0])
	} else {
		rows = Slice_NgingVhost(a.Objects())
	}
	return rows.AsKV(keyField, valueField)
}

func (a *NgingVhost) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv).List()
	}
	queryParam := a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv)
	if err := DBI.FireReading(a, queryParam); err != nil {
		return nil, err
	}
	cnt, err := queryParam.List()
	if err == nil {
		switch v := recv.(type) {
		case *[]*NgingVhost:
			err = DBI.FireReaded(a, queryParam, Slice_NgingVhost(*v))
		case []*NgingVhost:
			err = DBI.FireReaded(a, queryParam, Slice_NgingVhost(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingVhost) Insert() (pk interface{}, err error) {
	a.Created = uint(time.Now().Unix())
	a.Id = 0
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
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

func (a *NgingVhost) Update(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
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

func (a *NgingVhost) Updatex(mw func(db.Result) db.Result, args ...interface{}) (affected int64, err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(a).Updatex()
	}
	if err = DBI.Fire("updating", a, mw, args...); err != nil {
		return
	}
	if affected, err = a.Param(mw, args...).SetSend(a).Updatex(); err != nil {
		return
	}
	err = DBI.Fire("updated", a, mw, args...)
	return
}

func (a *NgingVhost) UpdateByFields(mw func(db.Result) db.Result, fields []string, args ...interface{}) (err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).UpdateByStruct(a, fields...)
	}
	editColumns := make([]string, len(fields))
	for index, field := range fields {
		editColumns[index] = com.SnakeCase(field)
	}
	if err = DBI.FireUpdate("updating", a, editColumns, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).UpdateByStruct(a, fields...); err != nil {
		return
	}
	err = DBI.FireUpdate("updated", a, editColumns, mw, args...)
	return
}

func (a *NgingVhost) UpdatexByFields(mw func(db.Result) db.Result, fields []string, args ...interface{}) (affected int64, err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).UpdatexByStruct(a, fields...)
	}
	editColumns := make([]string, len(fields))
	for index, field := range fields {
		editColumns[index] = com.SnakeCase(field)
	}
	if err = DBI.FireUpdate("updating", a, editColumns, mw, args...); err != nil {
		return
	}
	if affected, err = a.Param(mw, args...).UpdatexByStruct(a, fields...); err != nil {
		return
	}
	err = DBI.FireUpdate("updated", a, editColumns, mw, args...)
	return
}

func (a *NgingVhost) UpdateField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.UpdateFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingVhost) UpdateFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

	if val, ok := kvset["disabled"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["disabled"] = "N"
		}
	}
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

func (a *NgingVhost) UpdateValues(mw func(db.Result) db.Result, keysValues *db.KeysValues, args ...interface{}) (err error) {
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(keysValues).Update()
	}
	m := *a
	m.FromRow(keysValues.Map())
	if err = DBI.FireUpdate("updating", &m, keysValues.Keys(), mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(keysValues).Update(); err != nil {
		return
	}
	return DBI.FireUpdate("updated", &m, keysValues.Keys(), mw, args...)
}

func (a *NgingVhost) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
		a.Updated = uint(time.Now().Unix())
		if len(a.Disabled) == 0 {
			a.Disabled = "N"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		a.Created = uint(time.Now().Unix())
		a.Id = 0
		if len(a.Disabled) == 0 {
			a.Disabled = "N"
		}
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

func (a *NgingVhost) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

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

func (a *NgingVhost) Deletex(mw func(db.Result) db.Result, args ...interface{}) (affected int64, err error) {

	if !a.base.Eventable() {
		return a.Param(mw, args...).Deletex()
	}
	if err = DBI.Fire("deleting", a, mw, args...); err != nil {
		return
	}
	if affected, err = a.Param(mw, args...).Deletex(); err != nil {
		return
	}
	err = DBI.Fire("deleted", a, mw, args...)
	return
}

func (a *NgingVhost) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *NgingVhost) Exists(mw func(db.Result) db.Result, args ...interface{}) (bool, error) {
	return a.Param(mw, args...).Exists()
}

func (a *NgingVhost) Reset() *NgingVhost {
	a.Id = 0
	a.Name = ``
	a.GroupId = 0
	a.Domain = ``
	a.Root = ``
	a.Created = 0
	a.Updated = 0
	a.Setting = ``
	a.Disabled = ``
	return a
}

func (a *NgingVhost) AsMap(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["Id"] = a.Id
		r["Name"] = a.Name
		r["GroupId"] = a.GroupId
		r["Domain"] = a.Domain
		r["Root"] = a.Root
		r["Created"] = a.Created
		r["Updated"] = a.Updated
		r["Setting"] = a.Setting
		r["Disabled"] = a.Disabled
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "Id":
			r["Id"] = a.Id
		case "Name":
			r["Name"] = a.Name
		case "GroupId":
			r["GroupId"] = a.GroupId
		case "Domain":
			r["Domain"] = a.Domain
		case "Root":
			r["Root"] = a.Root
		case "Created":
			r["Created"] = a.Created
		case "Updated":
			r["Updated"] = a.Updated
		case "Setting":
			r["Setting"] = a.Setting
		case "Disabled":
			r["Disabled"] = a.Disabled
		}
	}
	return r
}

func (a *NgingVhost) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint(value)
		case "name":
			a.Name = param.AsString(value)
		case "group_id":
			a.GroupId = param.AsUint(value)
		case "domain":
			a.Domain = param.AsString(value)
		case "root":
			a.Root = param.AsString(value)
		case "created":
			a.Created = param.AsUint(value)
		case "updated":
			a.Updated = param.AsUint(value)
		case "setting":
			a.Setting = param.AsString(value)
		case "disabled":
			a.Disabled = param.AsString(value)
		}
	}
}

func (a *NgingVhost) Set(key interface{}, value ...interface{}) {
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
		case "Name":
			a.Name = param.AsString(vv)
		case "GroupId":
			a.GroupId = param.AsUint(vv)
		case "Domain":
			a.Domain = param.AsString(vv)
		case "Root":
			a.Root = param.AsString(vv)
		case "Created":
			a.Created = param.AsUint(vv)
		case "Updated":
			a.Updated = param.AsUint(vv)
		case "Setting":
			a.Setting = param.AsString(vv)
		case "Disabled":
			a.Disabled = param.AsString(vv)
		}
	}
}

func (a *NgingVhost) AsRow(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["id"] = a.Id
		r["name"] = a.Name
		r["group_id"] = a.GroupId
		r["domain"] = a.Domain
		r["root"] = a.Root
		r["created"] = a.Created
		r["updated"] = a.Updated
		r["setting"] = a.Setting
		r["disabled"] = a.Disabled
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "id":
			r["id"] = a.Id
		case "name":
			r["name"] = a.Name
		case "group_id":
			r["group_id"] = a.GroupId
		case "domain":
			r["domain"] = a.Domain
		case "root":
			r["root"] = a.Root
		case "created":
			r["created"] = a.Created
		case "updated":
			r["updated"] = a.Updated
		case "setting":
			r["setting"] = a.Setting
		case "disabled":
			r["disabled"] = a.Disabled
		}
	}
	return r
}

func (a *NgingVhost) ListPage(cond *db.Compounds, sorts ...interface{}) error {
	_, err := pagination.NewLister(a, nil, func(r db.Result) db.Result {
		return r.OrderBy(sorts...)
	}, cond.And()).Paging(a.Context())
	return err
}

func (a *NgingVhost) ListPageAs(recv interface{}, cond *db.Compounds, sorts ...interface{}) error {
	_, err := pagination.NewLister(a, recv, func(r db.Result) db.Result {
		return r.OrderBy(sorts...)
	}, cond.And()).Paging(a.Context())
	return err
}

func (a *NgingVhost) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return DBI.Fields.BatchValidate(a.Short_(), kvset)
}

func (a *NgingVhost) Validate(field string, value interface{}) error {
	return DBI.Fields.Validate(a.Short_(), field, value)
}