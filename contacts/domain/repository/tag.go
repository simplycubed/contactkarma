package repository

import (
	"context"

	"github.com/simplycubed/contactkarma/contacts/domain"
)

type ITag interface {
	//Get
	GetTagByID(ctx context.Context, uID domain.UserID, cID domain.UnifiedId, nID domain.TagID) (*domain.Tag, error)
	GetAllTags(context.Context, domain.UserID, domain.UnifiedId) ([]domain.Tag, error)
	GetTags(ctx context.Context, userId domain.UserID, contactId domain.UnifiedId, limit int, lastDocumentId *domain.TagID) ([]domain.Tag, error)

	//Patch
	UpdateTag(context.Context, domain.UserID, domain.UnifiedId, domain.TagID, domain.Tag) error

	//Save
	SaveTag(context.Context, domain.UserID, domain.UnifiedId, domain.Tag) (note *domain.Tag, err error)

	//Delete
	DeleteTag(context.Context, domain.UserID, domain.UnifiedId, domain.TagID) error
	DeleteAllTags(context.Context, domain.UserID, domain.UnifiedId) error
}
