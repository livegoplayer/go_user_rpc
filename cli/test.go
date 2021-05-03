package main

import (
	"github.com/livegoplayer/go_helper/config"
	myconfig "github.com/livegoplayer/go_user_rpc/config"
	"github.com/livegoplayer/go_user_rpc/repository_template_model/user_detail"
)

func main() {

	config.LoadEnv()
	myconfig.InitLog()
	myconfig.InitDb()

	user_detail.GetUserAuthority(26)
}
