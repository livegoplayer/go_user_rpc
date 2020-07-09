package user

import (
	userpb "github.com/livegoplayer/go_user_rpc/user/grpc"
	"log"

	"google.golang.org/grpc"
)

var userClientInstance userpb.UserClient

func initUserClient() {
	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer func() {
		_ = conn.Close()
	}()

	userClientInstance = userpb.NewUserClient(conn)
}

func GetUserClient() userpb.UserClient {
	if userClientInstance != nil {
		return userClientInstance
	}

	initUserClient()

	return GetUserClient()
}
