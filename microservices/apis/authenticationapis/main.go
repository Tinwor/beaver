package main

import (
	"encoding/json"
	"log"
	"net/http"

	context "golang.org/x/net/context"

	"fmt"

	"github.com/Tinwor/beaver/grpc/grpcauth"
	"github.com/Tinwor/beaver/utils/clients"
	"github.com/julienschmidt/httprouter"
)

const (
	port       = ":4001"
	authClient = ":51001"
)

var clientAuth grpcauth.AuthenticationClient

func main() {
	clientAuth = clients.NewGrpcAuthenticationClient(authClient)
	router := httprouter.New()
	router.POST("/user/login", login)
	router.POST("/user/register", register)
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
