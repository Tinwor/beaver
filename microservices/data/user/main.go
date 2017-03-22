package main

import (
	"log"
	"net"

	pb "github.com/Tinwor/beaver/grpc/data/grpcuser"
	"github.com/Tinwor/beaver/microservices/data/user/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":49001"
)

func main() {
	server := server.NewUserServer()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServer(s, server)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
