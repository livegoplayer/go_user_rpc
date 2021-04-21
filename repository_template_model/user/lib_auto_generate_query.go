package user

import "reflect"
import "github.com/livegoplayer/go_db_helper/mysql"

/**
@Collect
*/
type UserCollect []User

func NewUserQuery() *UserQuery {
	s := UserQuery{}
	s.SetBuild(mysql.NewBuild(&s))
	i, ok := reflect.ValueOf(&s).Interface().(BeforeHook)
	if ok {
		i.Before()
	}
	return &s
}

type _UserQueryColsStruct struct {
	Id          string
	Username    string
	Password    string
	AddDatetime string
	UptDatetime string
}

func GetUserQueryCols() *_UserQueryColsStruct {
	return &_UserQueryColsStruct{
		Id:          "id",
		Username:    "username",
		Password:    "password",
		AddDatetime: "add_datetime",
		UptDatetime: "upt_datetime",
	}
}

func (m *UserQuery) First() *User {
	s := make([]User, 0)
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	m.GetBuild().ModelType(&s).Limit(1).Get()
	if len(s) > 0 {
		return &s[0]
	}
	return &User{}
}

func (m *UserQuery) GetOne() *User {
	s := make([]User, 0)
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

func (m *UserQuery) Get() UserCollect {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := make([]User, 0)
	m.GetBuild().ModelType(&s).Get()
	return s
}

func (m *UserQuery) clone() *UserQuery {
	nm := NewUserQuery()
	nm.SetBuild(m.GetBuild().Clone())
	return nm
}

func (m *UserQuery) Count() int64 {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := User{}
	return m.GetBuild().ModelType(&s).Count()
}

func (m *UserQuery) sum(col string) float64 {
	s := User{}
	return m.GetBuild().ModelType(&s).Sum(col)
}

func (m *UserQuery) max(col string) float64 {
	s := User{}
	return m.GetBuild().ModelType(&s).Max(col)
}

func (m *UserQuery) DoneOperate() int64 {
	s := User{}
	return m.GetBuild().ModelType(&s).DoneOperate()
}

func (m *UserQuery) update(h *User) int64 {
	s := User{}
	return m.GetBuild().ModelType(&s).Update(h)
}

func (m *UserQuery) Delete() int64 {
	s := User{}
	return m.GetBuild().ModelType(&s).Delete()
}

func (m *UserQuery) save(x *User) {
	m.GetBuild().ModelType(x).Save()
}

func (m *UserQuery) error() error {
	return m.GetBuild().ModelType(m).Error()
}

//支持分表
func (m *UserQuery) insert(argu interface{}) {
	s := User{}
	m.GetBuild().ModelType(&s).Insert(argu)
}

func (m *UserQuery) increment(column string, amount int) int64 {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := User{}
	return m.GetBuild().ModelType(&s).Increment(column, amount)
}

func (m *UserQuery) Unscoped() *UserQuery {
	m.GetBuild().Unscoped()
	return m
}

func (m *UserQuery) Select(columns ...interface{}) *UserQuery {
	m.GetBuild().Select(columns...)
	return m
}

func (m *UserQuery) where(column string, args ...interface{}) *UserQuery {
	m.GetBuild().Where(column, args...)
	return m
}

func (m *UserQuery) whereMap(datas map[string]interface{}) *UserQuery {
	m.GetBuild().WhereMap(datas)
	return m
}

func (m *UserQuery) orWhere(column string, args ...interface{}) *UserQuery {
	m.GetBuild().OrWhere(column, args...)
	return m
}

func (m *UserQuery) orWhereCb(cb WhereCb) *UserQuery {
	m.GetBuild().OrWhereCb(cb)
	return m
}

func (m *UserQuery) whereCb(cb WhereCb) *UserQuery {
	m.GetBuild().WhereCb(cb)
	return m
}

func (m *UserQuery) whereRaw(sql string, bindings ...interface{}) *UserQuery {
	m.GetBuild().WhereRaw(sql, bindings...)
	return m
}

func (m *UserQuery) whereNull(column string) *UserQuery {
	m.GetBuild().WhereNull(column)
	return m
}

func (m *UserQuery) whereNotNull(column string) *UserQuery {
	m.GetBuild().WhereNotNull(column)
	return m
}

func (m *UserQuery) orWhereNull(column string) *UserQuery {
	m.GetBuild().OrWhereNull(column)
	return m
}

func (m *UserQuery) orWhereNotNull(column string) *UserQuery {
	m.GetBuild().OrWhereNotNull(column)
	return m
}

func (m *UserQuery) whereIn(column string, values interface{}) *UserQuery {
	m.GetBuild().WhereIn(column, values)
	return m
}

func (m *UserQuery) whereNotIn(column string, values interface{}) *UserQuery {
	m.GetBuild().WhereNotIn(column, values)
	return m
}

func (m *UserQuery) LeftJoin(table string, one string, args ...interface{}) *UserQuery {
	m.GetBuild().LeftJoin(table, one, args...)
	return m
}

func (m *UserQuery) RightJoin(table string, one string, args ...interface{}) *UserQuery {
	m.GetBuild().RightJoin(table, one, args...)
	return m
}

func (m *UserQuery) InnerJoin(table string, one string, args ...interface{}) *UserQuery {
	m.GetBuild().InnerJoin(table, one, args...)
	return m
}

func (m *UserQuery) OrderBy(column string, direction string) *UserQuery {
	m.GetBuild().OrderBy(column, direction)
	return m
}

func (m *UserQuery) OrderDescBy(column string) *UserQuery {
	m.GetBuild().OrderDescBy(column)
	return m
}

func (m *UserQuery) OrderAscBy(column string) *UserQuery {
	m.GetBuild().OrderAscBy(column)
	return m
}

func (m *UserQuery) Group(column string) *UserQuery {
	m.GetBuild().Group(column)
	return m
}

func (m *UserQuery) Inc(col string, num int) *UserQuery {
	m.GetBuild().Inc(col, num)
	return m
}

func (m *UserQuery) Set(col string, val interface{}) *UserQuery {
	m.GetBuild().Set(col, val)
	return m
}

func (m *UserQuery) Skip(lines int) *UserQuery {
	m.GetBuild().Skip(lines)
	return m
}

func (m *UserQuery) Offset(lines int) *UserQuery {
	m.GetBuild().Offset(lines)
	return m
}

func (m *UserQuery) Limit(lines int) *UserQuery {
	m.GetBuild().Limit(lines)
	return m
}

func (m *UserQuery) Take(lines int) *UserQuery {
	m.GetBuild().Take(lines)
	return m
}

func (m *UserQuery) ForUpdate() *UserQuery {
	m.GetBuild().ForUpdate()
	return m
}

func (m *UserQuery) SplitBy(val int64) *UserQuery {
	m.GetBuild().SplitBy(val)
	return m
}

func (m *UserQuery) ModelType(v interface{}) *UserQuery {
	m.GetBuild().ModelType(v)
	return m
}

func (m *UserQuery) DisablePanic() *UserQuery {
	m.GetBuild().DisablePanic()
	return m
}

func (m *UserQuery) OnWrite() *UserQuery {
	m.GetBuild().OnWrite()
	return m
}

func (m *UserQuery) kWheId(args ...interface{}) *UserQuery {
	return m.where("id", args...)
}

func (m *UserQuery) kWheUsername(args ...interface{}) *UserQuery {
	return m.where("username", args...)
}

func (m *UserQuery) kWhePassword(args ...interface{}) *UserQuery {
	return m.where("password", args...)
}

func (m *UserQuery) kWheAddDatetime(args ...interface{}) *UserQuery {
	return m.where("add_datetime", args...)
}

func (m *UserQuery) kWheUptDatetime(args ...interface{}) *UserQuery {
	return m.where("upt_datetime", args...)
}

func (m *UserQuery) kSetId(val interface{}) *UserQuery {
	return m.Set("id", val)
}

func (m *UserQuery) kSetUsername(val interface{}) *UserQuery {
	return m.Set("username", val)
}

func (m *UserQuery) kSetPassword(val interface{}) *UserQuery {
	return m.Set("password", val)
}

func (m *UserQuery) kSetAddDatetime(val interface{}) *UserQuery {
	return m.Set("add_datetime", val)
}

func (m *UserQuery) kSetUptDatetime(val interface{}) *UserQuery {
	return m.Set("upt_datetime", val)
}

func (m *UserQuery) kIncId(num int) *UserQuery {
	return m.Inc("id", num)
}

func (m *UserQuery) kWheIdIn(values interface{}) *UserQuery {
	return m.whereIn("id", values)
}

func (m *UserQuery) kWheUsernameIn(values interface{}) *UserQuery {
	return m.whereIn("username", values)
}

func (m *UserQuery) kWhePasswordIn(values interface{}) *UserQuery {
	return m.whereIn("password", values)
}

func (m *UserQuery) kWheAddDatetimeIn(values interface{}) *UserQuery {
	return m.whereIn("add_datetime", values)
}

func (m *UserQuery) kWheUptDatetimeIn(values interface{}) *UserQuery {
	return m.whereIn("upt_datetime", values)
}

func (m *UserQuery) kWheIdNotIn(values interface{}) *UserQuery {
	return m.whereNotIn("id", values)
}

func (m *UserQuery) kWheUsernameNotIn(values interface{}) *UserQuery {
	return m.whereNotIn("username", values)
}

func (m *UserQuery) kWhePasswordNotIn(values interface{}) *UserQuery {
	return m.whereNotIn("password", values)
}

func (m *UserQuery) kWheAddDatetimeNotIn(values interface{}) *UserQuery {
	return m.whereNotIn("add_datetime", values)
}

func (m *UserQuery) kWheUptDatetimeNotIn(values interface{}) *UserQuery {
	return m.whereNotIn("upt_datetime", values)
}

func Save(f *User) *User {
	NewUserQuery().save(f)
	return f
}

func UpdateUserAll(p *User) int64 {
	build := NewUserQuery()
	return build.update(p)
}

func FetchUserAll() UserCollect {
	build := NewUserQuery()
	return build.Get()
}

func FetchUserAllWithPageSize(page int, pageSize int) UserCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewUserQuery()
	return build.Skip(offset).Limit(pageSize).Get()
}

