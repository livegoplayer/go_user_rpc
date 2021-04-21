package ret_role_authority

import "reflect"
import "github.com/livegoplayer/go_db_helper/mysql"

/**
@Collect
*/
type RetRoleAuthorityCollect []RetRoleAuthority

func NewRetRoleAuthorityQuery() *RetRoleAuthorityQuery {
	s := RetRoleAuthorityQuery{}
	s.SetBuild(mysql.NewBuild(&s))
	i, ok := reflect.ValueOf(&s).Interface().(BeforeHook)
	if ok {
		i.Before()
	}
	return &s
}

type _RetRoleAuthorityQueryColsStruct struct {
	Id          string
	RoleId      string
	AuthorityId string
}

func GetRetRoleAuthorityQueryCols() *_RetRoleAuthorityQueryColsStruct {
	return &_RetRoleAuthorityQueryColsStruct{
		Id:          "id",
		RoleId:      "role_id",
		AuthorityId: "authority_id",
	}
}

func (m *RetRoleAuthorityQuery) First() *RetRoleAuthority {
	s := make([]RetRoleAuthority, 0)
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	m.GetBuild().ModelType(&s).Limit(1).Get()
	if len(s) > 0 {
		return &s[0]
	}
	return &RetRoleAuthority{}
}

func (m *RetRoleAuthorityQuery) GetOne() *RetRoleAuthority {
	s := make([]RetRoleAuthority, 0)
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

func (m *RetRoleAuthorityQuery) Get() RetRoleAuthorityCollect {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := make([]RetRoleAuthority, 0)
	m.GetBuild().ModelType(&s).Get()
	return s
}

func (m *RetRoleAuthorityQuery) clone() *RetRoleAuthorityQuery {
	nm := NewRetRoleAuthorityQuery()
	nm.SetBuild(m.GetBuild().Clone())
	return nm
}

func (m *RetRoleAuthorityQuery) Count() int64 {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := RetRoleAuthority{}
	return m.GetBuild().ModelType(&s).Count()
}

func (m *RetRoleAuthorityQuery) sum(col string) float64 {
	s := RetRoleAuthority{}
	return m.GetBuild().ModelType(&s).Sum(col)
}

func (m *RetRoleAuthorityQuery) max(col string) float64 {
	s := RetRoleAuthority{}
	return m.GetBuild().ModelType(&s).Max(col)
}

func (m *RetRoleAuthorityQuery) DoneOperate() int64 {
	s := RetRoleAuthority{}
	return m.GetBuild().ModelType(&s).DoneOperate()
}

func (m *RetRoleAuthorityQuery) update(h *RetRoleAuthority) int64 {
	s := RetRoleAuthority{}
	return m.GetBuild().ModelType(&s).Update(h)
}

func (m *RetRoleAuthorityQuery) Delete() int64 {
	s := RetRoleAuthority{}
	return m.GetBuild().ModelType(&s).Delete()
}

func (m *RetRoleAuthorityQuery) save(x *RetRoleAuthority) {
	m.GetBuild().ModelType(x).Save()
}

func (m *RetRoleAuthorityQuery) error() error {
	return m.GetBuild().ModelType(m).Error()
}

//支持分表
func (m *RetRoleAuthorityQuery) insert(argu interface{}) {
	s := RetRoleAuthority{}
	m.GetBuild().ModelType(&s).Insert(argu)
}

func (m *RetRoleAuthorityQuery) increment(column string, amount int) int64 {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := RetRoleAuthority{}
	return m.GetBuild().ModelType(&s).Increment(column, amount)
}

func (m *RetRoleAuthorityQuery) Unscoped() *RetRoleAuthorityQuery {
	m.GetBuild().Unscoped()
	return m
}

func (m *RetRoleAuthorityQuery) Select(columns ...interface{}) *RetRoleAuthorityQuery {
	m.GetBuild().Select(columns...)
	return m
}

func (m *RetRoleAuthorityQuery) where(column string, args ...interface{}) *RetRoleAuthorityQuery {
	m.GetBuild().Where(column, args...)
	return m
}

func (m *RetRoleAuthorityQuery) whereMap(datas map[string]interface{}) *RetRoleAuthorityQuery {
	m.GetBuild().WhereMap(datas)
	return m
}

func (m *RetRoleAuthorityQuery) orWhere(column string, args ...interface{}) *RetRoleAuthorityQuery {
	m.GetBuild().OrWhere(column, args...)
	return m
}

func (m *RetRoleAuthorityQuery) orWhereCb(cb WhereCb) *RetRoleAuthorityQuery {
	m.GetBuild().OrWhereCb(cb)
	return m
}

func (m *RetRoleAuthorityQuery) whereCb(cb WhereCb) *RetRoleAuthorityQuery {
	m.GetBuild().WhereCb(cb)
	return m
}

func (m *RetRoleAuthorityQuery) whereRaw(sql string, bindings ...interface{}) *RetRoleAuthorityQuery {
	m.GetBuild().WhereRaw(sql, bindings...)
	return m
}

func (m *RetRoleAuthorityQuery) whereNull(column string) *RetRoleAuthorityQuery {
	m.GetBuild().WhereNull(column)
	return m
}

func (m *RetRoleAuthorityQuery) whereNotNull(column string) *RetRoleAuthorityQuery {
	m.GetBuild().WhereNotNull(column)
	return m
}

func (m *RetRoleAuthorityQuery) orWhereNull(column string) *RetRoleAuthorityQuery {
	m.GetBuild().OrWhereNull(column)
	return m
}

func (m *RetRoleAuthorityQuery) orWhereNotNull(column string) *RetRoleAuthorityQuery {
	m.GetBuild().OrWhereNotNull(column)
	return m
}

func (m *RetRoleAuthorityQuery) whereIn(column string, values interface{}) *RetRoleAuthorityQuery {
	m.GetBuild().WhereIn(column, values)
	return m
}

func (m *RetRoleAuthorityQuery) whereNotIn(column string, values interface{}) *RetRoleAuthorityQuery {
	m.GetBuild().WhereNotIn(column, values)
	return m
}

func (m *RetRoleAuthorityQuery) LeftJoin(table string, one string, args ...interface{}) *RetRoleAuthorityQuery {
	m.GetBuild().LeftJoin(table, one, args...)
	return m
}

func (m *RetRoleAuthorityQuery) RightJoin(table string, one string, args ...interface{}) *RetRoleAuthorityQuery {
	m.GetBuild().RightJoin(table, one, args...)
	return m
}

func (m *RetRoleAuthorityQuery) InnerJoin(table string, one string, args ...interface{}) *RetRoleAuthorityQuery {
	m.GetBuild().InnerJoin(table, one, args...)
	return m
}

func (m *RetRoleAuthorityQuery) OrderBy(column string, direction string) *RetRoleAuthorityQuery {
	m.GetBuild().OrderBy(column, direction)
	return m
}

func (m *RetRoleAuthorityQuery) OrderDescBy(column string) *RetRoleAuthorityQuery {
	m.GetBuild().OrderDescBy(column)
	return m
}

func (m *RetRoleAuthorityQuery) OrderAscBy(column string) *RetRoleAuthorityQuery {
	m.GetBuild().OrderAscBy(column)
	return m
}

func (m *RetRoleAuthorityQuery) Group(column string) *RetRoleAuthorityQuery {
	m.GetBuild().Group(column)
	return m
}

func (m *RetRoleAuthorityQuery) Inc(col string, num int) *RetRoleAuthorityQuery {
	m.GetBuild().Inc(col, num)
	return m
}

func (m *RetRoleAuthorityQuery) Set(col string, val interface{}) *RetRoleAuthorityQuery {
	m.GetBuild().Set(col, val)
	return m
}

func (m *RetRoleAuthorityQuery) Skip(lines int) *RetRoleAuthorityQuery {
	m.GetBuild().Skip(lines)
	return m
}

func (m *RetRoleAuthorityQuery) Offset(lines int) *RetRoleAuthorityQuery {
	m.GetBuild().Offset(lines)
	return m
}

func (m *RetRoleAuthorityQuery) Limit(lines int) *RetRoleAuthorityQuery {
	m.GetBuild().Limit(lines)
	return m
}

func (m *RetRoleAuthorityQuery) Take(lines int) *RetRoleAuthorityQuery {
	m.GetBuild().Take(lines)
	return m
}

func (m *RetRoleAuthorityQuery) ForUpdate() *RetRoleAuthorityQuery {
	m.GetBuild().ForUpdate()
	return m
}

func (m *RetRoleAuthorityQuery) SplitBy(val int64) *RetRoleAuthorityQuery {
	m.GetBuild().SplitBy(val)
	return m
}

func (m *RetRoleAuthorityQuery) ModelType(v interface{}) *RetRoleAuthorityQuery {
	m.GetBuild().ModelType(v)
	return m
}

func (m *RetRoleAuthorityQuery) DisablePanic() *RetRoleAuthorityQuery {
	m.GetBuild().DisablePanic()
	return m
}

func (m *RetRoleAuthorityQuery) OnWrite() *RetRoleAuthorityQuery {
	m.GetBuild().OnWrite()
	return m
}

func (m *RetRoleAuthorityQuery) kWheId(args ...interface{}) *RetRoleAuthorityQuery {
	return m.where("id", args...)
}

func (m *RetRoleAuthorityQuery) kWheRoleId(args ...interface{}) *RetRoleAuthorityQuery {
	return m.where("role_id", args...)
}

func (m *RetRoleAuthorityQuery) kWheAuthorityId(args ...interface{}) *RetRoleAuthorityQuery {
	return m.where("authority_id", args...)
}

func (m *RetRoleAuthorityQuery) kSetId(val interface{}) *RetRoleAuthorityQuery {
	return m.Set("id", val)
}

func (m *RetRoleAuthorityQuery) kSetRoleId(val interface{}) *RetRoleAuthorityQuery {
	return m.Set("role_id", val)
}

func (m *RetRoleAuthorityQuery) kSetAuthorityId(val interface{}) *RetRoleAuthorityQuery {
	return m.Set("authority_id", val)
}

func (m *RetRoleAuthorityQuery) kIncId(num int) *RetRoleAuthorityQuery {
	return m.Inc("id", num)
}

func (m *RetRoleAuthorityQuery) kIncRoleId(num int) *RetRoleAuthorityQuery {
	return m.Inc("role_id", num)
}

func (m *RetRoleAuthorityQuery) kIncAuthorityId(num int) *RetRoleAuthorityQuery {
	return m.Inc("authority_id", num)
}

func (m *RetRoleAuthorityQuery) kWheIdIn(values interface{}) *RetRoleAuthorityQuery {
	return m.whereIn("id", values)
}

func (m *RetRoleAuthorityQuery) kWheRoleIdIn(values interface{}) *RetRoleAuthorityQuery {
	return m.whereIn("role_id", values)
}

func (m *RetRoleAuthorityQuery) kWheAuthorityIdIn(values interface{}) *RetRoleAuthorityQuery {
	return m.whereIn("authority_id", values)
}

func (m *RetRoleAuthorityQuery) kWheIdNotIn(values interface{}) *RetRoleAuthorityQuery {
	return m.whereNotIn("id", values)
}

func (m *RetRoleAuthorityQuery) kWheRoleIdNotIn(values interface{}) *RetRoleAuthorityQuery {
	return m.whereNotIn("role_id", values)
}

func (m *RetRoleAuthorityQuery) kWheAuthorityIdNotIn(values interface{}) *RetRoleAuthorityQuery {
	return m.whereNotIn("authority_id", values)
}

func Save(f *RetRoleAuthority) *RetRoleAuthority {
	NewRetRoleAuthorityQuery().save(f)
	return f
}

func UpdateRetRoleAuthorityAll(p *RetRoleAuthority) int64 {
	build := NewRetRoleAuthorityQuery()
	return build.update(p)
}

func FetchRetRoleAuthorityAll() RetRoleAuthorityCollect {
	build := NewRetRoleAuthorityQuery()
	return build.Get()
}

func FetchRetRoleAuthorityAllWithPageSize(page int, pageSize int) RetRoleAuthorityCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewRetRoleAuthorityQuery()
	return build.Skip(offset).Limit(pageSize).Get()
}

