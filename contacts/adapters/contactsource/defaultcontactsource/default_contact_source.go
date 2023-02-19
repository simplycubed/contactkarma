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

func (source *DefaultContactSource) Update(ctx context.Context, userId domain.UserID, sourceId domain.ContactSourceID, updates []domain.ContactSourceUpdate) (err error) {
	updateMap := map[domain.ContactID]*domain.Contact{}
	for _, update := range updates {
		contact := domain.Contact{}
		contact.FromUnified(update.Unified)
		contact.ID = domain.ContactID(update.ContactId)
		updateMap[update.ContactId] = &contact
	}
	err = source.repo.BulkUpdateContacts(ctx, userId, updateMap)
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

func (source *DefaultContactSource) Delete(ctx context.Context, userId domain.UserID, sourceId domain.ContactSourceID, contactIds []domain.ContactID) (err error) {
	return source.repo.BulkDeleteContacts(ctx, userId, contactIds)
}