func CountUserAll() int64 {
	build := NewUserQuery()
	return build.Count()
}

// uniIndex
func UpdateUserById(id int64, p *User) int64 {
	build := NewUserQuery()

	build.kWheId(id)

	return build.update(p)
}

func UpdateUserByIds(id []int64, p *User) int64 {
	build := NewUserQuery()

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

func UpdateUserByIdsWhatEver(id []int64, p *User) int64 {
	build := NewUserQuery()

	if len(id) == 1 {
		build.kWheId(id[0])
	} else {
		build.kWheIdIn(id)
	}

	return build.update(p)
}

func FetchByIds(id []int64) UserCollect {
	build := NewUserQuery()

	if len(id) == 0 {
		return make(UserCollect, 0)
	}

	if len(id) == 1 {
		build.kWheId(id[0])
	} else {
		build.kWheIdIn(id)
	}

	return build.Get()
}

func FetchByIdsWithPageSize(id []int64, page int, pageSize int) UserCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewUserQuery()

	if len(id) == 0 {
		return make(UserCollect, 0)
	}

	if len(id) == 1 {
		build.kWheId(id[0])
	} else {
		build.kWheIdIn(id)
	}

	return build.Skip(offset).Limit(pageSize).Get()
}

func CheckExistById(id int64) bool {
	build := NewUserQuery()

	build.kWheId(id)

	cnt := build.Count()
	return cnt > 0
}

