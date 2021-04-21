package ret_user_role

import "reflect"
import "github.com/livegoplayer/go_db_helper/mysql"

/**
@Collect
*/
type RetUserRoleCollect []RetUserRole

func NewRetUserRoleQuery() *RetUserRoleQuery {
	s := RetUserRoleQuery{}
	s.SetBuild(mysql.NewBuild(&s))
	i, ok := reflect.ValueOf(&s).Interface().(BeforeHook)
	if ok {
		i.Before()
	}
	return &s
}

type _RetUserRoleQueryColsStruct struct {
	Id     string
	Uid    string
	RoleId string
}

func GetRetUserRoleQueryCols() *_RetUserRoleQueryColsStruct {
	return &_RetUserRoleQueryColsStruct{
		Id:     "id",
		Uid:    "uid",
		RoleId: "role_id",
	}
}

func (m *RetUserRoleQuery) First() *RetUserRole {
	s := make([]RetUserRole, 0)
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	m.GetBuild().ModelType(&s).Limit(1).Get()
	if len(s) > 0 {
		return &s[0]
	}
	return &RetUserRole{}
}

func (m *RetUserRoleQuery) GetOne() *RetUserRole {
	s := make([]RetUserRole, 0)
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

func (m *RetUserRoleQuery) Get() RetUserRoleCollect {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := make([]RetUserRole, 0)
	m.GetBuild().ModelType(&s).Get()
	return s
}

func (m *RetUserRoleQuery) clone() *RetUserRoleQuery {
	nm := NewRetUserRoleQuery()
	nm.SetBuild(m.GetBuild().Clone())
	return nm
}

func (m *RetUserRoleQuery) Count() int64 {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := RetUserRole{}
	return m.GetBuild().ModelType(&s).Count()
}

func (m *RetUserRoleQuery) sum(col string) float64 {
	s := RetUserRole{}
	return m.GetBuild().ModelType(&s).Sum(col)
}

func (m *RetUserRoleQuery) max(col string) float64 {
	s := RetUserRole{}
	return m.GetBuild().ModelType(&s).Max(col)
}

func (m *RetUserRoleQuery) DoneOperate() int64 {
	s := RetUserRole{}
	return m.GetBuild().ModelType(&s).DoneOperate()
}

func (m *RetUserRoleQuery) update(h *RetUserRole) int64 {
	s := RetUserRole{}
	return m.GetBuild().ModelType(&s).Update(h)
}

func (m *RetUserRoleQuery) Delete() int64 {
	s := RetUserRole{}
	return m.GetBuild().ModelType(&s).Delete()
}

func (m *RetUserRoleQuery) save(x *RetUserRole) {
	m.GetBuild().ModelType(x).Save()
}

func (m *RetUserRoleQuery) error() error {
	return m.GetBuild().ModelType(m).Error()
}

//支持分表
func (m *RetUserRoleQuery) insert(argu interface{}) {
	s := RetUserRole{}
	m.GetBuild().ModelType(&s).Insert(argu)
}

func (m *RetUserRoleQuery) increment(column string, amount int) int64 {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := RetUserRole{}
	return m.GetBuild().ModelType(&s).Increment(column, amount)
}

func (m *RetUserRoleQuery) Unscoped() *RetUserRoleQuery {
	m.GetBuild().Unscoped()
	return m
}

func (m *RetUserRoleQuery) Select(columns ...interface{}) *RetUserRoleQuery {
	m.GetBuild().Select(columns...)
	return m
}

func (m *RetUserRoleQuery) where(column string, args ...interface{}) *RetUserRoleQuery {
	m.GetBuild().Where(column, args...)
	return m
}

func (m *RetUserRoleQuery) whereMap(datas map[string]interface{}) *RetUserRoleQuery {
	m.GetBuild().WhereMap(datas)
	return m
}

func (m *RetUserRoleQuery) orWhere(column string, args ...interface{}) *RetUserRoleQuery {
	m.GetBuild().OrWhere(column, args...)
	return m
}

func (m *RetUserRoleQuery) orWhereCb(cb WhereCb) *RetUserRoleQuery {
	m.GetBuild().OrWhereCb(cb)
	return m
}

func (m *RetUserRoleQuery) whereCb(cb WhereCb) *RetUserRoleQuery {
	m.GetBuild().WhereCb(cb)
	return m
}

func (m *RetUserRoleQuery) whereRaw(sql string, bindings ...interface{}) *RetUserRoleQuery {
	m.GetBuild().WhereRaw(sql, bindings...)
	return m
}

func (m *RetUserRoleQuery) whereNull(column string) *RetUserRoleQuery {
	m.GetBuild().WhereNull(column)
	return m
}

func (m *RetUserRoleQuery) whereNotNull(column string) *RetUserRoleQuery {
	m.GetBuild().WhereNotNull(column)
	return m
}

func (m *RetUserRoleQuery) orWhereNull(column string) *RetUserRoleQuery {
	m.GetBuild().OrWhereNull(column)
	return m
}

func (m *RetUserRoleQuery) orWhereNotNull(column string) *RetUserRoleQuery {
	m.GetBuild().OrWhereNotNull(column)
	return m
}

func (m *RetUserRoleQuery) whereIn(column string, values interface{}) *RetUserRoleQuery {
	m.GetBuild().WhereIn(column, values)
	return m
}

func (m *RetUserRoleQuery) whereNotIn(column string, values interface{}) *RetUserRoleQuery {
	m.GetBuild().WhereNotIn(column, values)
	return m
}

func (m *RetUserRoleQuery) LeftJoin(table string, one string, args ...interface{}) *RetUserRoleQuery {
	m.GetBuild().LeftJoin(table, one, args...)
	return m
}

func (m *RetUserRoleQuery) RightJoin(table string, one string, args ...interface{}) *RetUserRoleQuery {
	m.GetBuild().RightJoin(table, one, args...)
	return m
}

func (m *RetUserRoleQuery) InnerJoin(table string, one string, args ...interface{}) *RetUserRoleQuery {
	m.GetBuild().InnerJoin(table, one, args...)
	return m
}

func (m *RetUserRoleQuery) OrderBy(column string, direction string) *RetUserRoleQuery {
	m.GetBuild().OrderBy(column, direction)
	return m
}

func (m *RetUserRoleQuery) OrderDescBy(column string) *RetUserRoleQuery {
	m.GetBuild().OrderDescBy(column)
	return m
}

func (m *RetUserRoleQuery) OrderAscBy(column string) *RetUserRoleQuery {
	m.GetBuild().OrderAscBy(column)
	return m
}

func (m *RetUserRoleQuery) Group(column string) *RetUserRoleQuery {
	m.GetBuild().Group(column)
	return m
}

func (m *RetUserRoleQuery) Inc(col string, num int) *RetUserRoleQuery {
	m.GetBuild().Inc(col, num)
	return m
}

func (m *RetUserRoleQuery) Set(col string, val interface{}) *RetUserRoleQuery {
	m.GetBuild().Set(col, val)
	return m
}

func (m *RetUserRoleQuery) Skip(lines int) *RetUserRoleQuery {
	m.GetBuild().Skip(lines)
	return m
}

func (m *RetUserRoleQuery) Offset(lines int) *RetUserRoleQuery {
	m.GetBuild().Offset(lines)
	return m
}

func (m *RetUserRoleQuery) Limit(lines int) *RetUserRoleQuery {
	m.GetBuild().Limit(lines)
	return m
}

func (m *RetUserRoleQuery) Take(lines int) *RetUserRoleQuery {
	m.GetBuild().Take(lines)
	return m
}

func (m *RetUserRoleQuery) ForUpdate() *RetUserRoleQuery {
	m.GetBuild().ForUpdate()
	return m
}

func (m *RetUserRoleQuery) SplitBy(val int64) *RetUserRoleQuery {
	m.GetBuild().SplitBy(val)
	return m
}

func (m *RetUserRoleQuery) ModelType(v interface{}) *RetUserRoleQuery {
	m.GetBuild().ModelType(v)
	return m
}

func (m *RetUserRoleQuery) DisablePanic() *RetUserRoleQuery {
	m.GetBuild().DisablePanic()
	return m
}

func (m *RetUserRoleQuery) OnWrite() *RetUserRoleQuery {
	m.GetBuild().OnWrite()
	return m
}

func (m *RetUserRoleQuery) kWheId(args ...interface{}) *RetUserRoleQuery {
	return m.where("id", args...)
}

func (m *RetUserRoleQuery) kWheUid(args ...interface{}) *RetUserRoleQuery {
	return m.where("uid", args...)
}

func (m *RetUserRoleQuery) kWheRoleId(args ...interface{}) *RetUserRoleQuery {
	return m.where("role_id", args...)
}

func (m *RetUserRoleQuery) kSetId(val interface{}) *RetUserRoleQuery {
	return m.Set("id", val)
}

func (m *RetUserRoleQuery) kSetUid(val interface{}) *RetUserRoleQuery {
	return m.Set("uid", val)
}

func (m *RetUserRoleQuery) kSetRoleId(val interface{}) *RetUserRoleQuery {
	return m.Set("role_id", val)
}

func (m *RetUserRoleQuery) kIncId(num int) *RetUserRoleQuery {
	return m.Inc("id", num)
}

func (m *RetUserRoleQuery) kIncUid(num int) *RetUserRoleQuery {
	return m.Inc("uid", num)
}

func (m *RetUserRoleQuery) kIncRoleId(num int) *RetUserRoleQuery {
	return m.Inc("role_id", num)
}

func (m *RetUserRoleQuery) kWheIdIn(values interface{}) *RetUserRoleQuery {
	return m.whereIn("id", values)
}

func (m *RetUserRoleQuery) kWheUidIn(values interface{}) *RetUserRoleQuery {
	return m.whereIn("uid", values)
}

func (m *RetUserRoleQuery) kWheRoleIdIn(values interface{}) *RetUserRoleQuery {
	return m.whereIn("role_id", values)
}

func (m *RetUserRoleQuery) kWheIdNotIn(values interface{}) *RetUserRoleQuery {
	return m.whereNotIn("id", values)
}

func (m *RetUserRoleQuery) kWheUidNotIn(values interface{}) *RetUserRoleQuery {
	return m.whereNotIn("uid", values)
}

func (m *RetUserRoleQuery) kWheRoleIdNotIn(values interface{}) *RetUserRoleQuery {
	return m.whereNotIn("role_id", values)
}

func Save(f *RetUserRole) *RetUserRole {
	NewRetUserRoleQuery().save(f)
	return f
}

func UpdateRetUserRoleAll(p *RetUserRole) int64 {
	build := NewRetUserRoleQuery()
	return build.update(p)
}

func FetchRetUserRoleAll() RetUserRoleCollect {
	build := NewRetUserRoleQuery()
	return build.Get()
}

func FetchRetUserRoleAllWithPageSize(page int, pageSize int) RetUserRoleCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewRetUserRoleQuery()
	return build.Skip(offset).Limit(pageSize).Get()
}

