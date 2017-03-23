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
	expireToken := time.Now().Add(time.Hour * 168).Unix()
	claims := Claims{
		in.Guid,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
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
func (a *AuthorizationServer) RefreshToken(ctx context.Context, in *auth.TokenRefreshRequest) (*auth.TokenResponse, error) {
	response, err := a.AuthorizeUser(context.Background(), &auth.AuthorizationRequest{
		Token: in.Token,
	})
	if response.Response == auth.AuthorizationStatusResponse_OK {
		r, err := a.NewToken(context.Background(), &auth.TokenRequest{
			Guid: response.Guid,
		})
		if r.Response == auth.AuthorizationStatusResponse_OK {
			return &auth.TokenResponse{
				Response: auth.AuthorizationStatusResponse_OK,
				Token:    r.Token,
			}, nil
		}
		log.Println("Error creating new token: " + err.Error())
		return &auth.TokenResponse{
			Response: r.Response,
		}, nil

	}
	log.Println("Error refreshing token: " + err.Error())
	return &auth.TokenResponse{
		Response: response.Response,
	}, nil

}
func (a *AuthorizationServer) AuthorizeUser(ctx context.Context, in *auth.AuthorizationRequest) (*auth.AuthorizationResponse, error) {
	token, err := jwt.ParseWithClaims(in.Token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return &auth.AuthorizationResponse{
				Response: auth.AuthorizationStatusResponse_ERROR,
			}, nil
		}
		return []byte(secret), nil
	})
	if err != nil {
		log.Println("Error parsing token: " + err.Error())
		return &auth.AuthorizationResponse{
			Response: auth.AuthorizationStatusResponse_SERVER_ERROR,
		}, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return &auth.AuthorizationResponse{
			Guid:     claims.Guid,
			Response: auth.AuthorizationStatusResponse_OK,
		}, nil
	}
	return &auth.AuthorizationResponse{
		Response: auth.AuthorizationStatusResponse_SERVER_ERROR,
	}, nil

}
