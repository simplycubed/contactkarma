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

func Tags(api *operations.ContactsAPI, tagService application.TagService) {
	api.PostContactTagHandler = operations.PostContactTagHandlerFunc(
		func(p operations.PostContactTagParams) middleware.Responder {
			userID := p.HTTPRequest.Context().Value(gatewayUserContext).(firebase.UserInfo).UserID
			_, err := tagService.SaveTag(p.HTTPRequest.Context(), domain.UserID(userID), domain.UnifiedId(p.UnifiedID), p.Body)
			if err != nil {
				return operations.NewPostContactTagInternalServerError().WithPayload(&models.ErrorResponse{
					Description: "Tag creation failed",
					Error:       err.Error(),
				})
			}
			return operations.NewPostContactTagOK().WithPayload(&models.Message{
				Message: "Tag has been created",
			})
		},
	)

	api.PatchContactTagHandler = operations.PatchContactTagHandlerFunc(
		func(p operations.PatchContactTagParams) middleware.Responder {
			userID := p.HTTPRequest.Context().Value(gatewayUserContext).(firebase.UserInfo).UserID
			tag, err := tagService.UpdateTag(p.HTTPRequest.Context(), domain.UserID(userID), domain.UnifiedId(p.UnifiedID), domain.TagID(p.TagID), p.Body)
			if err != nil {
				if status.Code(err) == codes.NotFound {
					return operations.NewPatchContactTagNotFound().WithPayload(&models.ErrorResponse{
						Description: "Tag not found update failed",
						Error:       err.Error(),
					})
				}
				return operations.NewPatchContactTagInternalServerError().WithPayload(&models.ErrorResponse{
					Description: "Tag update failed",
					Error:       err.Error(),
				})
			}
			return operations.NewPatchContactTagOK().WithPayload(tag)
		},
	)

	api.DeleteContactTagHandler = operations.DeleteContactTagHandlerFunc(
		func(p operations.DeleteContactTagParams) middleware.Responder {
			userID := p.HTTPRequest.Context().Value(gatewayUserContext).(firebase.UserInfo).UserID
			err := tagService.DeleteTag(p.HTTPRequest.Context(), domain.UserID(userID), domain.UnifiedId(p.UnifiedID), domain.TagID(p.TagID))
			if err != nil {
				return operations.NewDeleteContactTagInternalServerError().WithPayload(&models.ErrorResponse{
					Description: "Tag delete failed",
					Error:       err.Error(),
				})
			}
			return operations.NewDeleteContactNoteOK().WithPayload(&models.Message{
				Message: "Tag deleted successfully",
			})
		},
	)

	api.GetContactTagsHandler = operations.GetContactTagsHandlerFunc(
		func(p operations.GetContactTagsParams) middleware.Responder {

			limit, validationErr := ValidateLimit(p.Limit)
			if validationErr != nil {
				return operations.NewGetContactTagsBadRequest().WithPayload(validationErr)
			}

			userID := p.HTTPRequest.Context().Value(gatewayUserContext).(firebase.UserInfo).UserID
			tags, err := tagService.GetTags(p.HTTPRequest.Context(), domain.UserID(userID), domain.UnifiedId(p.UnifiedID), limit, (*domain.TagID)(p.LastDocumentID))
			if err != nil {
				return operations.NewGetContactTagsInternalServerError().WithPayload(&models.ErrorResponse{
					Description: "Get tags failed",
					Error:       err.Error(),
				})
			}
			return operations.NewGetContactTagsOK().WithPayload(tags)
		},
	)

}