func CountRetRoleAuthorityAll() int64 {
	build := NewRetRoleAuthorityQuery()
	return build.Count()
}

// uniIndex
func UpdateRetRoleAuthorityById(id int64, p *RetRoleAuthority) int64 {
	build := NewRetRoleAuthorityQuery()

	build.kWheId(id)

	return build.update(p)
}

func UpdateRetRoleAuthorityByIds(id []int64, p *RetRoleAuthority) int64 {
	build := NewRetRoleAuthorityQuery()

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

func UpdateRetRoleAuthorityByIdsWhatEver(id []int64, p *RetRoleAuthority) int64 {
	build := NewRetRoleAuthorityQuery()

	if len(id) == 1 {
		build.kWheId(id[0])
	} else {
		build.kWheIdIn(id)
	}

	return build.update(p)
}

func FetchByIds(id []int64) RetRoleAuthorityCollect {
	build := NewRetRoleAuthorityQuery()

	if len(id) == 0 {
		return make(RetRoleAuthorityCollect, 0)
	}

	if len(id) == 1 {
		build.kWheId(id[0])
	} else {
		build.kWheIdIn(id)
	}

	return build.Get()
}

func FetchByIdsWithPageSize(id []int64, page int, pageSize int) RetRoleAuthorityCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewRetRoleAuthorityQuery()

	if len(id) == 0 {
		return make(RetRoleAuthorityCollect, 0)
	}

	if len(id) == 1 {
		build.kWheId(id[0])
	} else {
		build.kWheIdIn(id)
	}

	return build.Skip(offset).Limit(pageSize).Get()
}