func CountRetUserRoleAll() int64 {
	build := NewRetUserRoleQuery()
	return build.Count()
}

// uniIndex
func UpdateRetUserRoleById(id int64, p *RetUserRole) int64 {
	build := NewRetUserRoleQuery()

	build.kWheId(id)

	return build.update(p)
}

func UpdateRetUserRoleByIds(id []int64, p *RetUserRole) int64 {
	build := NewRetUserRoleQuery()

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

func UpdateRetUserRoleByIdsWhatEver(id []int64, p *RetUserRole) int64 {
	build := NewRetUserRoleQuery()

	if len(id) == 1 {
		build.kWheId(id[0])
	} else {
		build.kWheIdIn(id)
	}

	return build.update(p)
}

func FetchByIds(id []int64) RetUserRoleCollect {
	build := NewRetUserRoleQuery()

	if len(id) == 0 {
		return make(RetUserRoleCollect, 0)
	}

	if len(id) == 1 {
		build.kWheId(id[0])
	} else {
		build.kWheIdIn(id)
	}

	return build.Get()
}

func FetchByIdsWithPageSize(id []int64, page int, pageSize int) RetUserRoleCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewRetUserRoleQuery()

	if len(id) == 0 {
		return make(RetUserRoleCollect, 0)
	}

	if len(id) == 1 {
		build.kWheId(id[0])
	} else {
		build.kWheIdIn(id)
	}

	return build.Skip(offset).Limit(pageSize).Get()
}

