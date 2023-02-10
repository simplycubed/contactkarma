package repository

import (
	"context"
	"time"

	"github.com/simplycubed/contactkarma/contacts/domain"
)

type IContactSource interface {
	GetByEmail(ctx context.Context, id domain.UserID, email string, source domain.Source) (cs *domain.ContactSource, err error)
	Create(ctx context.Context, id domain.UserID, source domain.ContactSource) (created *domain.ContactSource, err error)
	GetAll(ctx context.Context, id domain.UserID) (sources []domain.ContactSource, err error)
	Get(ctx context.Context, id domain.UserID, limit int, lastDocumentId *domain.ContactSourceID) (sources []domain.ContactSource, err error)
	GetById(ctx context.Context, id domain.UserID, sourceId domain.ContactSourceID) (source domain.ContactSource, err error)
	Delete(ctx context.Context, id domain.UserID, sourceId domain.ContactSourceID) error
	UpdateByMap(ctx context.Context, uID domain.UserID, cID domain.ContactSourceID, updates map[string]interface{}) (err error)
	GetByNextUpdateAt(ctx context.Context, before time.Time) (sources []domain.ContactSource, err error)
}

type IGoogleContacts interface {
	Create(ctx context.Context, id domain.UserID, sourceId domain.ContactSourceID, personId string, contact domain.GoogleContact) (created *domain.GoogleContact, err error)
	Get(ctx context.Context, id domain.UserID, sourceId domain.ContactSourceID, personId string) (contact domain.GoogleContact, err error)
	Update(ctx context.Context, id domain.UserID, sourceId domain.ContactSourceID, personId string, contact domain.GoogleContact) (err error)
	List(ctx context.Context, id domain.UserID, sourceId domain.ContactSourceID, limit int, lastDocumentId *string) (sources []domain.GoogleContact, err error)
}
