package user_operation_log

import "reflect"
import "github.com/livegoplayer/go_db_helper/mysql"

/**
@Collect
*/
type UserOperationLogCollect []UserOperationLog

func NewUserOperationLogQuery() *UserOperationLogQuery {
	s := UserOperationLogQuery{}
	s.SetBuild(mysql.NewBuild(&s))
	i, ok := reflect.ValueOf(&s).Interface().(BeforeHook)
	if ok {
		i.Before()
	}
	return &s
}

type _UserOperationLogQueryColsStruct struct {
	Id          string
	Cat         string
	Uid         string
	Message     string
	OperatorUid string
	AddDatetime string
}

func GetUserOperationLogQueryCols() *_UserOperationLogQueryColsStruct {
	return &_UserOperationLogQueryColsStruct{
		Id:          "id",
		Cat:         "cat",
		Uid:         "uid",
		Message:     "message",
		OperatorUid: "operator_uid",
		AddDatetime: "add_datetime",
	}
}

func (m *UserOperationLogQuery) First() *UserOperationLog {
	s := make([]UserOperationLog, 0)
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	m.GetBuild().ModelType(&s).Limit(1).Get()
	if len(s) > 0 {
		return &s[0]
	}
	return &UserOperationLog{}
}

func (m *UserOperationLogQuery) GetOne() *UserOperationLog {
	s := make([]UserOperationLog, 0)
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

func (m *UserOperationLogQuery) Get() UserOperationLogCollect {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := make([]UserOperationLog, 0)
	m.GetBuild().ModelType(&s).Get()
	return s
}

func (m *UserOperationLogQuery) clone() *UserOperationLogQuery {
	nm := NewUserOperationLogQuery()
	nm.SetBuild(m.GetBuild().Clone())
	return nm
}

func (m *UserOperationLogQuery) Count() int64 {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := UserOperationLog{}
	return m.GetBuild().ModelType(&s).Count()
}

func (m *UserOperationLogQuery) sum(col string) float64 {
	s := UserOperationLog{}
	return m.GetBuild().ModelType(&s).Sum(col)
}

func (m *UserOperationLogQuery) max(col string) float64 {
	s := UserOperationLog{}
	return m.GetBuild().ModelType(&s).Max(col)
}

func (m *UserOperationLogQuery) DoneOperate() int64 {
	s := UserOperationLog{}
	return m.GetBuild().ModelType(&s).DoneOperate()
}

func (m *UserOperationLogQuery) update(h *UserOperationLog) int64 {
	s := UserOperationLog{}
	return m.GetBuild().ModelType(&s).Update(h)
}

func (m *UserOperationLogQuery) Delete() int64 {
	s := UserOperationLog{}
	return m.GetBuild().ModelType(&s).Delete()
}

func (m *UserOperationLogQuery) save(x *UserOperationLog) {
	m.GetBuild().ModelType(x).Save()
}

func (m *UserOperationLogQuery) error() error {
	return m.GetBuild().ModelType(m).Error()
}

//支持分表
func (m *UserOperationLogQuery) insert(argu interface{}) {
	s := UserOperationLog{}
	m.GetBuild().ModelType(&s).Insert(argu)
}

func (m *UserOperationLogQuery) increment(column string, amount int) int64 {
	i, ok := reflect.ValueOf(m).Interface().(MustHook)
	if ok {
		i.Must()
	}
	s := UserOperationLog{}
	return m.GetBuild().ModelType(&s).Increment(column, amount)
}

func (m *UserOperationLogQuery) Unscoped() *UserOperationLogQuery {
	m.GetBuild().Unscoped()
	return m
}

func (m *UserOperationLogQuery) Select(columns ...interface{}) *UserOperationLogQuery {
	m.GetBuild().Select(columns...)
	return m
}

func (m *UserOperationLogQuery) where(column string, args ...interface{}) *UserOperationLogQuery {
	m.GetBuild().Where(column, args...)
	return m
}

func (m *UserOperationLogQuery) whereMap(datas map[string]interface{}) *UserOperationLogQuery {
	m.GetBuild().WhereMap(datas)
	return m
}

func (m *UserOperationLogQuery) orWhere(column string, args ...interface{}) *UserOperationLogQuery {
	m.GetBuild().OrWhere(column, args...)
	return m
}

func (m *UserOperationLogQuery) orWhereCb(cb WhereCb) *UserOperationLogQuery {
	m.GetBuild().OrWhereCb(cb)
	return m
}

func (m *UserOperationLogQuery) whereCb(cb WhereCb) *UserOperationLogQuery {
	m.GetBuild().WhereCb(cb)
	return m
}

func (m *UserOperationLogQuery) whereRaw(sql string, bindings ...interface{}) *UserOperationLogQuery {
	m.GetBuild().WhereRaw(sql, bindings...)
	return m
}

func (m *UserOperationLogQuery) whereNull(column string) *UserOperationLogQuery {
	m.GetBuild().WhereNull(column)
	return m
}

func (m *UserOperationLogQuery) whereNotNull(column string) *UserOperationLogQuery {
	m.GetBuild().WhereNotNull(column)
	return m
}

func (m *UserOperationLogQuery) orWhereNull(column string) *UserOperationLogQuery {
	m.GetBuild().OrWhereNull(column)
	return m
}

func (m *UserOperationLogQuery) orWhereNotNull(column string) *UserOperationLogQuery {
	m.GetBuild().OrWhereNotNull(column)
	return m
}

func (m *UserOperationLogQuery) whereIn(column string, values interface{}) *UserOperationLogQuery {
	m.GetBuild().WhereIn(column, values)
	return m
}

func (m *UserOperationLogQuery) whereNotIn(column string, values interface{}) *UserOperationLogQuery {
	m.GetBuild().WhereNotIn(column, values)
	return m
}

func (m *UserOperationLogQuery) LeftJoin(table string, one string, args ...interface{}) *UserOperationLogQuery {
	m.GetBuild().LeftJoin(table, one, args...)
	return m
}

func (m *UserOperationLogQuery) RightJoin(table string, one string, args ...interface{}) *UserOperationLogQuery {
	m.GetBuild().RightJoin(table, one, args...)
	return m
}

func (m *UserOperationLogQuery) InnerJoin(table string, one string, args ...interface{}) *UserOperationLogQuery {
	m.GetBuild().InnerJoin(table, one, args...)
	return m
}

func (m *UserOperationLogQuery) OrderBy(column string, direction string) *UserOperationLogQuery {
	m.GetBuild().OrderBy(column, direction)
	return m
}

func (m *UserOperationLogQuery) OrderDescBy(column string) *UserOperationLogQuery {
	m.GetBuild().OrderDescBy(column)
	return m
}

func (m *UserOperationLogQuery) OrderAscBy(column string) *UserOperationLogQuery {
	m.GetBuild().OrderAscBy(column)
	return m
}

func (m *UserOperationLogQuery) Group(column string) *UserOperationLogQuery {
	m.GetBuild().Group(column)
	return m
}

func (m *UserOperationLogQuery) Inc(col string, num int) *UserOperationLogQuery {
	m.GetBuild().Inc(col, num)
	return m
}

func (m *UserOperationLogQuery) Set(col string, val interface{}) *UserOperationLogQuery {
	m.GetBuild().Set(col, val)
	return m
}

func (m *UserOperationLogQuery) Skip(lines int) *UserOperationLogQuery {
	m.GetBuild().Skip(lines)
	return m
}

func (m *UserOperationLogQuery) Offset(lines int) *UserOperationLogQuery {
	m.GetBuild().Offset(lines)
	return m
}

func (m *UserOperationLogQuery) Limit(lines int) *UserOperationLogQuery {
	m.GetBuild().Limit(lines)
	return m
}

func (m *UserOperationLogQuery) Take(lines int) *UserOperationLogQuery {
	m.GetBuild().Take(lines)
	return m
}

func (m *UserOperationLogQuery) ForUpdate() *UserOperationLogQuery {
	m.GetBuild().ForUpdate()
	return m
}

func (m *UserOperationLogQuery) SplitBy(val int64) *UserOperationLogQuery {
	m.GetBuild().SplitBy(val)
	return m
}

func (m *UserOperationLogQuery) ModelType(v interface{}) *UserOperationLogQuery {
	m.GetBuild().ModelType(v)
	return m
}

func (m *UserOperationLogQuery) DisablePanic() *UserOperationLogQuery {
	m.GetBuild().DisablePanic()
	return m
}

func (m *UserOperationLogQuery) OnWrite() *UserOperationLogQuery {
	m.GetBuild().OnWrite()
	return m
}

func (m *UserOperationLogQuery) kWheId(args ...interface{}) *UserOperationLogQuery {
	return m.where("id", args...)
}

func (m *UserOperationLogQuery) kWheCat(args ...interface{}) *UserOperationLogQuery {
	return m.where("cat", args...)
}

func (m *UserOperationLogQuery) kWheUid(args ...interface{}) *UserOperationLogQuery {
	return m.where("uid", args...)
}

func (m *UserOperationLogQuery) kWheMessage(args ...interface{}) *UserOperationLogQuery {
	return m.where("message", args...)
}

func (m *UserOperationLogQuery) kWheOperatorUid(args ...interface{}) *UserOperationLogQuery {
	return m.where("operator_uid", args...)
}

func (m *UserOperationLogQuery) kWheAddDatetime(args ...interface{}) *UserOperationLogQuery {
	return m.where("add_datetime", args...)
}

func (m *UserOperationLogQuery) kSetId(val interface{}) *UserOperationLogQuery {
	return m.Set("id", val)
}

func (m *UserOperationLogQuery) kSetCat(val interface{}) *UserOperationLogQuery {
	return m.Set("cat", val)
}

func (m *UserOperationLogQuery) kSetUid(val interface{}) *UserOperationLogQuery {
	return m.Set("uid", val)
}

func (m *UserOperationLogQuery) kSetMessage(val interface{}) *UserOperationLogQuery {
	return m.Set("message", val)
}

func (m *UserOperationLogQuery) kSetOperatorUid(val interface{}) *UserOperationLogQuery {
	return m.Set("operator_uid", val)
}

func (m *UserOperationLogQuery) kSetAddDatetime(val interface{}) *UserOperationLogQuery {
	return m.Set("add_datetime", val)
}

func (m *UserOperationLogQuery) kIncId(num int) *UserOperationLogQuery {
	return m.Inc("id", num)
}

func (m *UserOperationLogQuery) kIncCat(num int) *UserOperationLogQuery {
	return m.Inc("cat", num)
}

func (m *UserOperationLogQuery) kIncUid(num int) *UserOperationLogQuery {
	return m.Inc("uid", num)
}

func (m *UserOperationLogQuery) kIncOperatorUid(num int) *UserOperationLogQuery {
	return m.Inc("operator_uid", num)
}

func (m *UserOperationLogQuery) kWheIdIn(values interface{}) *UserOperationLogQuery {
	return m.whereIn("id", values)
}

func (m *UserOperationLogQuery) kWheCatIn(values interface{}) *UserOperationLogQuery {
	return m.whereIn("cat", values)
}

func (m *UserOperationLogQuery) kWheUidIn(values interface{}) *UserOperationLogQuery {
	return m.whereIn("uid", values)
}

func (m *UserOperationLogQuery) kWheMessageIn(values interface{}) *UserOperationLogQuery {
	return m.whereIn("message", values)
}

func (m *UserOperationLogQuery) kWheOperatorUidIn(values interface{}) *UserOperationLogQuery {
	return m.whereIn("operator_uid", values)
}

func (m *UserOperationLogQuery) kWheAddDatetimeIn(values interface{}) *UserOperationLogQuery {
	return m.whereIn("add_datetime", values)
}

func (m *UserOperationLogQuery) kWheIdNotIn(values interface{}) *UserOperationLogQuery {
	return m.whereNotIn("id", values)
}

func (m *UserOperationLogQuery) kWheCatNotIn(values interface{}) *UserOperationLogQuery {
	return m.whereNotIn("cat", values)
}

func (m *UserOperationLogQuery) kWheUidNotIn(values interface{}) *UserOperationLogQuery {
	return m.whereNotIn("uid", values)
}

func (m *UserOperationLogQuery) kWheMessageNotIn(values interface{}) *UserOperationLogQuery {
	return m.whereNotIn("message", values)
}

func (m *UserOperationLogQuery) kWheOperatorUidNotIn(values interface{}) *UserOperationLogQuery {
	return m.whereNotIn("operator_uid", values)
}

func (m *UserOperationLogQuery) kWheAddDatetimeNotIn(values interface{}) *UserOperationLogQuery {
	return m.whereNotIn("add_datetime", values)
}

func Save(f *UserOperationLog) *UserOperationLog {
	NewUserOperationLogQuery().save(f)
	return f
}

func UpdateUserOperationLogAll(p *UserOperationLog) int64 {
	build := NewUserOperationLogQuery()
	return build.update(p)
}

func FetchUserOperationLogAll() UserOperationLogCollect {
	build := NewUserOperationLogQuery()
	return build.Get()
}

func FetchUserOperationLogAllWithPageSize(page int, pageSize int) UserOperationLogCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewUserOperationLogQuery()
	return build.Skip(offset).Limit(pageSize).Get()
}

