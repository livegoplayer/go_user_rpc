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
	res, _ := service.CheckUserStatus(context.Background(), &user_grpc.CheckUserStatusRequest{Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjI2LCJ1c2VybmFtZSI6InhqeXBsYXllciIsInVzZXJfcm9sZV9saXN0Ijp7IjEiOiJjb21tb25fdXNlciJ9LCJ1c2VyX2F1dGhvcml0eV9saXN0Ijp7IjEyIjoiZ2V0X3VzZXJfbGlzdCIsIjgiOiJnZXRfdXNlcl9yb2xlcyJ9LCJhZGRfZGF0ZXRpbWUiOiIyMDIxLTA1LTAzIDA0OjAzOjQ5Ljg3MiIsInVwZGF0ZV9kYXRldGltZSI6IjIwMjEtMDUtMDMgMDQ6MDM6NDkuODcyIiwiZXhwIjoxNjIwMzU2NTY5LCJpYXQiOjE2MjAyNzAxNjl9.vIs7djGjuYXYQDKRh5ETI2lMjCpJB2nDP6tb7Rf1N5s"})
	print(res)
}
