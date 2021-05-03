package main

import (
	"context"
	"github.com/livegoplayer/go_helper/config"
	myconfig "github.com/livegoplayer/go_user_rpc/config"
	"github.com/livegoplayer/go_user_rpc/user"
	"github.com/livegoplayer/go_user_rpc/user/user_grpc"
)

func main() {

	config.LoadEnv()
	myconfig.InitLog()
	myconfig.InitDb()
	myconfig.InitRedis()

	service := user.UserService{}
	res, _ := service.GetUserRoleList(context.Background(), &user_grpc.GetUserRoleListRequest{Uid: 26})
	print(res)
}
