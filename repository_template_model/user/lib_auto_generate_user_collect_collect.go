package user

import (
	"sort"
)

/*
此文件为自动生成，所有修改都不会生效
*/

func (s UserCollect) PluckId() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.Id
	}
	return list
}

func (s UserCollect) PluckUsername() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.Username
	}
	return list
}

func (s UserCollect) PluckPassword() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.Password
	}
	return list
}

func (s UserCollect) PluckAddDatetime() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.AddDatetime
	}
	return list
}

func (s UserCollect) PluckUptDatetime() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.UptDatetime
	}
	return list
}

func (s UserCollect) PluckUniId() []int64 {
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

func (s UserCollect) PluckUniUsername() []string {
	uniMap := make(map[string]bool)
	list := make([]string, 0)
	for _, v := range s {
		_, ok := uniMap[v.Username]
		if !ok {
			uniMap[v.Username] = true
			list = append(list, v.Username)
		}
	}
	return list
}

func (s UserCollect) PluckUniPassword() []string {
	uniMap := make(map[string]bool)
	list := make([]string, 0)
	for _, v := range s {
		_, ok := uniMap[v.Password]
		if !ok {
			uniMap[v.Password] = true
			list = append(list, v.Password)
		}
	}
	return list
}

func (s UserCollect) PluckUniAddDatetime() []string {
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

func (s UserCollect) PluckUniUptDatetime() []string {
	uniMap := make(map[string]bool)
	list := make([]string, 0)
	for _, v := range s {
		_, ok := uniMap[v.UptDatetime]
		if !ok {
			uniMap[v.UptDatetime] = true
			list = append(list, v.UptDatetime)
		}
	}
	return list
}

func (s UserCollect) GroupById() map[int64]UserCollect {
	m := make(map[int64]UserCollect)
	for _, v := range s {
		if _, ok := m[v.Id]; !ok {
			m[v.Id] = make(UserCollect, 0)
		}
		m[v.Id] = append(m[v.Id], v)
	}
	return m
}

func (s UserCollect) GroupByUsername() map[string]UserCollect {
	m := make(map[string]UserCollect)
	for _, v := range s {
		if _, ok := m[v.Username]; !ok {
			m[v.Username] = make(UserCollect, 0)
		}
		m[v.Username] = append(m[v.Username], v)
	}
	return m
}

func (s UserCollect) GroupByPassword() map[string]UserCollect {
	m := make(map[string]UserCollect)
	for _, v := range s {
		if _, ok := m[v.Password]; !ok {
			m[v.Password] = make(UserCollect, 0)
		}
		m[v.Password] = append(m[v.Password], v)
	}
	return m
}

func (s UserCollect) GroupByAddDatetime() map[string]UserCollect {
	m := make(map[string]UserCollect)
	for _, v := range s {
		if _, ok := m[v.AddDatetime]; !ok {
			m[v.AddDatetime] = make(UserCollect, 0)
		}
		m[v.AddDatetime] = append(m[v.AddDatetime], v)
	}
	return m
}

func (s UserCollect) GroupByUptDatetime() map[string]UserCollect {
	m := make(map[string]UserCollect)
	for _, v := range s {
		if _, ok := m[v.UptDatetime]; !ok {
			m[v.UptDatetime] = make(UserCollect, 0)
		}
		m[v.UptDatetime] = append(m[v.UptDatetime], v)
	}
	return m
}

func (s UserCollect) SortByFunc(f func(i, j int) bool) UserCollect {
	sort.SliceStable(s, f)
	return s
}

func (s UserCollect) Filter(f func(item User) bool) UserCollect {
	m := make(UserCollect, 0)
	for _, v := range s {
		if f(v) {
			m = append(m, v)
		}
	}
	return m
}

func (s UserCollect) KeyById() map[int64]User {
	m := make(map[int64]User)
	for _, v := range s {
		m[v.Id] = v
	}
	return m
}

func (s UserCollect) KeyByUsername() map[string]User {
	m := make(map[string]User)
	for _, v := range s {
		m[v.Username] = v
	}
	return m
}

func (s UserCollect) KeyByPassword() map[string]User {
	m := make(map[string]User)
	for _, v := range s {
		m[v.Password] = v
	}
	return m
}

func (s UserCollect) KeyByAddDatetime() map[string]User {
	m := make(map[string]User)
	for _, v := range s {
		m[v.AddDatetime] = v
	}
	return m
}

func (s UserCollect) KeyByUptDatetime() map[string]User {
	m := make(map[string]User)
	for _, v := range s {
		m[v.UptDatetime] = v
	}
	return m
}
