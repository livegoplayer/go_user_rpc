package ret_role_authority

import (
	"sort"
)

/*
此文件为自动生成，所有修改都不会生效
*/

func (s RetRoleAuthorityCollect) PluckId() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.Id
	}
	return list
}

func (s RetRoleAuthorityCollect) PluckRoleId() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.RoleId
	}
	return list
}

func (s RetRoleAuthorityCollect) PluckAuthorityId() []int64 {
	list := make([]int64, len(s))
	for i, v := range s {
		list[i] = v.AuthorityId
	}
	return list
}

func (s RetRoleAuthorityCollect) PluckUniId() []int64 {
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

func (s RetRoleAuthorityCollect) PluckUniRoleId() []int64 {
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

func (s RetRoleAuthorityCollect) PluckUniAuthorityId() []int64 {
	uniMap := make(map[int64]bool)
	list := make([]int64, 0)
	for _, v := range s {
		_, ok := uniMap[v.AuthorityId]
		if !ok {
			uniMap[v.AuthorityId] = true
			list = append(list, v.AuthorityId)
		}
	}
	return list
}

func (s RetRoleAuthorityCollect) GroupById() map[int64]RetRoleAuthorityCollect {
	m := make(map[int64]RetRoleAuthorityCollect)
	for _, v := range s {
		if _, ok := m[v.Id]; !ok {
			m[v.Id] = make(RetRoleAuthorityCollect, 0)
		}
		m[v.Id] = append(m[v.Id], v)
	}
	return m
}

func (s RetRoleAuthorityCollect) GroupByRoleId() map[int64]RetRoleAuthorityCollect {
	m := make(map[int64]RetRoleAuthorityCollect)
	for _, v := range s {
		if _, ok := m[v.RoleId]; !ok {
			m[v.RoleId] = make(RetRoleAuthorityCollect, 0)
		}
		m[v.RoleId] = append(m[v.RoleId], v)
	}
	return m
}

func (s RetRoleAuthorityCollect) GroupByAuthorityId() map[int64]RetRoleAuthorityCollect {
	m := make(map[int64]RetRoleAuthorityCollect)
	for _, v := range s {
		if _, ok := m[v.AuthorityId]; !ok {
			m[v.AuthorityId] = make(RetRoleAuthorityCollect, 0)
		}
		m[v.AuthorityId] = append(m[v.AuthorityId], v)
	}
	return m
}

func (s RetRoleAuthorityCollect) SortByFunc(f func(i, j int) bool) RetRoleAuthorityCollect {
	sort.SliceStable(s, f)
	return s
}

func (s RetRoleAuthorityCollect) Filter(f func(item RetRoleAuthority) bool) RetRoleAuthorityCollect {
	m := make(RetRoleAuthorityCollect, 0)
	for _, v := range s {
		if f(v) {
			m = append(m, v)
		}
	}
	return m
}

func (s RetRoleAuthorityCollect) KeyById() map[int64]RetRoleAuthority {
	m := make(map[int64]RetRoleAuthority)
	for _, v := range s {
		m[v.Id] = v
	}
	return m
}

func (s RetRoleAuthorityCollect) KeyByRoleId() map[int64]RetRoleAuthority {
	m := make(map[int64]RetRoleAuthority)
	for _, v := range s {
		m[v.RoleId] = v
	}
	return m
}

func (s RetRoleAuthorityCollect) KeyByAuthorityId() map[int64]RetRoleAuthority {
	m := make(map[int64]RetRoleAuthority)
	for _, v := range s {
		m[v.AuthorityId] = v
	}
	return m
}
