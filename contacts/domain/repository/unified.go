//go:generate mockgen --build_flags=--mod=mod -package=mocks -destination=../../mocks/contact.go . IContact
package repository

import (
	"context"
	"errors"
	"time"

	"github.com/simplycubed/contactkarma/contacts/domain"
)

var ErrContactNotFound = errors.New("contact not found")

type IUnified interface {
	//Get
	GetAllContacts(context.Context, domain.UserID) ([]domain.Unified, error)

	GetContacts(ctx context.Context, uID domain.UserID, limit int, lastDocumentId *domain.UnifiedId) ([]domain.Unified, error)
	GetContactByID(context.Context, domain.UserID, domain.UnifiedId) (*domain.Unified, error)
	GetByIDs(ctx context.Context, uID domain.UserID, unifiedIds []domain.UnifiedId) (users []domain.Unified, err error)

	GetContactsByNextContact(ctx context.Context, user domain.UserID, before time.Time, limit int, lastDocumentInstant *time.Time, lastDocumentId *domain.UnifiedId) ([]domain.Unified, error)
	GetContactsByLastContact(ctx context.Context, user domain.UserID, after time.Time, limit int, lastDocumentInstant *time.Time, lastDocumentId *domain.UnifiedId) ([]domain.Unified, error)
	GetContactByOrigin(ctx context.Context, uID domain.UserID, origin domain.ContactOrigin) (unified domain.Unified, err error)
	GetContactBySearchTerms(ctx context.Context, uID domain.UserID, terms []string) (contacts []domain.Unified, err error)
	//Patch
	UpdateContact(context.Context, domain.UserID, domain.UnifiedId, *domain.Unified) error

	//Save
	SaveContact(context.Context, domain.UserID, domain.Unified) (*domain.Unified, error)

	//Delete
	DeleteContact(ctx context.Context, id domain.UserID, contactID domain.UnifiedId) error

	DeleteAllContacts(ctx context.Context, uID domain.UserID) error

	BulkDeleteContacts(ctx context.Context, uID domain.UserID, contactIDs []domain.UnifiedId) (err error)
	BulkUpdateContacts(ctx context.Context, uID domain.UserID, updates map[domain.UnifiedId]*domain.Unified) (err error)
}
