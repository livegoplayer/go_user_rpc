package main

import (
	"github.com/livegoplayer/go_helper/config"
	myconfig "github.com/livegoplayer/go_user_rpc/config"
)

func main() {

	config.LoadEnv()
	myconfig.InitLog()
	myconfig.InitDb()
}
