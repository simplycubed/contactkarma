package routes

import (
	"log"

	"github.com/simplycubed/contactkarma/contacts/adapters/firebase"
	"github.com/simplycubed/contactkarma/contacts/application"
	"github.com/simplycubed/contactkarma/contacts/domain"
	"github.com/simplycubed/contactkarma/contacts/gen/models"
	"github.com/simplycubed/contactkarma/contacts/gen/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func ContactSource(api *operations.ContactsAPI, service *application.ContactSourceService) {

	api.InitGoogleContactSourceHandler = operations.InitGoogleContactSourceHandlerFunc(
		func(p operations.InitGoogleContactSourceParams) middleware.Responder {
			url, err := service.GetGoogleRedirectUrl(p.HTTPRequest.Context())
			if err != nil {
				return operations.NewInitGoogleContactSourceInternalServerError().WithPayload(&models.ErrorResponse{
					Description: "failed to get redirect url",
					Error:       err.Error(),
				})
			}
			// Send back redirect url which will take user to Google's consent page
			return operations.NewInitGoogleContactSourceOK().WithPayload(&models.InitResponse{
				URL: url,
			})
		},
	)

	api.LinkGoogleContactSourceHandler = operations.LinkGoogleContactSourceHandlerFunc(
		func(p operations.LinkGoogleContactSourceParams) middleware.Responder {
			userInfo := p.HTTPRequest.Context().Value(gatewayUserContext).(firebase.UserInfo)
			userID := userInfo.UserID
			role := userInfo.StripeRole
			err := service.LinkGoogleContactSource(p.HTTPRequest.Context(), domain.UserID(userID), role, p.Body.AuthCode)
			if err != nil {
				log.Println(err)
				if err == domain.ErrContactSourcesLimitReached {
					return operations.NewLinkGoogleContactSourceForbidden().WithPayload(&models.ErrorResponse{
						Description: "quota limit exceeded",
						Error:       err.Error(),
					})
				}
				return operations.NewLinkGoogleContactSourceInternalServerError().WithPayload(&models.ErrorResponse{
					Description: "failed to link",
					Error:       err.Error(),
				})
			}
			return operations.NewLinkGoogleContactSourceOK()
		},
	)

	api.GetContactSourcesHandler = operations.GetContactSourcesHandlerFunc(
		func(p operations.GetContactSourcesParams) middleware.Responder {

			limit, validationErr := ValidateLimit(p.Limit)
			if validationErr != nil {
				return operations.NewGetContactTagsBadRequest().WithPayload(validationErr)
			}

			userID := p.HTTPRequest.Context().Value(gatewayUserContext).(firebase.UserInfo).UserID
			sources, err := service.GetContactSources(p.HTTPRequest.Context(), domain.UserID(userID), limit, (*domain.ContactSourceID)(p.LastDocumentID))
			if err != nil {
				log.Println(err)
				return operations.NewGetContactSourcesInternalServerError().WithPayload(&models.ErrorResponse{
					Description: "failed list contact sources",
					Error:       "failed list contact sources",
				})
			}
			return operations.NewGetContactSourcesOK().WithPayload(sources)
		},
	)

	api.DeleteContactSourceHandler = operations.DeleteContactSourceHandlerFunc(
		func(p operations.DeleteContactSourceParams) middleware.Responder {
			userID := p.HTTPRequest.Context().Value(gatewayUserContext).(firebase.UserInfo).UserID
			err := service.DeleteContactSource(p.HTTPRequest.Context(), domain.UserID(userID), domain.ContactSourceID(p.SourceID), p.Body.RemoveFromUnified)
			if err != nil {
				log.Println(err)
				return operations.NewGetContactSourcesInternalServerError().WithPayload(&models.ErrorResponse{
					Description: "failed to delete contact source",
					Error:       "failed to delete contact source",
				})
			}
			return operations.NewDeleteContactSourceOK().WithPayload(&models.Message{
				Message: "contact source deleted successfully",
			})
		},
	)
}
