package user

import (
	"google.golang.org/grpc"

	userpb "github.com/livegoplayer/go_user_rpc/user/grpc"
)

var userClientInstance userpb.UserClient

func initUserClient() {
	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err.Error())
	}

	userClientInstance = userpb.NewUserClient(conn)
}

func GetUserClient() userpb.UserClient {
	if userClientInstance != nil {
		return userClientInstance
	}

	initUserClient()

	return GetUserClient()
}
