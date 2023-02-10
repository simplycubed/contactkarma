package jobs

import (
	"encoding/json"
	"log"

	"github.com/simplycubed/contactkarma/contacts/application"
	"github.com/simplycubed/contactkarma/contacts/domain"
	"github.com/simplycubed/contactkarma/contacts/gen-jobs/models"
	"github.com/simplycubed/contactkarma/contacts/gen-jobs/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func Routes(api *operations.ContactsJobsAPI, service *application.ContactSourceService) {

	// Pull contacts from a single contact source of a user
	api.PullContactSourceHandler = operations.PullContactSourceHandlerFunc(
		func(p operations.PullContactSourceParams) middleware.Responder {
			log.Println("message", p.Body.Message)
			log.Println("data", string(p.Body.Message.Data))

			request := models.PullContactsRequest{}
			err := json.Unmarshal(p.Body.Message.Data, &request)
			if err != nil {
				log.Println(err)
				return operations.NewPullContactSourceBadRequest().WithPayload(&models.JobError{
					Description: "failed to parse pull contacts message",
					Error:       "failed to parse pull contacts message",
				})
			}

			err = service.SyncContacts(p.HTTPRequest.Context(), domain.UserID(request.UserID), domain.ContactSourceID(request.ContactSourceID))
			if err != nil {
				log.Println("failed to SyncContacts", err)
				return operations.NewPullContactSourceBadRequest().WithPayload(&models.JobError{
					Description: "failed to sync contacts from source",
					Error:       "failed to sync contacts from source",
				})
			}
			return operations.NewPullContactSourceOK().WithPayload(&models.JobSuccess{
				Message: "contacts synced successfully",
			})
		},
	)

	api.PullContactsHandler = operations.PullContactsHandlerFunc(
		func(p operations.PullContactsParams) middleware.Responder {
			log.Println("message", p.Body.Message)
			log.Println("data", string(p.Body.Message.Data))
			err := service.PullContacts(p.HTTPRequest.Context())
			if err != nil {
				log.Println("pull contacts failed", err)
				return operations.NewPullContactSourceBadRequest().WithPayload(&models.JobError{
					Description: "failed run pull contacts",
					Error:       "failed run pull contacts",
				})
			}
			return operations.NewPullContactsOK().WithPayload(&models.JobSuccess{
				Message: "pull-contacts ran successfully",
			})
		},
	)

	// Job to remove contacts created from a source. Also optionally remove those contacts unified collection.
	api.ContactSourceCleanUpHandler = operations.ContactSourceCleanUpHandlerFunc(
		func(p operations.ContactSourceCleanUpParams) middleware.Responder {
			log.Println("message", p.Body.Message)
			log.Println("data", string(p.Body.Message.Data))

			request := models.ContactSourceDeleted{}
			err := json.Unmarshal(p.Body.Message.Data, &request)
			if err != nil {
				log.Println(err)
				return operations.NewContactSourceCleanUpBadRequest().WithPayload(&models.JobError{
					Description: "parsing failed",
					Error:       "parsing failed: " + err.Error(),
				})
			}

			err = service.OnDeleteContactSource(p.HTTPRequest.Context(), domain.UserID(request.UserID), domain.ContactSourceID(request.ContactSourceID), domain.Source(request.Source), request.RemoveContactsFromUnified)
			if err != nil {
				log.Println("delete tasks failed", err)
				return operations.NewPullContactSourceBadRequest().WithPayload(&models.JobError{
					Description: "failed",
					Error:       "failed",
				})
			}
			return operations.NewContactSourceCleanUpOK().WithPayload(&models.JobSuccess{
				Message: "deletions tasks ran successfully",
			})
		},
	)
}
