package googlecontactsource

import (
	"context"

	"github.com/simplycubed/contactkarma/contacts/application"
	"github.com/simplycubed/contactkarma/contacts/domain"
	"github.com/simplycubed/contactkarma/contacts/domain/repository"
)

type GoogleContactsReader struct {
	googleContactsRepo repository.IGoogleContacts
	sourceId           domain.ContactSourceID
	userId             domain.UserID
	batchSize          int
	lastPersonId       *string
	isPullComplete     bool
}

func NewGoogleContactsReader(googleContactsRepo repository.IGoogleContacts, userId domain.UserID, sourceId domain.ContactSourceID, batchSize int) *GoogleContactsReader {
	return &GoogleContactsReader{
		userId:             userId,
		sourceId:           sourceId,
		googleContactsRepo: googleContactsRepo,
		batchSize:          batchSize,
	}
}

func (reader *GoogleContactsReader) Read(ctx context.Context) (contacts []domain.Contact, err error) {
	if reader.isPullComplete {
		err = application.ErrSourceReadCompleted
		return
	}
	googleContacts, err := reader.googleContactsRepo.List(ctx, reader.userId, reader.sourceId, reader.batchSize, reader.lastPersonId)
	if err != nil {
		return
	}
	for _, googleContact := range googleContacts {
		contacts = append(contacts, googleContact.MapToDomain())
	}
	if len(contacts) > 0 {
		lastId := string(contacts[len(contacts)-1].ID)
		reader.lastPersonId = &lastId
	}
	if len(contacts) == 0 {
		reader.isPullComplete = true
		err = application.ErrSourceReadCompleted
	}
	return
}
