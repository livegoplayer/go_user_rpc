package role

import (
	"sort"
)

/*
此文件为自动生成，所有修改都不会生效
*/

func (s RoleCollect) PluckId() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.Id
	}
	return list
}

func (s RoleCollect) PluckRoleName() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.RoleName
	}
	return list
}

func (s RoleCollect) PluckRoleLevel() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.RoleLevel
	}
	return list
}

func (s RoleCollect) PluckAddDatetime() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.AddDatetime
	}
	return list
}

func (s RoleCollect) PluckUptDatetime() []string {
	list := make([]string, len(s))
	for i, v := range s {
		list[i] = v.UptDatetime
	}
	return list
}

func (s RoleCollect) PluckUniId() []int64 {
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

func (s RoleCollect) PluckUniRoleName() []string {
	uniMap := make(map[string]bool)
	list := make([]string, 0)
	for _, v := range s {
		_, ok := uniMap[v.RoleName]
		if !ok {
			uniMap[v.RoleName] = true
			list = append(list, v.RoleName)
		}
	}
	return list
}

func (s RoleCollect) PluckUniRoleLevel() []int64 {
	uniMap := make(map[int64]bool)
	list := make([]int64, 0)
	for _, v := range s {
		_, ok := uniMap[v.RoleLevel]
		if !ok {
			uniMap[v.RoleLevel] = true
			list = append(list, v.RoleLevel)
		}
	}
	return list
}

func (s RoleCollect) PluckUniAddDatetime() []string {
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

func (s RoleCollect) PluckUniUptDatetime() []string {
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

func (s RoleCollect) GroupById() map[int64]RoleCollect {
	m := make(map[int64]RoleCollect)
	for _, v := range s {
		if _, ok := m[v.Id]; !ok {
			m[v.Id] = make(RoleCollect, 0)
		}
		m[v.Id] = append(m[v.Id], v)
	}
	return m
}

func (s RoleCollect) GroupByRoleName() map[string]RoleCollect {
	m := make(map[string]RoleCollect)
	for _, v := range s {
		if _, ok := m[v.RoleName]; !ok {
			m[v.RoleName] = make(RoleCollect, 0)
		}
		m[v.RoleName] = append(m[v.RoleName], v)
	}
	return m
}

func (s RoleCollect) GroupByRoleLevel() map[int64]RoleCollect {
	m := make(map[int64]RoleCollect)
	for _, v := range s {
		if _, ok := m[v.RoleLevel]; !ok {
			m[v.RoleLevel] = make(RoleCollect, 0)
		}
		m[v.RoleLevel] = append(m[v.RoleLevel], v)
	}
	return m
}

func (s RoleCollect) GroupByAddDatetime() map[string]RoleCollect {
	m := make(map[string]RoleCollect)
	for _, v := range s {
		if _, ok := m[v.AddDatetime]; !ok {
			m[v.AddDatetime] = make(RoleCollect, 0)
		}
		m[v.AddDatetime] = append(m[v.AddDatetime], v)
	}
	return m
}

func (s RoleCollect) GroupByUptDatetime() map[string]RoleCollect {
	m := make(map[string]RoleCollect)
	for _, v := range s {
		if _, ok := m[v.UptDatetime]; !ok {
			m[v.UptDatetime] = make(RoleCollect, 0)
		}
		m[v.UptDatetime] = append(m[v.UptDatetime], v)
	}
	return m
}

func (s RoleCollect) SortByFunc(f func(i, j int) bool) RoleCollect {
	sort.SliceStable(s, f)
	return s
}

func (s RoleCollect) Filter(f func(item Role) bool) RoleCollect {
	m := make(RoleCollect, 0)
	for _, v := range s {
		if f(v) {
			m = append(m, v)
		}
	}
	return m
}

func (s RoleCollect) KeyById() map[int64]Role {
	m := make(map[int64]Role)
	for _, v := range s {
		m[v.Id] = v
	}
	return m
}

func (s RoleCollect) KeyByRoleName() map[string]Role {
	m := make(map[string]Role)
	for _, v := range s {
		m[v.RoleName] = v
	}
	return m
}

func (s RoleCollect) KeyByRoleLevel() map[int64]Role {
	m := make(map[int64]Role)
	for _, v := range s {
		m[v.RoleLevel] = v
	}
	return m
}

func (s RoleCollect) KeyByAddDatetime() map[string]Role {
	m := make(map[string]Role)
	for _, v := range s {
		m[v.AddDatetime] = v
	}
	return m
}

func (s RoleCollect) KeyByUptDatetime() map[string]Role {
	m := make(map[string]Role)
	for _, v := range s {
		m[v.UptDatetime] = v
	}
	return m
}
