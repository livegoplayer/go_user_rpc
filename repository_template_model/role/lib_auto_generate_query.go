package role

import "reflect"
import "github.com/livegoplayer/go_db_helper/mysql"

/**
@Collect
*/
type RoleCollect []Role

func NewRoleQuery() *RoleQuery {
	s := RoleQuery{}
	s.SetBuild(mysql.NewBuild(&s))
	i, ok := reflect.ValueOf(&s).Interface().(BeforeHook)
	if ok {
		i.Before()
	}
	return &s
}

type _RoleQueryColsStruct struct {
	Id          string
	RoleName    string
	RoleLevel   string
	AddDatetime string
	UptDatetime string
}

func GetRoleQueryCols() *_RoleQueryColsStruct {
	return &_RoleQueryColsStruct{
		Id:          "id",
		RoleName:    "role_name",
		RoleLevel:   "role_level",
		AddDatetime: "add_datetime",
		UptDatetime: "upt_datetime",
	}
}

func (m *RoleQuery) First() *Role {
	s := make([]Role, 0)
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	m.GetBuild().ModelType(&s).Limit(1).Get()
	if len(s) > 0 {
		return &s[0]
	}
	return &Role{}
}

func (m *RoleQuery) GetOne() *Role {
	s := make([]Role, 0)
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	m.GetBuild().ModelType(&s).Limit(1).Get()
	if len(s) > 0 {
		return &s[0]
	}
	return nil
}

func (m *RoleQuery) Get() RoleCollect {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := make([]Role, 0)
	m.GetBuild().ModelType(&s).Get()
	return s
}

func (m *RoleQuery) clone() *RoleQuery {
	nm := NewRoleQuery()
	nm.SetBuild(m.GetBuild().Clone())
	return nm
}

func (m *RoleQuery) Count() int64 {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := Role{}
	return m.GetBuild().ModelType(&s).Count()
}

func (m *RoleQuery) sum(col string) float64 {
	s := Role{}
	return m.GetBuild().ModelType(&s).Sum(col)
}

func (m *RoleQuery) max(col string) float64 {
	s := Role{}
	return m.GetBuild().ModelType(&s).Max(col)
}

func (m *RoleQuery) DoneOperate() int64 {
	s := Role{}
	return m.GetBuild().ModelType(&s).DoneOperate()
}

func (m *RoleQuery) update(h *Role) int64 {
	s := Role{}
	return m.GetBuild().ModelType(&s).Update(h)
}

func (m *RoleQuery) Delete() int64 {
	s := Role{}
	return m.GetBuild().ModelType(&s).Delete()
}

func (m *RoleQuery) save(x *Role) {
	m.GetBuild().ModelType(x).Save()
}

func (m *RoleQuery) error() error {
	return m.GetBuild().ModelType(m).Error()
}

//支持分表
func (m *RoleQuery) insert(argu interface{}) {
	s := Role{}
	m.GetBuild().ModelType(&s).Insert(argu)
}

func (m *RoleQuery) increment(column string, amount int) int64 {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := Role{}
	return m.GetBuild().ModelType(&s).Increment(column, amount)
}

func (m *RoleQuery) Unscoped() *RoleQuery {
	m.GetBuild().Unscoped()
	return m
}

func (m *RoleQuery) Select(columns ...interface{}) *RoleQuery {
	m.GetBuild().Select(columns...)
	return m
}

func (m *RoleQuery) where(column string, args ...interface{}) *RoleQuery {
	m.GetBuild().Where(column, args...)
	return m
}

func (m *RoleQuery) whereMap(datas map[string]interface{}) *RoleQuery {
	m.GetBuild().WhereMap(datas)
	return m
}

func (m *RoleQuery) orWhere(column string, args ...interface{}) *RoleQuery {
	m.GetBuild().OrWhere(column, args...)
	return m
}

func (m *RoleQuery) orWhereCb(cb WhereCb) *RoleQuery {
	m.GetBuild().OrWhereCb(cb)
	return m
}

func (m *RoleQuery) whereCb(cb WhereCb) *RoleQuery {
	m.GetBuild().WhereCb(cb)
	return m
}

func (m *RoleQuery) whereRaw(sql string, bindings ...interface{}) *RoleQuery {
	m.GetBuild().WhereRaw(sql, bindings...)
	return m
}

func (m *RoleQuery) whereNull(column string) *RoleQuery {
	m.GetBuild().WhereNull(column)
	return m
}

