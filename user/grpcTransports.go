package user

import (
	"context"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/icowan/grpc-world/pkg/encode"
	ep "github.com/icowan/grpc-world/pkg/endpoint"
	"github.com/icowan/grpc-world/pkg/grpc/pb"
)

type userGrpcServer struct {
	get grpc.Handler
}

func (g *userGrpcServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.ServiceResponse, error) {
	_, rep, err := g.get.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ServiceResponse), nil
}

func MakeGRPCHandler(eps ep.Endpoints, opts ...grpc.ServerOption) pb.ServiceServer {
	return &grpcServer{
		get: grpc.NewServer(
			eps.GetEndpoint,
			decodeGetRequest,
			encodeResponse,
			opts...,
		),
	}
}
func decodeGetRequest(_ context.Context, r interface{}) (interface{}, error) {
	return ep.GetRequest{
		Key: r.(*pb.GetRequest).Key,
		Val: r.(*pb.GetRequest).Val,
	}, nil
}
func encodeResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(encode.Response)
	// ......省略
	return &pb.ServiceResponse{
		Success: resp.Success,
		Code:    int64(resp.Code),
	}, err
}
