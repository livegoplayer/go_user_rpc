package go_user_rpc

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"strconv"

	"github.com/go-kit/kit/log"

	"go_user_rpc/user"
)

const rateBucketNum = 20

var (
	logger log.Logger
	//fs       = flag.NewFlagSet("world", flag.ExitOnError)
	//httpAddr = fs.String("http-addr", ":8082", "HTTP listen address")
	//grpcAddr = fs.String("grpc-addr", ":8083", "gRPC listen address")
)

func Run() {
	initHttpHandler()
}

func initHttpHandler() {
	svc := user.UserService{}

	r := user.MakeHTTPHandler(user.MakeUserEndpoInt32(svc))
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Printf("server start error : " + err.Error())
		return
	}
}

func initRpcHandler() {

	// Publish our Handler methods
	rpc.Register(user.MakeRpcHandle(user.MakeUserEndpoInt32(svc)))

	// Create a TCP listener that will listen on `Port`
	listener, _ = net.Listen("tcp", ":8083")

	// Close the listener whenever we stop
	defer listener.Close()

	// Wait for incoming connections
	rpc.Accept(listener)
}
