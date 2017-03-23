package server

import (
	"math/rand"

	"golang.org/x/net/context"

	"log"

	"github.com/Tinwor/beaver/grpc/data/grpcuser"
	client "github.com/Tinwor/beaver/grpc/data/grpcuser"
	"github.com/Tinwor/beaver/grpc/grpcauth"
	"github.com/Tinwor/beaver/utils/clients"
	"github.com/twinj/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	userClient = "localhost:49001"
)

type AuthenticationServer struct {
	userClient client.UserClient
}

func NewAuthenticationServer() *AuthenticationServer {
	var auth AuthenticationServer
	cc := clients.NewGrpcDataUserClient(userClient)
	auth.userClient = cc
	return &auth
}
func (a *AuthenticationServer) RegisterNewUser(context context.Context, in *grpcauth.RegisterRequest) (*grpcauth.AuthenticationResponse, error) {
	guid := uuid.NewV4().String()
	salt := randStringRunes(6)
	password, err := bcrypt.GenerateFromPassword([]byte(salt+in.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password: " + err.Error())
		return &grpcauth.AuthenticationResponse{
			Status: grpcauth.StatusResponse_SERVER_ERROR,
		}, err
	}
	response, err := a.userClient.NewUser(context, &grpcuser.RegisterUser{
		Guid:     guid,
		Username: in.Username,
		Email:    in.Email,
		Password: string(password),
		Salt:     salt,
	})
	switch response.Status {
	case grpcuser.StatusResponse_SERVER_ERROR:
		return &grpcauth.AuthenticationResponse{
			Status: grpcauth.StatusResponse_SERVER_ERROR,
		}, err
	case grpcuser.StatusResponse_CREDENTIAL_EXIST:
		return &grpcauth.AuthenticationResponse{
			Status: grpcauth.StatusResponse_FAILURE,
		}, err
	case grpcuser.StatusResponse_OK:
		return &grpcauth.AuthenticationResponse{
			Status: grpcauth.StatusResponse_OK,
		}, nil
	default:
		return &grpcauth.AuthenticationResponse{
			Status: grpcauth.StatusResponse_SERVER_ERROR,
		}, err
	}
}
func (a *AuthenticationServer) Login(context.Context, *grpcauth.LoginRequest) (*grpcauth.AuthenticationResponse, error) {
	return nil, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
