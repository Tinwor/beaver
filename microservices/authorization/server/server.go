package server

import (
	"time"

	"log"

	auth "github.com/Tinwor/beaver/grpc/grpcauthorization"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/net/context"
)

const (
	secret = "MySuperSecret"
)

type AuthorizationServer struct {
	address string
}
type Claims struct {
	Guid string `json:"guid"`
	jwt.StandardClaims
}

func NewAuthorizationServer(address string) *AuthorizationServer {
	return &AuthorizationServer{address}
}

func (a *AuthorizationServer) NewToken(context context.Context, in *auth.TokenRequest) (*auth.TokenResponse, error) {
	expireToken := time.Now().Add(time.Hour * 1).Unix()
	claims := Claims{
		in.Guid,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			IssuedAt:  time.Now().Unix(),
			Issuer:    a.address,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Println("Error signin token: " + err.Error())
		return &auth.TokenResponse{Response: auth.AuthorizationStatusResponse_SERVER_ERROR}, err
	}
	return &auth.TokenResponse{
		Response: auth.AuthorizationStatusResponse_OK,
		Token:    signedToken}, nil
}
func (a *AuthorizationServer) RefreshToken(context.Context, *auth.TokenRefreshRequest) (*auth.TokenResponse, error) {
	return nil, nil
}
func (a *AuthorizationServer) AuthorizeUser(context.Context, *auth.AuthorizationRequest) (*auth.AuthorizationResponse, error) {
	return nil, nil
}
