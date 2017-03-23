package clients

import (
	"log"

	client "github.com/Tinwor/beaver/grpc/data/grpcuser"
	"google.golang.org/grpc"
)

type GrpcUserClient struct {
	grpcClient client.UserClient
}

func NewGrpcDataUserClient(address string) client.UserClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		defer conn.Close()
		log.Fatal("did not connect: %v", err)
	}

	return client.NewUserClient(conn)

}
