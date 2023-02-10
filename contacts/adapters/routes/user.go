package routes

import (
	"github.com/simplycubed/contactkarma/contacts/adapters/firebase"
	"github.com/simplycubed/contactkarma/contacts/application"
	"github.com/simplycubed/contactkarma/contacts/domain"
	"github.com/simplycubed/contactkarma/contacts/gen/models"
	"github.com/simplycubed/contactkarma/contacts/gen/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

const gatewayUserContext = "GATEWAY_USER"

func Users(api *operations.ContactsAPI, app application.UserService) {

	// create user
	api.CreateUserHandler = operations.CreateUserHandlerFunc(
		func(p operations.CreateUserParams) middleware.Responder {
			userID := p.HTTPRequest.Context().Value(gatewayUserContext).(firebase.UserInfo).UserID
			user, err := app.SaveUser(p.HTTPRequest.Context(), domain.UserID(userID), p.Body)
			if err != nil {
				return operations.NewCreateUserInternalServerError().WithPayload(&models.ErrorResponse{
					Description: "User creation failed",
					Error:       err.Error(),
				})
			}
			return operations.NewCreateUserOK().WithPayload(user)
		},
	)

	// update user
	api.UpdateUserHandler = operations.UpdateUserHandlerFunc(
		func(p operations.UpdateUserParams) middleware.Responder {
			userID := p.HTTPRequest.Context().Value(gatewayUserContext).(firebase.UserInfo).UserID
			user, err := app.UpdateUser(p.HTTPRequest.Context(), domain.UserID(userID), p.Body)
			if err != nil {
				return operations.NewUpdateUserInternalServerError().WithPayload(&models.ErrorResponse{
					Description: "User update failed",
					Error:       err.Error(),
				})
			}
			return operations.NewUpdateUserOK().WithPayload(user)
		},
	)

	// get user by id
	api.GetUserHandler = operations.GetUserHandlerFunc(
		func(p operations.GetUserParams) middleware.Responder {
			userID := p.HTTPRequest.Context().Value(gatewayUserContext).(firebase.UserInfo).UserID
			if string(userID) != string(p.UserID) {
				return operations.NewGetUserUnauthorized().WithPayload(&models.ErrorResponse{
					Description: "Unauthorized to access users/" + string(p.UserID),
					Error:       "Request is unauthorized",
				})
			}
			user, err := app.GetUser(p.HTTPRequest.Context(), domain.UserID(p.UserID))
			if err != nil {
				return operations.NewGetUserInternalServerError().WithPayload(&models.ErrorResponse{
					Description: "User fetching failed",
					Error:       err.Error(),
				})
			}
			return operations.NewGetUserOK().WithPayload(user)
		},
	)

	// delete user by id
	api.DeleteUserHandler = operations.DeleteUserHandlerFunc(
		func(p operations.DeleteUserParams) middleware.Responder {
			userID := p.HTTPRequest.Context().Value(gatewayUserContext).(firebase.UserInfo).UserID
			if string(userID) != string(p.UserID) {
				return operations.NewDeleteUserUnauthorized().WithPayload(&models.ErrorResponse{
					Description: "Unauthorized to access users/" + string(p.UserID),
					Error:       "Request is unauthorized",
				})
			}
			err := app.DeleteUser(p.HTTPRequest.Context(), domain.UserID(p.UserID))
			if err != nil {
				return operations.NewDeleteUserInternalServerError().WithPayload(&models.ErrorResponse{
					Description: "User fetching failed",
					Error:       err.Error(),
				})
			}
			return operations.NewDeleteUserOK().WithPayload(&models.Message{
				Message: "User has been successfully deleted",
			})
		},
	)
}
