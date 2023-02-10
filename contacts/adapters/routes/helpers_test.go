package routes

import (
	"context"
	"time"

	"github.com/simplycubed/contactkarma/contacts/domain"
	"github.com/simplycubed/contactkarma/contacts/domain/repository"
)

func createTestContact(contactRepo repository.IContact, userId domain.UserID, modify func(v *domain.Contact)) (created *domain.Contact, err error) {

	testContact := &domain.Contact{
		Addresses: []domain.Address{{
			City:        "London",
			Country:     "United Kingdom",
			CountryCode: "UK",
		}},
		Birthdays:      []domain.Birthday{{Date: "27-06-2000"}},
		EmailAddresses: []domain.EmailAddress{{Value: "john@gmail.com"}},
		Genders:        []domain.Gender{{Value: "male"}},
		Names:          []*domain.UserNames{{DisplayName: "John Doe"}},
	}

	// add any modification to model if any
	if modify != nil {
		modify(testContact)
	}

	created, err = contactRepo.SaveContact(context.Background(), userId, *testContact)
	return
}

func createTestUnified(unifiedRepo repository.IUnified, userId domain.UserID, modify func(v *domain.Unified)) (created *domain.Unified, err error) {

	testContact := domain.Unified{
		Addresses: []domain.Address{{
			City:        "London",
			Country:     "United Kingdom",
			CountryCode: "UK",
		}},
		Birthdays:      []domain.Birthday{{Date: "27-06-2000"}},
		EmailAddresses: []domain.EmailAddress{{Value: "john@gmail.com"}},
		Genders:        []domain.Gender{{Value: "male"}},
		Names:          []*domain.UserNames{{DisplayName: "John Doe"}},
	}

	// add any modification to model if any
	if modify != nil {
		modify(&testContact)
	}

	created, err = unifiedRepo.SaveContact(context.Background(), userId, testContact)
	return
}

func createTestContactNote(noteRepo repository.INote, userId domain.UserID, contactId domain.UnifiedId, modify func(v *domain.Note)) (created *domain.Note, err error) {

	testNote := &domain.Note{
		CreatedAt: time.Now(),
		IsUpdated: false,
		Note:      "test-note",
	}

	// add any modification to model if any
	if modify != nil {
		modify(testNote)
	}

	created, err = noteRepo.SaveNote(context.Background(), userId, contactId, *testNote)
	return
}

func createTestContactTag(tagRepo repository.ITag, userId domain.UserID, contactId domain.UnifiedId, modify func(v *domain.Tag)) (created *domain.Tag, err error) {

	testTag := &domain.Tag{
		CreatedAt: time.Now(),
		TagName:   "test-tag",
	}

	// add any modification to model if any
	if modify != nil {
		modify(testTag)
	}

	created, err = tagRepo.SaveTag(context.Background(), userId, contactId, *testTag)
	return
}
