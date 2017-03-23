package clients

import (
	"log"

	client "github.com/Tinwor/beaver/grpc/grpcauth"
	"google.golang.org/grpc"
)

type GrpcAuthenticationClient struct {
}

func NewGrpcAuthenticationClient(address string) client.AuthenticationClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		defer conn.Close()
		log.Fatal("did not connect: %v", err)
	}

	return client.NewAuthenticationClient(conn)

}
