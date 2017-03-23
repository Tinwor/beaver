package clients

import (
	"log"

	client "github.com/Tinwor/beaver/grpc/grpcauthorization"
	"google.golang.org/grpc"
)

type GrpcAuthorizationClient struct {
}

func NewGrpcAuthorizationClient(address string) client.AuthorizationClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		defer conn.Close()
		log.Fatal("did not connect: %v", err)
	}

	return client.NewAuthorizationClient(conn)

}
