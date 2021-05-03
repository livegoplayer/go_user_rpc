package main

import (
	"github.com/livegoplayer/go_helper/config"
	myconfig "github.com/livegoplayer/go_user_rpc/config"
	"github.com/livegoplayer/go_user_rpc/user"
)

func main() {

	config.LoadEnv()
	myconfig.InitLog()
	myconfig.InitDb()
	myconfig.InitRedis()

	user.GetUserRoleList(26)
}