func CheckExistById(id int64) bool {
	build := NewRetRoleAuthorityQuery()

	build.kWheId(id)

	cnt := build.Count()
	return cnt > 0
}

func GetOneById(id int64) *RetRoleAuthority {
	build := NewRetRoleAuthorityQuery()

	build.kWheId(id)

	return build.GetOne()
}

func DeleteById(id int64) int64 {
	build := NewRetRoleAuthorityQuery()

	build.kWheId(id)

	return build.Delete()
}

func GetFirstById(id int64) *RetRoleAuthority {
	build := NewRetRoleAuthorityQuery()

	build.kWheId(id)

	return build.First()
}

// uniIndex
func UpdateRetRoleAuthorityByRoleIdAndAuthorityId(roleId int64, authorityId int64, p *RetRoleAuthority) int64 {
	build := NewRetRoleAuthorityQuery()

	build.kWheRoleId(roleId)

	build.kWheAuthorityId(authorityId)

	return build.update(p)
}

func CheckExistByRoleIdAndAuthorityId(roleId int64, authorityId int64) bool {
	build := NewRetRoleAuthorityQuery()

	build.kWheRoleId(roleId)

	build.kWheAuthorityId(authorityId)

	cnt := build.Count()
	return cnt > 0
}

func GetOneByRoleIdAndAuthorityId(roleId int64, authorityId int64) *RetRoleAuthority {
	build := NewRetRoleAuthorityQuery()

	build.kWheRoleId(roleId)

	build.kWheAuthorityId(authorityId)

	return build.GetOne()
}

