package jobs

import (
	"log"

	"github.com/simplycubed/contactkarma/contacts/conf"
	"github.com/simplycubed/contactkarma/contacts/gen-jobs/restapi"
	"github.com/simplycubed/contactkarma/contacts/gen-jobs/restapi/operations"
	"github.com/go-openapi/loads"
)

func CreateApi() *operations.ContactsJobsAPI {
	swaggerSpec, err := loads.Analyzed(
		restapi.FlatSwaggerJSON,
		"2.0",
	)
	api := operations.NewContactsJobsAPI(swaggerSpec)

	if err != nil {
		log.Panicln("Unable to analyze swaggerSpec", err)
	}
	return api
}

func CreateServer(api *operations.ContactsJobsAPI) (srv *restapi.Server) {
	srv = restapi.NewServer(api)
	srv.ConfigureAPI()
	srv.EnabledListeners = []string{"http"}
	srv.Port = conf.JobPort()
	return
}
