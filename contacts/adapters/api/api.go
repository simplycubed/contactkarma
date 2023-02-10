package api

import (
	"log"
	"net/http"

	"github.com/simplycubed/contactkarma/contacts/adapters/firebase"
	"github.com/simplycubed/contactkarma/contacts/conf"
	"github.com/simplycubed/contactkarma/contacts/gen/restapi"
	"github.com/simplycubed/contactkarma/contacts/gen/restapi/operations"
	"github.com/rs/cors"

	"github.com/go-openapi/loads"
)

func Create() *operations.ContactsAPI {
	swaggerSpec, err := loads.Analyzed(
		restapi.FlatSwaggerJSON,
		"2.0",
	)
	api := operations.NewContactsAPI(swaggerSpec)

	if err != nil {
		log.Panicln("Unable to analyze swaggerSpec", err)
	}
	return api
}

func CreateServer(api *operations.ContactsAPI) (srv *restapi.Server) {
	srv = restapi.NewServer(api)
	srv.ConfigureAPI()

	handler := api.Serve(func(h http.Handler) http.Handler {
		// attach X-Apigateway-Api-Userinfo parser
		return firebase.HandleAuth(h)
	})

	corsHandler := cors.New(cors.Options{
		Debug:          false,
		AllowedHeaders: []string{"*"},
		AllowedOrigins: []string{conf.AllowedOrigin()},
		AllowedMethods: []string{"PUT", "POST", "DELETE", "PATCH", "GET"},
		MaxAge:         10000,
	})

	srv.SetHandler(corsHandler.Handler(handler))

	srv.EnabledListeners = []string{"http"}

	srv.Port = conf.PORT()
	return
}