func DeleteByRoleIdAndAuthorityId(roleId int64, authorityId int64) int64 {
	build := NewRetRoleAuthorityQuery()

	build.kWheRoleId(roleId)

	build.kWheAuthorityId(authorityId)

	return build.Delete()
}

func GetFirstByRoleIdAndAuthorityId(roleId int64, authorityId int64) *RetRoleAuthority {
	build := NewRetRoleAuthorityQuery()

	build.kWheRoleId(roleId)

	build.kWheAuthorityId(authorityId)

	return build.First()
}

// MultiIndex
func CheckExistByRoleId(roleId int64) bool {
	build := NewRetRoleAuthorityQuery()

	build.kWheRoleId(roleId)

	cnt := build.Count()
	return cnt > 0
}

func GetFirstByRoleId(roleId int64) *RetRoleAuthority {
	build := NewRetRoleAuthorityQuery()

	build.kWheRoleId(roleId)

	return build.First()
}

func DeleteAllByRoleId(roleId int64) int64 {
	build := NewRetRoleAuthorityQuery()

	build.kWheRoleId(roleId)

	return build.Delete()
}

func FetchByRoleId(roleId int64) RetRoleAuthorityCollect {
	build := NewRetRoleAuthorityQuery()

	build.kWheRoleId(roleId)

	return build.Get()
}

