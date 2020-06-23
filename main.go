package go_user_rpc

import (
	"net"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/oklog/oklog/pkg/group"
	"go_user_rpc/user"
	userpb "go_user_rpc/user/grpc"
	realgrpc "google.golang.org/grpc"
)

const rateBucketNum = 20

var (
	logger log.Logger
	//fs       = flag.NewFlagSet("world", flag.ExitOnError)
	//httpAddr = fs.String("http-addr", ":8082", "HTTP listen address")
	//grpcAddr = fs.String("grpc-addr", ":8083", "gRPC listen address")
)

func Run() {
	//initHttpHandler()
	g := &group.Group{}

	initUserRpcHandler(g)
}

//func initHttpHandler() {
//	svc := user.UserService{}
//
//	r := user.MakeHTTPHandler(user.MakeUserEndpoints(svc))
//	err := http.ListenAndServe(":8080", r)
//	if err != nil {
//		fmt.Printf("server start error : " + err.Error())
//		return
//	}
//}

func initUserRpcHandler(g *group.Group) {

	grpcOpts := []grpctransport.ServerOption{
		grpctransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}

	grpcListener, err := net.Listen("tcp", ":8888")
	if err != nil {
		//todo log
	}

	endpoints := user.MakeUserEndpoints(&user.UserServiceServer{}, nil)
	baseServer := realgrpc.NewServer()

	g.Add(func() error {
		//这里是执行函数
		userpb.RegisterUserServer(baseServer, user.MakeGRPCHandler(*endpoints, grpcOpts...))
		return baseServer.Serve(grpcListener)
	}, func(error) {
		//这里是遇到错误的中断处理函数
		baseServer.Stop()
	})
}