func (m *RoleQuery) whereNotNull(column string) *RoleQuery {
	m.GetBuild().WhereNotNull(column)
	return m
}

func (m *RoleQuery) orWhereNull(column string) *RoleQuery {
	m.GetBuild().OrWhereNull(column)
	return m
}

func (m *RoleQuery) orWhereNotNull(column string) *RoleQuery {
	m.GetBuild().OrWhereNotNull(column)
	return m
}

func (m *RoleQuery) whereIn(column string, values interface{}) *RoleQuery {
	m.GetBuild().WhereIn(column, values)
	return m
}

func (m *RoleQuery) whereNotIn(column string, values interface{}) *RoleQuery {
	m.GetBuild().WhereNotIn(column, values)
	return m
}

func (m *RoleQuery) LeftJoin(table string, one string, args ...interface{}) *RoleQuery {
	m.GetBuild().LeftJoin(table, one, args...)
	return m
}

func (m *RoleQuery) RightJoin(table string, one string, args ...interface{}) *RoleQuery {
	m.GetBuild().RightJoin(table, one, args...)
	return m
}

func (m *RoleQuery) InnerJoin(table string, one string, args ...interface{}) *RoleQuery {
	m.GetBuild().InnerJoin(table, one, args...)
	return m
}

func (m *RoleQuery) OrderBy(column string, direction string) *RoleQuery {
	m.GetBuild().OrderBy(column, direction)
	return m
}

func (m *RoleQuery) OrderDescBy(column string) *RoleQuery {
	m.GetBuild().OrderDescBy(column)
	return m
}

func (m *RoleQuery) OrderAscBy(column string) *RoleQuery {
	m.GetBuild().OrderAscBy(column)
	return m
}

func (m *RoleQuery) Group(column string) *RoleQuery {
	m.GetBuild().Group(column)
	return m
}

func (m *RoleQuery) Inc(col string, num int) *RoleQuery {
	m.GetBuild().Inc(col, num)
	return m
}

func (m *RoleQuery) Set(col string, val interface{}) *RoleQuery {
	m.GetBuild().Set(col, val)
	return m
}

func (m *RoleQuery) Skip(lines int) *RoleQuery {
	m.GetBuild().Skip(lines)
	return m
}

func (m *RoleQuery) Offset(lines int) *RoleQuery {
	m.GetBuild().Offset(lines)
	return m
}

func (m *RoleQuery) Limit(lines int) *RoleQuery {
	m.GetBuild().Limit(lines)
	return m
}

func (m *RoleQuery) Take(lines int) *RoleQuery {
	m.GetBuild().Take(lines)
	return m
}

func (m *RoleQuery) ForUpdate() *RoleQuery {
	m.GetBuild().ForUpdate()
	return m
}

func (m *RoleQuery) SplitBy(val int64) *RoleQuery {
	m.GetBuild().SplitBy(val)
	return m
}

func (m *RoleQuery) ModelType(v interface{}) *RoleQuery {
	m.GetBuild().ModelType(v)
	return m
}

func (m *RoleQuery) DisablePanic() *RoleQuery {
	m.GetBuild().DisablePanic()
	return m
}

func (m *RoleQuery) OnWrite() *RoleQuery {
	m.GetBuild().OnWrite()
	return m
}

func (m *RoleQuery) kWheId(args ...interface{}) *RoleQuery {
	return m.where("id", args...)
}

func (m *RoleQuery) kWheRoleName(args ...interface{}) *RoleQuery {
	return m.where("role_name", args...)
}

func (m *RoleQuery) kWheRoleLevel(args ...interface{}) *RoleQuery {
	return m.where("role_level", args...)
}

func (m *RoleQuery) kWheAddDatetime(args ...interface{}) *RoleQuery {
	return m.where("add_datetime", args...)
}

func (m *RoleQuery) kWheUptDatetime(args ...interface{}) *RoleQuery {
	return m.where("upt_datetime", args...)
}

func (m *RoleQuery) kSetId(val interface{}) *RoleQuery {
	return m.Set("id", val)
}

func (m *RoleQuery) kSetRoleName(val interface{}) *RoleQuery {
	return m.Set("role_name", val)
}

func (m *RoleQuery) kSetRoleLevel(val interface{}) *RoleQuery {
	return m.Set("role_level", val)
}

