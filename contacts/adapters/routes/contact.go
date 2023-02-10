package routes

import (
	"github.com/simplycubed/contactkarma/contacts/adapters/firebase"
	"github.com/simplycubed/contactkarma/contacts/application"
	"github.com/simplycubed/contactkarma/contacts/domain"
	"github.com/simplycubed/contactkarma/contacts/gen/models"
	"github.com/simplycubed/contactkarma/contacts/gen/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Contacts(api *operations.ContactsAPI, contactService application.ContactService, csvImporter application.ICsvImporter) {

	// create contact
	api.CreateUserContactHandler = operations.CreateUserContactHandlerFunc(
		func(p operations.CreateUserContactParams) middleware.Responder {
			userID := p.HTTPRequest.Context().Value(gatewayUserContext).(firebase.UserInfo).UserID
			contact, err := contactService.SaveContact(p.HTTPRequest.Context(), domain.UserID(userID), p.Body)
			if err != nil {
				return operations.NewCreateUserContactInternalServerError().WithPayload(&models.ErrorResponse{
					Description: "Contact creation failed",
					Error:       err.Error(),
				})
			}
			return operations.NewCreateUserContactOK().WithPayload(contact)
		},
	)

	// update contact
	api.UpdateUserContactHandler = operations.UpdateUserContactHandlerFunc(
		func(p operations.UpdateUserContactParams) middleware.Responder {
			userID := p.HTTPRequest.Context().Value(gatewayUserContext).(firebase.UserInfo).UserID
			contact, err := contactService.UpdateContact(p.HTTPRequest.Context(), domain.UserID(userID), domain.UnifiedId(p.UnifiedID), p.Body)
			if err != nil {
				if status.Code(err) == codes.NotFound {
					return operations.NewUpdateUserContactNotFound().WithPayload(&models.ErrorResponse{
						Description: "Contact not found update failed",
						Error:       err.Error(),
					})
				}
				return operations.NewUpdateUserContactInternalServerError().WithPayload(&models.ErrorResponse{
					Description: "Update contact failed",
					Error:       err.Error(),
				})
			}
			return operations.NewUpdateUserContactOK().WithPayload(contact)
		},
	)

	// update contact category
	api.UpdateContactCategoryHandler = operations.UpdateContactCategoryHandlerFunc(
		func(p operations.UpdateContactCategoryParams) middleware.Responder {
			userInfo := p.HTTPRequest.Context().Value(gatewayUserContext).(firebase.UserInfo)
			userID := userInfo.UserID
			role := userInfo.StripeRole
			category := domain.ContactCatgeory(p.Body.Category)
			if !category.Valid() {
				return operations.NewUpdateUserContactBadRequest().WithPayload(&models.ErrorResponse{
					Description: "Invalid Category",
					Error:       "Invalid category value: " + string(category),
				})
			}
			contact, err := contactService.UpdateCategory(p.HTTPRequest.Context(), domain.UserID(userID), role, domain.UnifiedId(p.UnifiedID), category)
			if err != nil {
				if status.Code(err) == codes.NotFound {
					return operations.NewUpdateContactCategoryNotFound().WithPayload(&models.ErrorResponse{
						Description: "Contact not found, update failed",
						Error:       err.Error(),
					})
				}
				return operations.NewUpdateUserContactInternalServerError().WithPayload(&models.ErrorResponse{
					Description: "Update contact failed",
					Error:       err.Error(),
				})
			}
			return operations.NewUpdateContactCategoryOK().WithPayload(contact)
		},
	)

	// get contact by id
	api.GetUserContactByIDHandler = operations.GetUserContactByIDHandlerFunc(
		func(p operations.GetUserContactByIDParams) middleware.Responder {
			userID := p.HTTPRequest.Context().Value(gatewayUserContext).(firebase.UserInfo).UserID
			contact, err := contactService.GetContactByID(p.HTTPRequest.Context(), domain.UserID(userID), domain.UnifiedId(p.UnifiedID))
			if err != nil {
				if status.Code(err) == codes.NotFound {
					return operations.NewGetUserContactByIDNotFound().WithPayload(&models.ErrorResponse{
						Description: "Contact not found",
						Error:       err.Error(),
					})
				}
				return operations.NewGetUserContactByIDInternalServerError().WithPayload(&models.ErrorResponse{
					Description: "Get contact failed",
					Error:       err.Error(),
				})
			}
			return operations.NewGetUserContactByIDOK().WithPayload(contact)
		},
	)

	// delete contact by id
	api.DeleteUserContactHandler = operations.DeleteUserContactHandlerFunc(
		func(p operations.DeleteUserContactParams) middleware.Responder {
			userID := p.HTTPRequest.Context().Value(gatewayUserContext).(firebase.UserInfo).UserID
			err := contactService.DeleteContact(p.HTTPRequest.Context(), domain.UserID(userID), domain.UnifiedId(p.UnifiedID))
			if err != nil {
				return operations.NewDeleteUserContactInternalServerError().WithPayload(&models.ErrorResponse{
					Description: "Delete contact failed",
					Error:       err.Error(),
				})
			}
			return operations.NewDeleteUserContactOK().WithPayload(&models.Message{
				Message: "Contact has been successfully deleted",
			})
		},
	)

	api.UploadContactsCsvHandler = operations.UploadContactsCsvHandlerFunc(
		func(p operations.UploadContactsCsvParams) middleware.Responder {
			userID := p.HTTPRequest.Context().Value(gatewayUserContext).(firebase.UserInfo).UserID
			err := csvImporter.Import(p.HTTPRequest.Context(), p.File, domain.UserID(userID))
			if err != nil {
				return operations.NewUploadContactsCsvBadRequest().WithPayload(&models.ErrorResponse{
					Description: "Failed to upload contacts",
					Error:       err.Error(),
				})
			}
			return operations.NewUploadContactsCsvOK().WithPayload(&models.Message{
				Message: "Contacts have been uploaded successfully",
			})
		},
	)
}