func FetchByRoleIdWithPageSize(roleId int64, page int, pageSize int) RetRoleAuthorityCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewRetRoleAuthorityQuery()
	return build.Skip(offset).Limit(pageSize).Get()
}

func GetOneByRoleId(roleId int64) *RetRoleAuthority {
	build := NewRetRoleAuthorityQuery()

	build.kWheRoleId(roleId)

	return build.GetOne()
}

func UpdateByRoleId(roleId int64, p *RetRoleAuthority) int64 {
	build := NewRetRoleAuthorityQuery()

	build.kWheRoleId(roleId)

	return build.update(p)
}

func UpdateRetRoleAuthorityByRoleId(roleId []int64, p *RetRoleAuthority) int64 {
	build := NewRetRoleAuthorityQuery()

	if len(roleId) == 0 {
		return 0
	}

	if len(roleId) == 1 {
		build.kWheRoleId(roleId[0])
	} else {
		build.kWheRoleIdIn(roleId)
	}

	return build.update(p)
}

func FetchByRoleIds(roleId []int64) RetRoleAuthorityCollect {
	build := NewRetRoleAuthorityQuery()

	if len(roleId) == 0 {
		return make(RetRoleAuthorityCollect, 0)
	}

	if len(roleId) == 1 {
		build.kWheRoleId(roleId[0])
	} else {
		build.kWheRoleIdIn(roleId)
	}

	return build.Get()
}

func FetchByRoleIdsWithPageSize(roleId []int64, page int, pageSize int) RetRoleAuthorityCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewRetRoleAuthorityQuery()

	if len(roleId) == 0 {
		return make(RetRoleAuthorityCollect, 0)
	}

	if len(roleId) == 1 {
		build.kWheRoleId(roleId[0])
	} else {
		build.kWheRoleIdIn(roleId)
	}

	return build.Skip(offset).Limit(pageSize).Get()
}

func UpdateRetRoleAuthorityByRoleIdsWhatEver(roleId []int64, p *RetRoleAuthority) int64 {
	build := NewRetRoleAuthorityQuery()

	if len(roleId) == 1 {
		build.kWheRoleId(roleId[0])
	} else {
		build.kWheRoleIdIn(roleId)
	}

	return build.update(p)
}

func CountRetRoleAuthorityByRoleIds(roleId []int64) int64 {
	if len(roleId) == 0 {
		return 0
	}
	build := NewRetRoleAuthorityQuery()
	if len(roleId) == 1 {
		build.kWheRoleId(roleId[0])
	} else {
		build.kWheRoleIdIn(roleId)
	}

	return build.Count()
}
