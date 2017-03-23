package main

import (
	"encoding/json"
	"log"
	"net/http"

	context "golang.org/x/net/context"

	"fmt"

	"github.com/Tinwor/beaver/grpc/grpcauth"
	authorization "github.com/Tinwor/beaver/grpc/grpcauthorization"
	"github.com/Tinwor/beaver/utils"
	"github.com/Tinwor/beaver/utils/clients"
	"github.com/Tinwor/beaver/utils/middlewares"
	"github.com/julienschmidt/httprouter"
)

const (
	port                 = ":4001"
	authClient           = ":51001"
	authorizationAddress = "localhost:50001"
)

var clientAuth grpcauth.AuthenticationClient
var authorizationClient authorization.AuthorizationClient

func main() {
	mw := middlewares.NewAuthorizationUserMiddleware()
	authorizationClient = clients.NewGrpcAuthorizationClient(authorizationAddress)
	clientAuth = clients.NewGrpcAuthenticationClient(authClient)
	router := httprouter.New()
	router.POST("/user/login", login)
	router.POST("/user/register", register)
	router.GET("/user/refresh/token", mw.AuthorizeUser(refreshToken))
	log.Fatal(http.ListenAndServe(port, router))
}

func login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var t grpcauth.AuthenticationRequest
	err := decoder.Decode(&t)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	response, err := clientAuth.Login(context.Background(), &t)
	switch response.Status {
	case grpcauth.AuthenticationStatusResponse_SERVER_ERROR:
		w.WriteHeader(500)
		return
	case grpcauth.AuthenticationStatusResponse_FAILURE:
		w.WriteHeader(402)
		fmt.Fprintf(w, "Username or password not correct")
	case grpcauth.AuthenticationStatusResponse_OK:
		w.WriteHeader(200)
		fmt.Fprintf(w, response.Token)
	}
}
func register(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var t grpcauth.RegisterRequest
	err := decoder.Decode(&t)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	response, err := clientAuth.RegisterNewUser(context.Background(), &t)
	switch response.Status {
	case grpcauth.AuthenticationStatusResponse_SERVER_ERROR:
		w.WriteHeader(500)
	case grpcauth.AuthenticationStatusResponse_FAILURE:
		w.WriteHeader(400)
		fmt.Fprintf(w, "Username or email already exist!")
	case grpcauth.AuthenticationStatusResponse_OK:
		w.WriteHeader(200)
		fmt.Fprintf(w, "Registration completed")
	}
}
func refreshToken(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	guid, ok := r.Context().Value(utils.GuidKey).(string)
	if !ok {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Impossible refresh token")
	}
	response, err := authorizationClient.RefreshToken(context.Background(), &authorization.TokenRefreshRequest{
		Token: guid,
	})
	if response.Response == authorization.AuthorizationStatusResponse_OK {
		w.WriteHeader(200)
		fmt.Fprintf(w, response.Token)
	} else {
		log.Println("Error refreshing token: " + err.Error())
		w.WriteHeader(500)
	}
}