func CountUserOperationLogAll() int64 {
	build := NewUserOperationLogQuery()
	return build.Count()
}

// uniIndex
func UpdateUserOperationLogById(id int64, p *UserOperationLog) int64 {
	build := NewUserOperationLogQuery()

	build.kWheId(id)

	return build.update(p)
}

func UpdateUserOperationLogByIds(id []int64, p *UserOperationLog) int64 {
	build := NewUserOperationLogQuery()

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

func UpdateUserOperationLogByIdsWhatEver(id []int64, p *UserOperationLog) int64 {
	build := NewUserOperationLogQuery()

	if len(id) == 1 {
		build.kWheId(id[0])
	} else {
		build.kWheIdIn(id)
	}

	return build.update(p)
}

func FetchByIds(id []int64) UserOperationLogCollect {
	build := NewUserOperationLogQuery()

	if len(id) == 0 {
		return make(UserOperationLogCollect, 0)
	}

	if len(id) == 1 {
		build.kWheId(id[0])
	} else {
		build.kWheIdIn(id)
	}

	return build.Get()
}

func FetchByIdsWithPageSize(id []int64, page int, pageSize int) UserOperationLogCollect {
	if page == 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	build := NewUserOperationLogQuery()

	if len(id) == 0 {
		return make(UserOperationLogCollect, 0)
	}

	if len(id) == 1 {
		build.kWheId(id[0])
	} else {
		build.kWheIdIn(id)
	}

	return build.Skip(offset).Limit(pageSize).Get()
}

func CheckExistById(id int64) bool {
	build := NewUserOperationLogQuery()

	build.kWheId(id)

	cnt := build.Count()
	return cnt > 0
}

func GetOneById(id int64) *UserOperationLog {
	build := NewUserOperationLogQuery()

	build.kWheId(id)

	return build.GetOne()
}

func DeleteById(id int64) int64 {
	build := NewUserOperationLogQuery()

	build.kWheId(id)

	return build.Delete()
}

func GetFirstById(id int64) *UserOperationLog {
	build := NewUserOperationLogQuery()

	build.kWheId(id)

	return build.First()
}
