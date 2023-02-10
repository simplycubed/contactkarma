package routes

import (
	"log"
	"time"

	"github.com/simplycubed/contactkarma/contacts/adapters/firebase"
	"github.com/simplycubed/contactkarma/contacts/application"
	"github.com/simplycubed/contactkarma/contacts/domain"
	"github.com/simplycubed/contactkarma/contacts/gen/models"
	"github.com/simplycubed/contactkarma/contacts/gen/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func Unified(api *operations.ContactsAPI, service application.IUnifiedContactService, linkSuggestionService application.ILinkSuggestionService, searchService application.ContactSearchService) {
	// get all contacts
	api.GetUnifiedContactsHandler = operations.GetUnifiedContactsHandlerFunc(
		func(p operations.GetUnifiedContactsParams) middleware.Responder {

			userID := p.HTTPRequest.Context().Value(gatewayUserContext).(firebase.UserInfo).UserID
			limit, validationErr := ValidateLimit(p.Limit)
			if validationErr != nil {
				return operations.NewGetUnifiedContactsBadRequest().WithPayload(validationErr)
			}

			contacts, err := service.GetContacts(p.HTTPRequest.Context(), domain.UserID(userID), limit, (*domain.UnifiedId)(p.LastDocumentID))
			if err != nil {
				return operations.NewGetUnifiedContactsInternalServerError().WithPayload(&models.ErrorResponse{
					Description: "Get unified contacts failed",
					Error:       err.Error(),
				})
			}
			return operations.NewGetUnifiedContactsOK().WithPayload(contacts)
		},
	)

	api.GetPendingFollowUpsHandler = operations.GetPendingFollowUpsHandlerFunc(
		func(p operations.GetPendingFollowUpsParams) middleware.Responder {

			limit, validationErr := ValidateLimit(p.Limit)
			if validationErr != nil {
				return operations.NewGetContactTagsBadRequest().WithPayload(validationErr)
			}
			userID := p.HTTPRequest.Context().Value(gatewayUserContext).(firebase.UserInfo).UserID
			pending, err := service.GetPendingContacts(p.HTTPRequest.Context(), domain.UserID(userID), limit, (*time.Time)(p.LastDocumentNextContact), (*domain.UnifiedId)(p.LastDocumentID))
			if err != nil {
				log.Println(err)
				return operations.NewGetPendingFollowUpsInternalServerError().WithPayload(&models.ErrorResponse{
					Description: "Failed to retrieve pending follow ups",
					Error:       "Failed to retrieve pending follow ups",
				})
			}
			return operations.NewGetPendingFollowUpsOK().WithPayload(pending)
		},
	)

	api.GetRecentContactsHandler = operations.GetRecentContactsHandlerFunc(
		func(p operations.GetRecentContactsParams) middleware.Responder {

			limit, validationErr := ValidateLimit(p.Limit)
			if validationErr != nil {
				return operations.NewGetContactTagsBadRequest().WithPayload(validationErr)
			}
			userID := p.HTTPRequest.Context().Value(gatewayUserContext).(firebase.UserInfo).UserID
			pending, err := service.GetRecentContacts(p.HTTPRequest.Context(), domain.UserID(userID), p.MaxDays, limit, (*time.Time)(p.LastDocumentLastContact), (*domain.UnifiedId)(p.LastDocumentID))
			if err != nil {
				log.Println(err)
				return operations.NewGetRecentContactsInternalServerError().WithPayload(&models.ErrorResponse{
					Description: "Failed to retrieve recent contacts",
					Error:       "Failed to retrieve recent contacts",
				})
			}
			return operations.NewGetRecentContactsOK().WithPayload(pending)
		},
	)

	api.GetLinkSuggestionsHandler = operations.GetLinkSuggestionsHandlerFunc(
		func(p operations.GetLinkSuggestionsParams) middleware.Responder {

			limit, validationErr := ValidateLimit(p.Limit)
			if validationErr != nil {
				return operations.NewGetContactTagsBadRequest().WithPayload(validationErr)
			}
			userID := p.HTTPRequest.Context().Value(gatewayUserContext).(firebase.UserInfo).UserID
			suggestions, err := linkSuggestionService.GetLinkSuggestions(p.HTTPRequest.Context(), domain.UserID(userID), limit, (*domain.LinkSuggestionID)(p.LastDocumentID))
			if err != nil {
				log.Println(err)
				return operations.NewGetRecentContactsInternalServerError().WithPayload(&models.ErrorResponse{
					Description: "Failed to retrieve suggestions",
					Error:       "Failed to retrieve suggestions",
				})
			}
			return operations.NewGetLinkSuggestionsOK().WithPayload(suggestions)
		},
	)

	api.ApplyLinkSuggestionHandler = operations.ApplyLinkSuggestionHandlerFunc(
		func(p operations.ApplyLinkSuggestionParams) middleware.Responder {
			userID := p.HTTPRequest.Context().Value(gatewayUserContext).(firebase.UserInfo).UserID
			unifiedIds := domain.UnifiedIds{}
			unifiedIds.FromArray(p.Body.UnifiedIds)
			err := linkSuggestionService.ApplyLinkSuggestion(p.HTTPRequest.Context(), domain.UserID(userID), domain.LinkSuggestionID(p.SuggestionID), unifiedIds)
			if err != nil {
				log.Println(err)
				return operations.NewApplyLinkSuggestionInternalServerError().WithPayload(&models.ErrorResponse{
					Description: "Failed to apply suggestion",
					Error:       "Failed to apply suggestion",
				})
			}
			return operations.NewApplyLinkSuggestionOK()
		},
	)

	api.SearchUserContactHandler = operations.SearchUserContactHandlerFunc(
		func(p operations.SearchUserContactParams) middleware.Responder {
			userID := p.HTTPRequest.Context().Value(gatewayUserContext).(firebase.UserInfo).UserID
			unified, err := searchService.SearchContacts(p.HTTPRequest.Context(), domain.UserID(userID), *p.Body)
			if err != nil {
				log.Println(err.Error())
				return operations.NewSearchUserContactInternalServerError().WithPayload(&models.ErrorResponse{
					Description: "Failed to search contacts",
					Error:       "Failed to search contacts",
				})
			}
			return operations.NewSearchUserContactOK().WithPayload(unified)
		},
	)

}
