package defaultcontactsource

import (
	"context"

	"github.com/simplycubed/contactkarma/contacts/adapters/firestore"
	"github.com/simplycubed/contactkarma/contacts/application"
	"github.com/simplycubed/contactkarma/contacts/domain"
)

type DefaultContactSource struct {
	repo *firestore.DefaultContactsFirestore
}

func NewDefaultContactSource(repo *firestore.DefaultContactsFirestore) *DefaultContactSource {
	return &DefaultContactSource{repo: repo}
}

func (source *DefaultContactSource) Update(ctx context.Context, userId domain.UserID, sourceId domain.ContactSourceID, contactId domain.ContactID, unified domain.Unified) (err error) {
	contact := domain.Contact{}
	contact.FromUnified(unified)
	contact.ID = domain.ContactID(contactId)
	err = source.repo.UpdateContact(ctx, userId, contactId, &contact)
	return
}

func (source *DefaultContactSource) Puller(ctx context.Context, userId domain.UserID, contactSource domain.ContactSource) (puller application.IContactSourcePuller) {
	return NewDefaultContactSourcePuller()
}

func (source *DefaultContactSource) Reader(ctx context.Context, userId domain.UserID, contactSource domain.ContactSourceID) (puller application.IContactSourceReader) {
	batchSize := 100
	return NewDefaultContactsReader(source.repo, userId, batchSize)
}

func (source *DefaultContactSource) Remove(ctx context.Context, userId domain.UserID, sourceId domain.ContactSourceID, contactIds []domain.ContactID) (err error) {
	return source.repo.BulkDeleteContacts(ctx, userId, contactIds)
}
