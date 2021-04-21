package authority

import "reflect"
import "github.com/livegoplayer/go_db_helper/mysql"

/**
@Collect
*/
type AuthorityCollect []Authority

func NewAuthorityQuery() *AuthorityQuery {
	s := AuthorityQuery{}
	s.SetBuild(mysql.NewBuild(&s))
	i, ok := reflect.ValueOf(&s).Interface().(BeforeHook)
	if ok {
		i.Before()
	}
	return &s
}

type _AuthorityQueryColsStruct struct {
	Id            string
	AuthorityName string
}

func GetAuthorityQueryCols() *_AuthorityQueryColsStruct {
	return &_AuthorityQueryColsStruct{
		Id:            "id",
		AuthorityName: "authority_name",
	}
}

func (m *AuthorityQuery) First() *Authority {
	s := make([]Authority, 0)
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	m.GetBuild().ModelType(&s).Limit(1).Get()
	if len(s) > 0 {
		return &s[0]
	}
	return &Authority{}
}

func (m *AuthorityQuery) GetOne() *Authority {
	s := make([]Authority, 0)
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

func (m *AuthorityQuery) Get() AuthorityCollect {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := make([]Authority, 0)
	m.GetBuild().ModelType(&s).Get()
	return s
}

func (m *AuthorityQuery) clone() *AuthorityQuery {
	nm := NewAuthorityQuery()
	nm.SetBuild(m.GetBuild().Clone())
	return nm
}

func (m *AuthorityQuery) Count() int64 {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := Authority{}
	return m.GetBuild().ModelType(&s).Count()
}

func (m *AuthorityQuery) sum(col string) float64 {
	s := Authority{}
	return m.GetBuild().ModelType(&s).Sum(col)
}

func (m *AuthorityQuery) max(col string) float64 {
	s := Authority{}
	return m.GetBuild().ModelType(&s).Max(col)
}

func (m *AuthorityQuery) DoneOperate() int64 {
	s := Authority{}
	return m.GetBuild().ModelType(&s).DoneOperate()
}

func (m *AuthorityQuery) update(h *Authority) int64 {
	s := Authority{}
	return m.GetBuild().ModelType(&s).Update(h)
}

func (m *AuthorityQuery) Delete() int64 {
	s := Authority{}
	return m.GetBuild().ModelType(&s).Delete()
}

func (m *AuthorityQuery) save(x *Authority) {
	m.GetBuild().ModelType(x).Save()
}

func (m *AuthorityQuery) error() error {
	return m.GetBuild().ModelType(m).Error()
}

//支持分表
func (m *AuthorityQuery) insert(argu interface{}) {
	s := Authority{}
	m.GetBuild().ModelType(&s).Insert(argu)
}

func (m *AuthorityQuery) increment(column string, amount int) int64 {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := Authority{}
	return m.GetBuild().ModelType(&s).Increment(column, amount)
}

func (m *AuthorityQuery) Unscoped() *AuthorityQuery {
	m.GetBuild().Unscoped()
	return m
}

func (m *AuthorityQuery) Select(columns ...interface{}) *AuthorityQuery {
	m.GetBuild().Select(columns...)
	return m
}

func (m *AuthorityQuery) where(column string, args ...interface{}) *AuthorityQuery {
	m.GetBuild().Where(column, args...)
	return m
}

func (m *AuthorityQuery) whereMap(datas map[string]interface{}) *AuthorityQuery {
	m.GetBuild().WhereMap(datas)
	return m
}

func (m *AuthorityQuery) orWhere(column string, args ...interface{}) *AuthorityQuery {
	m.GetBuild().OrWhere(column, args...)
	return m
}

func (m *AuthorityQuery) orWhereCb(cb WhereCb) *AuthorityQuery {
	m.GetBuild().OrWhereCb(cb)
	return m
}

func (m *AuthorityQuery) whereCb(cb WhereCb) *AuthorityQuery {
	m.GetBuild().WhereCb(cb)
	return m
}

func (m *AuthorityQuery) whereRaw(sql string, bindings ...interface{}) *AuthorityQuery {
	m.GetBuild().WhereRaw(sql, bindings...)
	return m
}

func (m *AuthorityQuery) whereNull(column string) *AuthorityQuery {
	m.GetBuild().WhereNull(column)
	return m
}

func (m *AuthorityQuery) whereNotNull(column string) *AuthorityQuery {
	m.GetBuild().WhereNotNull(column)
	return m
}

func (m *AuthorityQuery) orWhereNull(column string) *AuthorityQuery {
	m.GetBuild().OrWhereNull(column)
	return m
}

func (m *AuthorityQuery) orWhereNotNull(column string) *AuthorityQuery {
	m.GetBuild().OrWhereNotNull(column)
	return m
}

func (m *AuthorityQuery) whereIn(column string, values interface{}) *AuthorityQuery {
	m.GetBuild().WhereIn(column, values)
	return m
}

func (m *AuthorityQuery) whereNotIn(column string, values interface{}) *AuthorityQuery {
	m.GetBuild().WhereNotIn(column, values)
	return m
}

func (m *AuthorityQuery) LeftJoin(table string, one string, args ...interface{}) *AuthorityQuery {
	m.GetBuild().LeftJoin(table, one, args...)
	return m
}

func (m *AuthorityQuery) RightJoin(table string, one string, args ...interface{}) *AuthorityQuery {
	m.GetBuild().RightJoin(table, one, args...)
	return m
}

func (m *AuthorityQuery) InnerJoin(table string, one string, args ...interface{}) *AuthorityQuery {
	m.GetBuild().InnerJoin(table, one, args...)
	return m
}

func (m *AuthorityQuery) OrderBy(column string, direction string) *AuthorityQuery {
	m.GetBuild().OrderBy(column, direction)
	return m
}

func (m *AuthorityQuery) OrderDescBy(column string) *AuthorityQuery {
	m.GetBuild().OrderDescBy(column)
	return m
}

func (m *AuthorityQuery) OrderAscBy(column string) *AuthorityQuery {
	m.GetBuild().OrderAscBy(column)
	return m
}

func (m *AuthorityQuery) Group(column string) *AuthorityQuery {
	m.GetBuild().Group(column)
	return m
}

func (m *AuthorityQuery) Inc(col string, num int) *AuthorityQuery {
	m.GetBuild().Inc(col, num)
	return m
}

func (m *AuthorityQuery) Set(col string, val interface{}) *AuthorityQuery {
	m.GetBuild().Set(col, val)
	return m
}

func (m *AuthorityQuery) Skip(lines int) *AuthorityQuery {
	m.GetBuild().Skip(lines)
	return m
}

func (m *AuthorityQuery) Offset(lines int) *AuthorityQuery {
	m.GetBuild().Offset(lines)
	return m
}

func (m *AuthorityQuery) Limit(lines int) *AuthorityQuery {
	m.GetBuild().Limit(lines)
	return m
}

func (m *AuthorityQuery) Take(lines int) *AuthorityQuery {
	m.GetBuild().Take(lines)
	return m
}

func (m *AuthorityQuery) ForUpdate() *AuthorityQuery {
	m.GetBuild().ForUpdate()
	return m
}

func (m *AuthorityQuery) SplitBy(val int64) *AuthorityQuery {
	m.GetBuild().SplitBy(val)
	return m
}

func (m *AuthorityQuery) ModelType(v interface{}) *AuthorityQuery {
	m.GetBuild().ModelType(v)
	return m
}

func (m *AuthorityQuery) DisablePanic() *AuthorityQuery {
	m.GetBuild().DisablePanic()
	return m
}

func (m *AuthorityQuery) OnWrite() *AuthorityQuery {
	m.GetBuild().OnWrite()
	return m
}

func (m *AuthorityQuery) kWheId(args ...interface{}) *AuthorityQuery {
	return m.where("id", args...)
}

func (m *AuthorityQuery) kWheAuthorityName(args ...interface{}) *AuthorityQuery {
	return m.where("authority_name", args...)
}

func (m *AuthorityQuery) kSetId(val interface{}) *AuthorityQuery {
	return m.Set("id", val)
}

func (m *AuthorityQuery) kSetAuthorityName(val interface{}) *AuthorityQuery {
	return m.Set("authority_name", val)
}

func (m *AuthorityQuery) kIncId(num int) *AuthorityQuery {
	return m.Inc("id", num)
}

func (m *AuthorityQuery) kWheIdIn(values interface{}) *AuthorityQuery {
	return m.whereIn("id", values)
}

func (m *AuthorityQuery) kWheAuthorityNameIn(values interface{}) *AuthorityQuery {
	return m.whereIn("authority_name", values)
}

func (m *AuthorityQuery) kWheIdNotIn(values interface{}) *AuthorityQuery {
	return m.whereNotIn("id", values)
}

func (m *AuthorityQuery) kWheAuthorityNameNotIn(values interface{}) *AuthorityQuery {
	return m.whereNotIn("authority_name", values)
}

func Save(f *Authority) *Authority {
	NewAuthorityQuery().save(f)
	return f
}

func UpdateAuthorityAll(p *Authority) int64 {
	build := NewAuthorityQuery()
	return build.update(p)
}

func FetchAuthorityAll() AuthorityCollect {
	build := NewAuthorityQuery()
	return build.Get()
}

func FetchAuthorityAllWithPageSize(page int, pageSize int) AuthorityCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewAuthorityQuery()
	return build.Skip(offset).Limit(pageSize).Get()
}