func CheckExistById(id int64) bool {
	build := NewRetUserRoleQuery()

	build.kWheId(id)

	cnt := build.Count()
	return cnt > 0
}

func GetOneById(id int64) *RetUserRole {
	build := NewRetUserRoleQuery()

	build.kWheId(id)

	return build.GetOne()
}

func DeleteById(id int64) int64 {
	build := NewRetUserRoleQuery()

	build.kWheId(id)

	return build.Delete()
}

func GetFirstById(id int64) *RetUserRole {
	build := NewRetUserRoleQuery()

	build.kWheId(id)

	return build.First()
}

// uniIndex
func UpdateRetUserRoleByUidAndRoleId(uid int64, roleId int64, p *RetUserRole) int64 {
	build := NewRetUserRoleQuery()

	build.kWheUid(uid)

	build.kWheRoleId(roleId)

	return build.update(p)
}

func CheckExistByUidAndRoleId(uid int64, roleId int64) bool {
	build := NewRetUserRoleQuery()

	build.kWheUid(uid)

	build.kWheRoleId(roleId)

	cnt := build.Count()
	return cnt > 0
}

func GetOneByUidAndRoleId(uid int64, roleId int64) *RetUserRole {
	build := NewRetUserRoleQuery()

	build.kWheUid(uid)

	build.kWheRoleId(roleId)

	return build.GetOne()
}