func (m *RoleQuery) kSetAddDatetime(val interface{}) *RoleQuery {
	return m.Set("add_datetime", val)
}

func (m *RoleQuery) kSetUptDatetime(val interface{}) *RoleQuery {
	return m.Set("upt_datetime", val)
}

func (m *RoleQuery) kIncId(num int) *RoleQuery {
	return m.Inc("id", num)
}

func (m *RoleQuery) kIncRoleLevel(num int) *RoleQuery {
	return m.Inc("role_level", num)
}

func (m *RoleQuery) kWheIdIn(values interface{}) *RoleQuery {
	return m.whereIn("id", values)
}

func (m *RoleQuery) kWheRoleNameIn(values interface{}) *RoleQuery {
	return m.whereIn("role_name", values)
}

func (m *RoleQuery) kWheRoleLevelIn(values interface{}) *RoleQuery {
	return m.whereIn("role_level", values)
}

func (m *RoleQuery) kWheAddDatetimeIn(values interface{}) *RoleQuery {
	return m.whereIn("add_datetime", values)
}

func (m *RoleQuery) kWheUptDatetimeIn(values interface{}) *RoleQuery {
	return m.whereIn("upt_datetime", values)
}

func (m *RoleQuery) kWheIdNotIn(values interface{}) *RoleQuery {
	return m.whereNotIn("id", values)
}

func (m *RoleQuery) kWheRoleNameNotIn(values interface{}) *RoleQuery {
	return m.whereNotIn("role_name", values)
}

func (m *RoleQuery) kWheRoleLevelNotIn(values interface{}) *RoleQuery {
	return m.whereNotIn("role_level", values)
}

func (m *RoleQuery) kWheAddDatetimeNotIn(values interface{}) *RoleQuery {
	return m.whereNotIn("add_datetime", values)
}

func (m *RoleQuery) kWheUptDatetimeNotIn(values interface{}) *RoleQuery {
	return m.whereNotIn("upt_datetime", values)
}

func Save(f *Role) *Role {
	NewRoleQuery().save(f)
	return f
}

func UpdateRoleAll(p *Role) int64 {
	build := NewRoleQuery()
	return build.update(p)
}

func FetchRoleAll() RoleCollect {
	build := NewRoleQuery()
	return build.Get()
}

func FetchRoleAllWithPageSize(page int, pageSize int) RoleCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewRoleQuery()
	return build.Skip(offset).Limit(pageSize).Get()
}

func CountRoleAll() int64 {
	build := NewRoleQuery()
	return build.Count()
}

// uniIndex
func UpdateRoleById(id int64, p *Role) int64 {
	build := NewRoleQuery()

	build.kWheId(id)

	return build.update(p)
}

func UpdateRoleByIds(id []int64, p *Role) int64 {
	build := NewRoleQuery()

	if len(id) == 0 {
		return 0
	}

	if len(id) == 1 {
		build.kWheId(id[0])
	} else {
		build.kWheIdIn(id)
	}

	return build.update(p)
}

func UpdateRoleByIdsWhatEver(id []int64, p *Role) int64 {
	build := NewRoleQuery()

	if len(id) == 1 {
		build.kWheId(id[0])
	} else {
		build.kWheIdIn(id)
	}

	return build.update(p)
}

func FetchByIds(id []int64) RoleCollect {
	build := NewRoleQuery()

	if len(id) == 0 {
		return make(RoleCollect, 0)
	}

	if len(id) == 1 {
		build.kWheId(id[0])
	} else {
		build.kWheIdIn(id)
	}

	return build.Get()
}

func FetchByIdsWithPageSize(id []int64, page int, pageSize int) RoleCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewRoleQuery()

	if len(id) == 0 {
		return make(RoleCollect, 0)
	}

	if len(id) == 1 {
		build.kWheId(id[0])
	} else {
		build.kWheIdIn(id)
	}

	return build.Skip(offset).Limit(pageSize).Get()
}

func CheckExistById(id int64) bool {
	build := NewRoleQuery()

	build.kWheId(id)

	cnt := build.Count()
	return cnt > 0
}

func GetOneById(id int64) *Role {
	build := NewRoleQuery()

	build.kWheId(id)

	return build.GetOne()
}

func DeleteById(id int64) int64 {
	build := NewRoleQuery()

	build.kWheId(id)

	return build.Delete()
}

func GetFirstById(id int64) *Role {
	build := NewRoleQuery()

	build.kWheId(id)

	return build.First()
}