func CountAuthorityAll() int64 {
	build := NewAuthorityQuery()
	return build.Count()
}

// uniIndex
func UpdateAuthorityById(id int64, p *Authority) int64 {
	build := NewAuthorityQuery()

	build.kWheId(id)

	return build.update(p)
}

func UpdateAuthorityByIds(id []int64, p *Authority) int64 {
	build := NewAuthorityQuery()

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

func UpdateAuthorityByIdsWhatEver(id []int64, p *Authority) int64 {
	build := NewAuthorityQuery()

	if len(id) == 1 {
		build.kWheId(id[0])
	} else {
		build.kWheIdIn(id)
	}

	return build.update(p)
}

func FetchByIds(id []int64) AuthorityCollect {
	build := NewAuthorityQuery()

	if len(id) == 0 {
		return make(AuthorityCollect, 0)
	}

	if len(id) == 1 {
		build.kWheId(id[0])
	} else {
		build.kWheIdIn(id)
	}

	return build.Get()
}

func FetchByIdsWithPageSize(id []int64, page int, pageSize int) AuthorityCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewAuthorityQuery()

	if len(id) == 0 {
		return make(AuthorityCollect, 0)
	}

	if len(id) == 1 {
		build.kWheId(id[0])
	} else {
		build.kWheIdIn(id)
	}

	return build.Skip(offset).Limit(pageSize).Get()
}

func CheckExistById(id int64) bool {
	build := NewAuthorityQuery()

	build.kWheId(id)

	cnt := build.Count()
	return cnt > 0
}

func GetOneById(id int64) *Authority {
	build := NewAuthorityQuery()

	build.kWheId(id)

	return build.GetOne()
}

func DeleteById(id int64) int64 {
	build := NewAuthorityQuery()

	build.kWheId(id)

	return build.Delete()
}

func GetFirstById(id int64) *Authority {
	build := NewAuthorityQuery()

	build.kWheId(id)

	return build.First()
}
