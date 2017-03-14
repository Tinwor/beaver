package clients

import (
	"context"

	"log"

	client "github.com/Tinwor/beaver/grpc/data/grpcuser"
	"google.golang.org/grpc"
)

type GrpcUserClient struct {
	grpcClient client.UserClient
}

func NewGrpcUserClient(uri string) *GrpcUserClient {
	conn, err := grpc.Dial(uri, grpc.WithInsecure())
	if err != nil {
		defer conn.Close()
		log.Fatal(err)
	}
	ret := GrpcUserClient{
		grpcClient: client.NewUserClient(conn),
	}
	return &ret
}

func (g *GrpcUserClient) UserLogin(ctx context.Context, in *client.UserRequest, opts ...grpc.CallOption) (*client.LoginResponse, error) {
	return nil, nil
}
func (g *GrpcUserClient) UserNewRegister(ctx context.Context, in *client.UserRegister, opts ...grpc.CallOption) (*client.RegisterResponse, error) {
	return nil, nil
}
