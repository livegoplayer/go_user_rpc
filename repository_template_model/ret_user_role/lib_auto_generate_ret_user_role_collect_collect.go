package ret_user_role

import (
	"sort"
)

/*
此文件为自动生成，所有修改都不会生效
*/

func (s RetUserRoleCollect) PluckId() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.Id
	}
	return list
}

func (s RetUserRoleCollect) PluckUid() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.Uid
	}
	return list
}

func (s RetUserRoleCollect) PluckRoleId() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.RoleId
	}
	return list
}

func (s RetUserRoleCollect) PluckUniId() []int64 {
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

func (s RetUserRoleCollect) PluckUniUid() []int64 {
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

func (s RetUserRoleCollect) PluckUniRoleId() []int64 {
	uniMap := make(map[int64]bool)
	list := make([]int64, 0)
	for _, v := range s {
		_, ok := uniMap[v.RoleId]
		if !ok {
			uniMap[v.RoleId] = true
			list = append(list, v.RoleId)
		}
	}
	return list
}

func (s RetUserRoleCollect) GroupById() map[int64]RetUserRoleCollect {
	m := make(map[int64]RetUserRoleCollect)
	for _, v := range s {
		if _, ok := m[v.Id]; !ok {
			m[v.Id] = make(RetUserRoleCollect, 0)
		}
		m[v.Id] = append(m[v.Id], v)
	}
	return m
}

func (s RetUserRoleCollect) GroupByUid() map[int64]RetUserRoleCollect {
	m := make(map[int64]RetUserRoleCollect)
	for _, v := range s {
		if _, ok := m[v.Uid]; !ok {
			m[v.Uid] = make(RetUserRoleCollect, 0)
		}
		m[v.Uid] = append(m[v.Uid], v)
	}
	return m
}

func (s RetUserRoleCollect) GroupByRoleId() map[int64]RetUserRoleCollect {
	m := make(map[int64]RetUserRoleCollect)
	for _, v := range s {
		if _, ok := m[v.RoleId]; !ok {
			m[v.RoleId] = make(RetUserRoleCollect, 0)
		}
		m[v.RoleId] = append(m[v.RoleId], v)
	}
	return m
}

func (s RetUserRoleCollect) SortByFunc(f func(i, j int) bool) RetUserRoleCollect {
	sort.SliceStable(s, f)
	return s
}

func (s RetUserRoleCollect) Filter(f func(item RetUserRole) bool) RetUserRoleCollect {
	m := make(RetUserRoleCollect, 0)
	for _, v := range s {
		if f(v) {
			m = append(m, v)
		}
	}
	return m
}

func (s RetUserRoleCollect) KeyById() map[int64]RetUserRole {
	m := make(map[int64]RetUserRole)
	for _, v := range s {
		m[v.Id] = v
	}
	return m
}

func (s RetUserRoleCollect) KeyByUid() map[int64]RetUserRole {
	m := make(map[int64]RetUserRole)
	for _, v := range s {
		m[v.Uid] = v
	}
	return m
}

func (s RetUserRoleCollect) KeyByRoleId() map[int64]RetUserRole {
	m := make(map[int64]RetUserRole)
	for _, v := range s {
		m[v.RoleId] = v
	}
	return m
}
