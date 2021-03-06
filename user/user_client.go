package user

import (
	"google.golang.org/grpc"

	userpb "github.com/livegoplayer/go_user_rpc/user/user_grpc"
)

var userClientInstance userpb.UserClient

func InitUserClient(appHost string) {
	conn, err := grpc.Dial(appHost, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err.Error())
	}

	userClientInstance = userpb.NewUserClient(conn)
}

func GetUserClient() userpb.UserClient {
	if userClientInstance != nil {
		return userClientInstance
	}

	return GetUserClient()
}
