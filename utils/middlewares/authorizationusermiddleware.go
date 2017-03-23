package middlewares

import (
	"net/http"

	"golang.org/x/net/context"

	"log"

	"strings"

	client "github.com/Tinwor/beaver/grpc/grpcauthorization"
	"github.com/Tinwor/beaver/utils"
	"github.com/Tinwor/beaver/utils/clients"
	"github.com/julienschmidt/httprouter"
)

type AuthorizationUserMiddleware struct {
	authClient client.AuthorizationClient
}

const (
	header  = "Authorization"
	address = "localhost:50001"
)

func NewAuthorizationUserMiddleware() *AuthorizationUserMiddleware {
	cc := clients.NewGrpcAuthorizationClient(address)
	return &AuthorizationUserMiddleware{cc}
}
func (a *AuthorizationUserMiddleware) AuthorizeUser(page httprouter.Handle) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		log.Println("Entering middleware authorization")
		header := r.Header.Get(header)
		log.Println(len(strings.Split(header, ".")))
		if header == "" {
			log.Println("Request is null!")
			http.NotFound(w, r)
			return
		}

		response, err := a.authClient.AuthorizeUser(context.Background(), &client.AuthorizationRequest{
			Token: header,
		})

		switch response.Response {
		case client.AuthorizationStatusResponse_ERROR:
			w.WriteHeader(400)
			return
		case client.AuthorizationStatusResponse_SERVER_ERROR:
			log.Println("Server error: " + err.Error())
			w.WriteHeader(500)
			return
		case client.AuthorizationStatusResponse_OK:

			context := context.WithValue(r.Context(), utils.GuidKey, response.Guid)

			page(w, r.WithContext(context), ps)
		}
	})
}
