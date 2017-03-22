package server

import "github.com/dgrijalva/jwt-go"

const (
	secret = "MySuperSecret"
)

type AuthorizationServer struct {
}
type Claims struct {
	Guid string `json:"guid"`
	jwt.StandardClaims
}

func NewAuthorizationServer() *AuthorizationServer {
	return &AuthorizationServer{}
}
