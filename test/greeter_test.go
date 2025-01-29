package test

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"testing"
	"time"

	pb "github.com/MaluMSiza/gRPC-microservice-template/gen/greet/v1"
	"github.com/MaluMSiza/gRPC-microservice-template/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestGreeterService(t *testing.T) {
	lis := bufconn.Listen(1024 * 1024)
	defer lis.Close()

	s := grpc.NewServer()
	defer s.Stop()

	service.RegisterService(s)

	go func() {
		if err := s.Serve(lis); err != nil {
			t.Errorf("Falha ao iniciar o servidor: %v", err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		"bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatalf("Falha ao conectar: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	t.Run("ValidRequest", func(t *testing.T) {
		res, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Test"})
		if err != nil {
			t.Fatalf("SayHello falhou: %v", err)
		}

		expected := "Olá, Test"
		if res.Message != expected {
			t.Errorf("Resposta inesperada:\nEsperado: %s\nObtido: %s", expected, res.Message)
		}
	})
}
