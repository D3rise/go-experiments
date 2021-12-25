package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "awesomeProject/test"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedTestServer
}

func (s *server) GetHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", request.GetName())
	return &pb.HelloResponse{
		Response: "Hello, " + request.GetName() + "!",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 3000))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTestServer(s, &server{})
	log.Printf("server listening on %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
