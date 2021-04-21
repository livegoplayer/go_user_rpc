package ret_role_authority

import "github.com/livegoplayer/go_db_helper/mysql"

const PREFIX = "us_"

//
type RetRoleAuthority struct {
	Id          int64 `gorm:"column:id;PRIMARY_KEY" json:"id"`                                  //
	RoleId      int64 `gorm:"column:role_id;id_index_1_UNIQUE;id_index_2_MULTI" json:"role_id"` // 角色id
	AuthorityId int64 `gorm:"column:authority_id;id_index_1_UNIQUE" json:"authority_id"`        // 权限id
}

func (RetRoleAuthority) TableName() string {
	return PREFIX + "ret_role_authority"
}

type RetRoleAuthorityQuery struct {
	mysql.Query
}

func (RetRoleAuthority) Connect() string {
	return "user"
}
