package main

import (
	"log"
	"net"

	pb "github.com/Tinwor/beaver/grpc/grpcauthorization"
	"github.com/Tinwor/beaver/microservices/authorization/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	address = "localhost:50001"
)

func main() {
	server := server.NewAuthorizationServer(address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthorizationServer(s, server)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
