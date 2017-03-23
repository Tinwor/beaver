package server

import (
	"math/rand"

	context "golang.org/x/net/context"

	"log"

	"github.com/Tinwor/beaver/grpc/data/grpcuser"
	client "github.com/Tinwor/beaver/grpc/data/grpcuser"
	"github.com/Tinwor/beaver/grpc/grpcauth"
	authorization "github.com/Tinwor/beaver/grpc/grpcauthorization"
	"github.com/Tinwor/beaver/utils/clients"
	"github.com/twinj/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	userClient = "localhost:49001"
	authClient = "localhost:50001"
)

type AuthenticationServer struct {
	userClient client.UserClient
	authClient authorization.AuthorizationClient
}

func NewAuthenticationServer() *AuthenticationServer {
	var auth AuthenticationServer
	cc := clients.NewGrpcDataUserClient(userClient)
	ca := clients.NewGrpcAuthorizationClient(authClient)
	auth.userClient = cc
	auth.authClient = ca
	return &auth
}
func (a *AuthenticationServer) RegisterNewUser(context context.Context, in *grpcauth.RegisterRequest) (*grpcauth.AuthenticationResponse, error) {
	guid := uuid.NewV4().String()
	salt := randStringRunes(6)
	password, err := bcrypt.GenerateFromPassword([]byte(salt+in.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password: " + err.Error())
		return &grpcauth.AuthenticationResponse{
			Status: grpcauth.AuthenticationStatusResponse_SERVER_ERROR,
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
	case grpcuser.UserStatusResponse_SERVER_ERROR:
		return &grpcauth.AuthenticationResponse{
			Status: grpcauth.AuthenticationStatusResponse_SERVER_ERROR,
		}, err
	case grpcuser.UserStatusResponse_CREDENTIAL_EXIST:
		return &grpcauth.AuthenticationResponse{
			Status: grpcauth.AuthenticationStatusResponse_FAILURE,
		}, err
	case grpcuser.UserStatusResponse_OK:
		return &grpcauth.AuthenticationResponse{
			Status: grpcauth.AuthenticationStatusResponse_OK,
		}, nil
	default:
		return &grpcauth.AuthenticationResponse{
			Status: grpcauth.AuthenticationStatusResponse_SERVER_ERROR,
		}, err
	}
}
func (a *AuthenticationServer) Login(ctx context.Context, in *grpcauth.AuthenticationRequest) (*grpcauth.AuthenticationResponse, error) {

	if a.userClient == nil {
		log.Println("Error, nil reference")
	}

	response, err := a.userClient.UserLogin(context.Background(), &grpcuser.LoginRequest{
		Username: in.Username,
		Password: in.Password,
	})
	if response != nil {
		log.Println("There is something here!")
	} else {
		log.Println("Nothing to see")
	}
	switch response.Status {
	case grpcuser.UserStatusResponse_SERVER_ERROR:
		log.Println("Server error trying to log in")
		return &grpcauth.AuthenticationResponse{
			Status: grpcauth.AuthenticationStatusResponse_SERVER_ERROR,
		}, err
	case grpcuser.UserStatusResponse_FAILED:
		log.Println("Username or password not ok")
		return &grpcauth.AuthenticationResponse{
			Status: grpcauth.AuthenticationStatusResponse_FAILURE,
		}, nil
	case grpcuser.UserStatusResponse_OK:
		log.Println("Trying to create new Token")
		authResponse, err := a.authClient.NewToken(context.Background(), &authorization.TokenRequest{
			Guid: response.Token,
		})
		switch authResponse.Response {

		case authorization.AuthorizationStatusResponse_OK:
			return &grpcauth.AuthenticationResponse{
				Status: grpcauth.AuthenticationStatusResponse_OK,
				Token:  authResponse.Token,
			}, nil
		case authorization.AuthorizationStatusResponse_SERVER_ERROR:
		default:
			return &grpcauth.AuthenticationResponse{
				Status: grpcauth.AuthenticationStatusResponse_SERVER_ERROR,
			}, err

		}
	default:
		return &grpcauth.AuthenticationResponse{
			Status: grpcauth.AuthenticationStatusResponse_SERVER_ERROR,
		}, nil
	}
	return &grpcauth.AuthenticationResponse{}, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
