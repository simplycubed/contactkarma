//go:generate mockgen --build_flags=--mod=mod -package=mocks -destination=../mocks/application.go . Application
package application

import (
	"context"

	"github.com/simplycubed/contactkarma/contacts/domain"
	"github.com/simplycubed/contactkarma/contacts/gen/models"
)

type UserService interface {
	SaveUser(context.Context, domain.UserID, *models.User) (*models.User, error)
	UpdateUser(context.Context, domain.UserID, *models.User) (*models.User, error)
	GetUser(context.Context, domain.UserID) (*models.User, error)
	DeleteUser(context.Context, domain.UserID) error
}

type ContactService interface {
	SaveContact(context.Context, domain.UserID, *models.CreateContactDto) (updated *models.Unified, err error)
	UpdateContact(context.Context, domain.UserID, domain.UnifiedId, *models.UpdateUnifiedDto) (*models.Unified, error)
	UpdateCategory(ctx context.Context, userId domain.UserID, role *domain.Role, unifiedId domain.UnifiedId, category domain.ContactCatgeory) (*models.Unified, error)
	GetContacts(ctx context.Context, id domain.UserID, limit int, lastDocumentId *domain.ContactID) ([]*models.Contact, error)
	GetContactByID(context.Context, domain.UserID, domain.UnifiedId) (*models.Unified, error)
	DeleteContact(context.Context, domain.UserID, domain.UnifiedId) error
}

type NoteService interface {
	SaveNote(context.Context, domain.UserID, domain.UnifiedId, *models.Note) (updated *models.Note, err error)
	UpdateNote(context.Context, domain.UserID, domain.UnifiedId, domain.NoteID, *models.Note) (*models.Note, error)
	GetAllNotes(context.Context, domain.UserID, domain.UnifiedId) ([]*models.Note, error)
	GetNotes(ctx context.Context, userID domain.UserID, contactId domain.UnifiedId, limit int, lastDocumentId *domain.NoteID) ([]*models.Note, error)
	DeleteNote(context.Context, domain.UserID, domain.UnifiedId, domain.NoteID) error
}

type TagService interface {
	SaveTag(context.Context, domain.UserID, domain.UnifiedId, *models.Tag) (updated *models.Tag, err error)
	UpdateTag(context.Context, domain.UserID, domain.UnifiedId, domain.TagID, *models.Tag) (*models.Tag, error)
	GetAllTags(context.Context, domain.UserID, domain.UnifiedId) ([]*models.Tag, error)
	GetTags(ctx context.Context, userId domain.UserID, contactId domain.UnifiedId, limit int, lastDocumentId *domain.TagID) ([]*models.Tag, error)
	DeleteTag(context.Context, domain.UserID, domain.UnifiedId, domain.TagID) error
}