func GetOneById(id int64) *User {
	build := NewUserQuery()

	build.kWheId(id)

	return build.GetOne()
}

func DeleteById(id int64) int64 {
	build := NewUserQuery()

	build.kWheId(id)

	return build.Delete()
}

func GetFirstById(id int64) *User {
	build := NewUserQuery()

	build.kWheId(id)

	return build.First()
}

// uniIndex
func UpdateUserByUsernameAndPassword(username string, password string, p *User) int64 {
	build := NewUserQuery()

	build.kWheUsername(username)

	build.kWhePassword(password)

	return build.update(p)
}

func CheckExistByUsernameAndPassword(username string, password string) bool {
	build := NewUserQuery()

	build.kWheUsername(username)

	build.kWhePassword(password)

	cnt := build.Count()
	return cnt > 0
}

func GetOneByUsernameAndPassword(username string, password string) *User {
	build := NewUserQuery()

	build.kWheUsername(username)

	build.kWhePassword(password)

	return build.GetOne()
}

func DeleteByUsernameAndPassword(username string, password string) int64 {
	build := NewUserQuery()

	build.kWheUsername(username)

	build.kWhePassword(password)

	return build.Delete()
}

func GetFirstByUsernameAndPassword(username string, password string) *User {
	build := NewUserQuery()

	build.kWheUsername(username)

	build.kWhePassword(password)

	return build.First()
}

