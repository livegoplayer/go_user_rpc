package user_operation_log

import (
	"sort"
)

/*
此文件为自动生成，所有修改都不会生效
*/

func (s UserOperationLogCollect) PluckId() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.Id
	}
	return list
}

func (s UserOperationLogCollect) PluckCat() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.Cat
	}
	return list
}

func (s UserOperationLogCollect) PluckUid() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.Uid
	}
	return list
}

func (s UserOperationLogCollect) PluckMessage() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.Message
	}
	return list
}

func (s UserOperationLogCollect) PluckOperatorUid() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.OperatorUid
	}
	return list
}

func (s UserOperationLogCollect) PluckAddDatetime() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.AddDatetime
	}
	return list
}

func (s UserOperationLogCollect) PluckUniId() []int64 {
	uniMap := make(map[int64]bool)
	list := make([]int64, 0)
	for _, v := range s {
		_, ok := uniMap[v.Id]
		if !ok {
			uniMap[v.Id] = true
			list = append(list, v.Id)
		}
	}
	return list
}

func (s UserOperationLogCollect) PluckUniCat() []int64 {
	uniMap := make(map[int64]bool)
	list := make([]int64, 0)
	for _, v := range s {
		_, ok := uniMap[v.Cat]
		if !ok {
			uniMap[v.Cat] = true
			list = append(list, v.Cat)
		}
	}
	return list
}

func (s UserOperationLogCollect) PluckUniUid() []int64 {
	uniMap := make(map[int64]bool)
	list := make([]int64, 0)
	for _, v := range s {
		_, ok := uniMap[v.Uid]
		if !ok {
			uniMap[v.Uid] = true
			list = append(list, v.Uid)
		}
	}
	return list
}

func (s UserOperationLogCollect) PluckUniMessage() []string {
	uniMap := make(map[string]bool)
	list := make([]string, 0)
	for _, v := range s {
		_, ok := uniMap[v.Message]
		if !ok {
			uniMap[v.Message] = true
			list = append(list, v.Message)
		}
	}
	return list
}

func (s UserOperationLogCollect) PluckUniOperatorUid() []int64 {
	uniMap := make(map[int64]bool)
	list := make([]int64, 0)
	for _, v := range s {
		_, ok := uniMap[v.OperatorUid]
		if !ok {
			uniMap[v.OperatorUid] = true
			list = append(list, v.OperatorUid)
		}
	}
	return list
}

func (s UserOperationLogCollect) PluckUniAddDatetime() []string {
	uniMap := make(map[string]bool)
	list := make([]string, 0)
	for _, v := range s {
		_, ok := uniMap[v.AddDatetime]
		if !ok {
			uniMap[v.AddDatetime] = true
			list = append(list, v.AddDatetime)
		}
	}
	return list
}

func (s UserOperationLogCollect) GroupById() map[int64]UserOperationLogCollect {
	m := make(map[int64]UserOperationLogCollect)
	for _, v := range s {
		if _, ok := m[v.Id]; !ok {
			m[v.Id] = make(UserOperationLogCollect, 0)
		}
		m[v.Id] = append(m[v.Id], v)
	}
	return m
}

func (s UserOperationLogCollect) GroupByCat() map[int64]UserOperationLogCollect {
	m := make(map[int64]UserOperationLogCollect)
	for _, v := range s {
		if _, ok := m[v.Cat]; !ok {
			m[v.Cat] = make(UserOperationLogCollect, 0)
		}
		m[v.Cat] = append(m[v.Cat], v)
	}
	return m
}

func (s UserOperationLogCollect) GroupByUid() map[int64]UserOperationLogCollect {
	m := make(map[int64]UserOperationLogCollect)
	for _, v := range s {
		if _, ok := m[v.Uid]; !ok {
			m[v.Uid] = make(UserOperationLogCollect, 0)
		}
		m[v.Uid] = append(m[v.Uid], v)
	}
	return m
}

func (s UserOperationLogCollect) GroupByMessage() map[string]UserOperationLogCollect {
	m := make(map[string]UserOperationLogCollect)
	for _, v := range s {
		if _, ok := m[v.Message]; !ok {
			m[v.Message] = make(UserOperationLogCollect, 0)
		}
		m[v.Message] = append(m[v.Message], v)
	}
	return m
}

func (s UserOperationLogCollect) GroupByOperatorUid() map[int64]UserOperationLogCollect {
	m := make(map[int64]UserOperationLogCollect)
	for _, v := range s {
		if _, ok := m[v.OperatorUid]; !ok {
			m[v.OperatorUid] = make(UserOperationLogCollect, 0)
		}
		m[v.OperatorUid] = append(m[v.OperatorUid], v)
	}
	return m
}

func (s UserOperationLogCollect) GroupByAddDatetime() map[string]UserOperationLogCollect {
	m := make(map[string]UserOperationLogCollect)
	for _, v := range s {
		if _, ok := m[v.AddDatetime]; !ok {
			m[v.AddDatetime] = make(UserOperationLogCollect, 0)
		}
		m[v.AddDatetime] = append(m[v.AddDatetime], v)
	}
	return m
}

func (s UserOperationLogCollect) SortByFunc(f func(i, j int) bool) UserOperationLogCollect {
	sort.SliceStable(s, f)
	return s
}

func (s UserOperationLogCollect) Filter(f func(item UserOperationLog) bool) UserOperationLogCollect {
	m := make(UserOperationLogCollect, 0)
	for _, v := range s {
		if f(v) {
			m = append(m, v)
		}
	}
	return m
}

func (s UserOperationLogCollect) KeyById() map[int64]UserOperationLog {
	m := make(map[int64]UserOperationLog)
	for _, v := range s {
		m[v.Id] = v
	}
	return m
}

func (s UserOperationLogCollect) KeyByCat() map[int64]UserOperationLog {
	m := make(map[int64]UserOperationLog)
	for _, v := range s {
		m[v.Cat] = v
	}
	return m
}

func (s UserOperationLogCollect) KeyByUid() map[int64]UserOperationLog {
	m := make(map[int64]UserOperationLog)
	for _, v := range s {
		m[v.Uid] = v
	}
	return m
}

func (s UserOperationLogCollect) KeyByMessage() map[string]UserOperationLog {
	m := make(map[string]UserOperationLog)
	for _, v := range s {
		m[v.Message] = v
	}
	return m
}

func (s UserOperationLogCollect) KeyByOperatorUid() map[int64]UserOperationLog {
	m := make(map[int64]UserOperationLog)
	for _, v := range s {
		m[v.OperatorUid] = v
	}
	return m
}

func (s UserOperationLogCollect) KeyByAddDatetime() map[string]UserOperationLog {
	m := make(map[string]UserOperationLog)
	for _, v := range s {
		m[v.AddDatetime] = v
	}
	return m
}
