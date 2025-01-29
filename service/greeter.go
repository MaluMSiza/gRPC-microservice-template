package service

import (
	"context"
	pb "github.com/MaluMSiza/gRPC-microservice-template/gen/greet/v1"
	"google.golang.org/grpc"
)

type GreeterServer struct {
	pb.UnimplementedGreeterServer
}

func (s *GreeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: "Olá, " + req.Name,
	}, nil
}

func RegisterService(s *grpc.Server) {
	pb.RegisterGreeterServer(s, &GreeterServer{})
}
