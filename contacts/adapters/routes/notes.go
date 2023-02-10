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

func Notes(api *operations.ContactsAPI, noteService application.NoteService) {
	api.PostContactNoteHandler = operations.PostContactNoteHandlerFunc(
		func(p operations.PostContactNoteParams) middleware.Responder {
			userID := p.HTTPRequest.Context().Value(gatewayUserContext).(firebase.UserInfo).UserID
			_, err := noteService.SaveNote(p.HTTPRequest.Context(), domain.UserID(userID), domain.UnifiedId(p.UnifiedID), p.Body)
			if err != nil {
				return operations.NewPostContactNoteInternalServerError().WithPayload(&models.ErrorResponse{
					Description: "Note creation failed",
					Error:       err.Error(),
				})
			}
			return operations.NewPostContactNoteOK().WithPayload(&models.Message{
				Message: "Note has been created",
			})
		},
	)

	api.PatchContactNoteHandler = operations.PatchContactNoteHandlerFunc(
		func(p operations.PatchContactNoteParams) middleware.Responder {
			userID := p.HTTPRequest.Context().Value(gatewayUserContext).(firebase.UserInfo).UserID
			note, err := noteService.UpdateNote(p.HTTPRequest.Context(), domain.UserID(userID), domain.UnifiedId(p.UnifiedID), domain.NoteID(p.NoteID), p.Body)
			if err != nil {
				if status.Code(err) == codes.NotFound {
					return operations.NewPatchContactNoteNotFound().WithPayload(&models.ErrorResponse{
						Description: "Note not found update failed",
						Error:       err.Error(),
					})
				}
				return operations.NewPatchContactNoteInternalServerError().WithPayload(&models.ErrorResponse{
					Description: "Note update failed",
					Error:       err.Error(),
				})
			}
			return operations.NewPatchContactNoteOK().WithPayload(note)
		},
	)

	api.DeleteContactNoteHandler = operations.DeleteContactNoteHandlerFunc(
		func(p operations.DeleteContactNoteParams) middleware.Responder {
			userID := p.HTTPRequest.Context().Value(gatewayUserContext).(firebase.UserInfo).UserID
			err := noteService.DeleteNote(p.HTTPRequest.Context(), domain.UserID(userID), domain.UnifiedId(p.UnifiedID), domain.NoteID(p.NoteID))
			if err != nil {
				return operations.NewDeleteContactNoteInternalServerError().WithPayload(&models.ErrorResponse{
					Description: "Note delete failed",
					Error:       err.Error(),
				})
			}
			return operations.NewDeleteContactNoteOK().WithPayload(&models.Message{
				Message: "Note deleted successfully",
			})
		},
	)

	api.GetContactNotesHandler = operations.GetContactNotesHandlerFunc(
		func(p operations.GetContactNotesParams) middleware.Responder {

			limit, validationErr := ValidateLimit(p.Limit)
			if validationErr != nil {
				return operations.NewGetContactNotesBadRequest().WithPayload(validationErr)
			}

			userID := p.HTTPRequest.Context().Value(gatewayUserContext).(firebase.UserInfo).UserID
			notes, err := noteService.GetNotes(p.HTTPRequest.Context(), domain.UserID(userID), domain.UnifiedId(p.UnifiedID), limit, (*domain.NoteID)(p.LastDocumentID))
			if err != nil {
				return operations.NewGetContactNotesInternalServerError().WithPayload(&models.ErrorResponse{
					Description: "Get notes failed",
					Error:       err.Error(),
				})
			}
			return operations.NewGetContactNotesOK().WithPayload(notes)
		},
	)

}
