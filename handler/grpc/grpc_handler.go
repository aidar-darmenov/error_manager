package grpc

import (
	"context"
	"flag"
	"fmt"
	"github.com/aidar-darmenov/error_manager/config"
	grpcGenerated "github.com/aidar-darmenov/error_manager/handler/grpc/generated"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	grpcGenerated.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *grpcGenerated.HelloRequest) (*grpcGenerated.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &grpcGenerated.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *grpcGenerated.HelloRequest) (*grpcGenerated.HelloReply, error) {
	return &grpcGenerated.HelloReply{Message: "Hello again " + in.GetName()}, nil
}

func InitRouter(c config.Configuration) *grpc.Server {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", c.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	grpcGenerated.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return s
}
