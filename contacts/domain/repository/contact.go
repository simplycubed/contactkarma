//go:generate mockgen --build_flags=--mod=mod -package=mocks -destination=../../mocks/contact.go . IContact
package repository

import (
	"context"
	"time"

	"github.com/simplycubed/contactkarma/contacts/domain"
)

type IContact interface {
	//Get
	GetAllContacts(context.Context, domain.UserID) ([]domain.Contact, error)

	GetContacts(ctx context.Context, uID domain.UserID, limit int, lastDocumentId *domain.ContactID) ([]domain.Contact, error)
	GetContactByID(context.Context, domain.UserID, domain.ContactID) (*domain.Contact, error)
	GetContactsByNextContact(ctx context.Context, user domain.UserID, before time.Time, limit int, lastDocumentInstant *time.Time, lastDocumentId *domain.ContactID) ([]domain.Contact, error)
	GetContactsByLastContact(ctx context.Context, user domain.UserID, after time.Time, limit int, lastDocumentInstant *time.Time, lastDocumentId *domain.ContactID) ([]domain.Contact, error)
	//Patch
	UpdateContact(context.Context, domain.UserID, domain.ContactID, *domain.Contact) error

	//Save
	SaveContact(context.Context, domain.UserID, domain.Contact) (*domain.Contact, error)

	//Delete
	DeleteContact(ctx context.Context, id domain.UserID, contactID domain.ContactID) error

	DeleteAllContacts(ctx context.Context, uID domain.UserID) error
}
