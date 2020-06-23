package user

import (
	"context"
	"encoding/json"
	"net"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	grpc_transport "github.com/go-kit/kit/transport/grpc"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

//一个断电对应一个服务
func MakeHTTPHandler(Endpoint endpoint.Endpoint, opts ...kithttp.ServerOption) http.Handler {
	r := mux.NewRouter()
	r.Handle("/get/{key}", kithttp.NewServer(
		Endpoint,
		decodeRequest,
		encodeHttpResponse, //暂且借用一下rpc封装好的方法
		opts...,
	)).Methods(http.MethodGet)
	return r
}

func MakeRpcHandle(Endpoint endpoint.Endpoint, opts ...grpc_transport.ServerOption) *grpc_transport.Server {

	bookListHandler := grpc_transport.NewServer(
		Endpoint,
		decodeRequestRpc,
		encodeResponseRpc,
		opts...,
	)
	//启动grpc服务
	serviceAddress := ":50052"
	ls, _ := net.Listen("tcp", serviceAddress)
	gs := grpc.NewServer()
	book.RegisterBookServiceServer(gs, bookServer)
	gs.Serve(ls)

	return bookListHandler
}

func decodeRequestRpc(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func decodeRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	requestType := vars["request_type"]

	requestBody := map[string]string{}
	for key, val := range vars {
		if key != "request_type" {
			requestBody[key] = val
		}
	}
	var userRequest = &UserRequest{
		RequestType: requestType,
		RequestBody: requestBody,
	}

	return userRequest, nil
}

func encodeHttpResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	res := &UserResponse{}
	ok := false

	if res, ok = response.(*UserResponse); !ok {
		res.Code = 1
		res.Msg = "response加密出错"
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(res)
}
