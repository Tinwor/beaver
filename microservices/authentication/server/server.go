package server

import "context"
import "github.com/Tinwor/beaver/grpc/grpcauth"

type Server struct {
}

func RegisterNewUser(context.Context, *grpcauth.RegisterRequest) (*grpcauth.AuthenticationResponse, error) {
	return nil, nil
}
func Login(context.Context, *grpcauth.LoginRequest) (*grpcauth.AuthenticationResponse, error) {
	return nil, nil
}
