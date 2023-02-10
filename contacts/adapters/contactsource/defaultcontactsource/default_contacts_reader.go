package defaultcontactsource

import (
	"context"

	"github.com/simplycubed/contactkarma/contacts/application"
	"github.com/simplycubed/contactkarma/contacts/domain"
	"github.com/simplycubed/contactkarma/contacts/domain/repository"
)

type DefaultContactsReader struct {
	contactsRepo   repository.IContact
	userId         domain.UserID
	batchSize      int
	lastDocumentId *domain.ContactID
	isPullComplete bool
}

func NewDefaultContactsReader(contactsRepo repository.IContact, userId domain.UserID, batchSize int) *DefaultContactsReader {
	return &DefaultContactsReader{
		userId:       userId,
		contactsRepo: contactsRepo,
		batchSize:    batchSize,
	}
}

func (reader *DefaultContactsReader) Read(ctx context.Context) (contacts []domain.Contact, err error) {
	if reader.isPullComplete {
		err = application.ErrSourceReadCompleted
		return
	}
	contacts, err = reader.contactsRepo.GetContacts(ctx, reader.userId, reader.batchSize, reader.lastDocumentId)
	if err != nil {
		return
	}
	if len(contacts) > 0 {
		lastId := contacts[len(contacts)-1].ID
		reader.lastDocumentId = &lastId
	}
	if len(contacts) == 0 {
		reader.isPullComplete = true
		err = application.ErrSourceReadCompleted
	}
	return
}
