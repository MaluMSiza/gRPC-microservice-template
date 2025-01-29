package main

import (
	"github.com/MaluMSiza/gRPC-microservice-template/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Falha ao escutar: %v", err)
	}

	grpcServer := grpc.NewServer()
	service.RegisterService(grpcServer)

	log.Println("Servidor gRPC rodando na porta :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Falha ao servir: %v", err)
	}
}