func DeleteByUidAndRoleId(uid int64, roleId int64) int64 {
	build := NewRetUserRoleQuery()

	build.kWheUid(uid)

	build.kWheRoleId(roleId)

	return build.Delete()
}

func GetFirstByUidAndRoleId(uid int64, roleId int64) *RetUserRole {
	build := NewRetUserRoleQuery()

	build.kWheUid(uid)

	build.kWheRoleId(roleId)

	return build.First()
}

// MultiIndex
func CheckExistByUid(uid int64) bool {
	build := NewRetUserRoleQuery()

	build.kWheUid(uid)

	cnt := build.Count()
	return cnt > 0
}

func GetFirstByUid(uid int64) *RetUserRole {
	build := NewRetUserRoleQuery()

	build.kWheUid(uid)

	return build.First()
}

func DeleteAllByUid(uid int64) int64 {
	build := NewRetUserRoleQuery()

	build.kWheUid(uid)

	return build.Delete()
}

func FetchByUid(uid int64) RetUserRoleCollect {
	build := NewRetUserRoleQuery()

	build.kWheUid(uid)

	return build.Get()
}

func FetchByUidWithPageSize(uid int64, page int, pageSize int) RetUserRoleCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewRetUserRoleQuery()
	return build.Skip(offset).Limit(pageSize).Get()
}

func GetOneByUid(uid int64) *RetUserRole {
	build := NewRetUserRoleQuery()

	build.kWheUid(uid)

	return build.GetOne()
}

func UpdateByUid(uid int64, p *RetUserRole) int64 {
	build := NewRetUserRoleQuery()

	build.kWheUid(uid)

	return build.update(p)
}

func UpdateRetUserRoleByUid(uid []int64, p *RetUserRole) int64 {
	build := NewRetUserRoleQuery()

	if len(uid) == 0 {
		return 0
	}

	if len(uid) == 1 {
		build.kWheUid(uid[0])
	} else {
		build.kWheUidIn(uid)
	}

	return build.update(p)
}

func FetchByUids(uid []int64) RetUserRoleCollect {
	build := NewRetUserRoleQuery()

	if len(uid) == 0 {
		return make(RetUserRoleCollect, 0)
	}

	if len(uid) == 1 {
		build.kWheUid(uid[0])
	} else {
		build.kWheUidIn(uid)
	}

	return build.Get()
}

func FetchByUidsWithPageSize(uid []int64, page int, pageSize int) RetUserRoleCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewRetUserRoleQuery()

	if len(uid) == 0 {
		return make(RetUserRoleCollect, 0)
	}

	if len(uid) == 1 {
		build.kWheUid(uid[0])
	} else {
		build.kWheUidIn(uid)
	}

	return build.Skip(offset).Limit(pageSize).Get()
}

func UpdateRetUserRoleByUidsWhatEver(uid []int64, p *RetUserRole) int64 {
	build := NewRetUserRoleQuery()

	if len(uid) == 1 {
		build.kWheUid(uid[0])
	} else {
		build.kWheUidIn(uid)
	}

	return build.update(p)
}

func CountRetUserRoleByUids(uid []int64) int64 {
	if len(uid) == 0 {
		return 0
	}
	build := NewRetUserRoleQuery()
	if len(uid) == 1 {
		build.kWheUid(uid[0])
	} else {
		build.kWheUidIn(uid)
	}

	return build.Count()
}
