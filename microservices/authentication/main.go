package main

import (
	"log"
	"net"

	pb "github.com/Tinwor/beaver/grpc/grpcauth"
	"github.com/Tinwor/beaver/microservices/authentication/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	address = ":51001"
)

func main() {
	server := server.NewAuthenticationServer()
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthenticationServer(s, server)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
