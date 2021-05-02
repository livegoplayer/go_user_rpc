package user_detail

import (
	"github.com/livegoplayer/go_db_helper/mysql"
	"github.com/livegoplayer/go_user_rpc/repository_template_model/authority"
	"github.com/livegoplayer/go_user_rpc/repository_template_model/ret_role_authority"
	"github.com/livegoplayer/go_user_rpc/repository_template_model/ret_user_role"
	"github.com/livegoplayer/go_user_rpc/repository_template_model/role"
	"github.com/livegoplayer/go_user_rpc/repository_template_model/user"
)

func GetUserAuthority(uid int64) []UserDetail {
	UserDetails := make([]UserDetail, 0)
	clos := user.GetUserQueryCols()
	mysql.NewBuild(&UserDetails).Select("`us_user`.id as uid", "`us_user`.username as username", "`us_user`.add_datetime as add_datetime", "`us_user`.upt_datetime as upt_datetime", "`c`.id as role_id", "`c`.role_name as role_name", "`e`.id as authority_id", "`e`.authority_name as authority_name").
		LeftJoin(ret_user_role.RetUserRole{}.TableName()+" as a", "a.uid = us_user.id").
		LeftJoin(role.Role{}.TableName()+" as c", "a.role_id = c.id").
		LeftJoin(ret_role_authority.RetRoleAuthority{}.TableName()+" as d", "c.id = d.role_id").
		LeftJoin(authority.Authority{}.TableName()+" as e", "d.authority_id = e.id").Where(user.User{}.TableName()+"."+clos.Id, uid).
		Get()

	return UserDetails
}