// MultiIndex
func CheckExistByUsername(username string) bool {
	build := NewUserQuery()

	build.kWheUsername(username)

	cnt := build.Count()
	return cnt > 0
}

func GetFirstByUsername(username string) *User {
	build := NewUserQuery()

	build.kWheUsername(username)

	return build.First()
}

func DeleteAllByUsername(username string) int64 {
	build := NewUserQuery()

	build.kWheUsername(username)

	return build.Delete()
}

func FetchByUsername(username string) UserCollect {
	build := NewUserQuery()

	build.kWheUsername(username)

	return build.Get()
}

func FetchByUsernameWithPageSize(username string, page int, pageSize int) UserCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewUserQuery()
	return build.Skip(offset).Limit(pageSize).Get()
}

func GetOneByUsername(username string) *User {
	build := NewUserQuery()

	build.kWheUsername(username)

	return build.GetOne()
}

func UpdateByUsername(username string, p *User) int64 {
	build := NewUserQuery()

	build.kWheUsername(username)

	return build.update(p)
}

func UpdateUserByUsername(username []string, p *User) int64 {
	build := NewUserQuery()

	if len(username) == 0 {
		return 0
	}

	if len(username) == 1 {
		build.kWheUsername(username[0])
	} else {
		build.kWheUsernameIn(username)
	}

	return build.update(p)
}

func FetchByUsernames(username []string) UserCollect {
	build := NewUserQuery()

	if len(username) == 0 {
		return make(UserCollect, 0)
	}

	if len(username) == 1 {
		build.kWheUsername(username[0])
	} else {
		build.kWheUsernameIn(username)
	}

	return build.Get()
}

func FetchByUsernamesWithPageSize(username []string, page int, pageSize int) UserCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewUserQuery()

	if len(username) == 0 {
		return make(UserCollect, 0)
	}

	if len(username) == 1 {
		build.kWheUsername(username[0])
	} else {
		build.kWheUsernameIn(username)
	}

	return build.Skip(offset).Limit(pageSize).Get()
}

func UpdateUserByUsernamesWhatEver(username []string, p *User) int64 {
	build := NewUserQuery()

	if len(username) == 1 {
		build.kWheUsername(username[0])
	} else {
		build.kWheUsernameIn(username)
	}

	return build.update(p)
}

func CountUserByUsernames(username []string) int64 {
	if len(username) == 0 {
		return 0
	}
	build := NewUserQuery()
	if len(username) == 1 {
		build.kWheUsername(username[0])
	} else {
		build.kWheUsernameIn(username)
	}

	return build.Count()
}
