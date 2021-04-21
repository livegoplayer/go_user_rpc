package authority

import (
	"sort"
)

/*
此文件为自动生成，所有修改都不会生效
*/

func (s AuthorityCollect) PluckId() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.Id
	}
	return list
}

func (s AuthorityCollect) PluckAuthorityName() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.AuthorityName
	}
	return list
}

func (s AuthorityCollect) PluckUniId() []int64 {
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

func (s AuthorityCollect) PluckUniAuthorityName() []string {
	uniMap := make(map[string]bool)
	list := make([]string, 0)
	for _, v := range s {
		_, ok := uniMap[v.AuthorityName]
		if !ok {
			uniMap[v.AuthorityName] = true
			list = append(list, v.AuthorityName)
		}
	}
	return list
}

func (s AuthorityCollect) GroupById() map[int64]AuthorityCollect {
	m := make(map[int64]AuthorityCollect)
	for _, v := range s {
		if _, ok := m[v.Id]; !ok {
			m[v.Id] = make(AuthorityCollect, 0)
		}
		m[v.Id] = append(m[v.Id], v)
	}
	return m
}

func (s AuthorityCollect) GroupByAuthorityName() map[string]AuthorityCollect {
	m := make(map[string]AuthorityCollect)
	for _, v := range s {
		if _, ok := m[v.AuthorityName]; !ok {
			m[v.AuthorityName] = make(AuthorityCollect, 0)
		}
		m[v.AuthorityName] = append(m[v.AuthorityName], v)
	}
	return m
}

func (s AuthorityCollect) SortByFunc(f func(i, j int) bool) AuthorityCollect {
	sort.SliceStable(s, f)
	return s
}

func (s AuthorityCollect) Filter(f func(item Authority) bool) AuthorityCollect {
	m := make(AuthorityCollect, 0)
	for _, v := range s {
		if f(v) {
			m = append(m, v)
		}
	}
	return m
}

func (s AuthorityCollect) KeyById() map[int64]Authority {
	m := make(map[int64]Authority)
	for _, v := range s {
		m[v.Id] = v
	}
	return m
}

func (s AuthorityCollect) KeyByAuthorityName() map[string]Authority {
	m := make(map[string]Authority)
	for _, v := range s {
		m[v.AuthorityName] = v
	}
	return m
}
