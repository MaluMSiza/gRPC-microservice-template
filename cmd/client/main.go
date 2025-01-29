package main

import (
	"context"
	pb "github.com/MaluMSiza/gRPC-microservice-template/gen/greet/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Falha ao conectar: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Mundo"})
	if err != nil {
		log.Fatalf("Erro ao chamar SayHello: %v", err)
	}

	log.Printf("Resposta do servidor: %s", res.Message)
}
