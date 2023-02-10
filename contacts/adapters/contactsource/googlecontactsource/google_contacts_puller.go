package googlecontactsource

import (
	"context"
	"log"

	"github.com/simplycubed/contactkarma/contacts/application"
	"github.com/simplycubed/contactkarma/contacts/domain"
	"github.com/simplycubed/contactkarma/contacts/domain/repository"
	"google.golang.org/api/people/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GoogleContactSourcePuller struct {
	userId             domain.UserID
	contactSourceId    domain.ContactSourceID
	googleContactsRepo repository.IGoogleContacts
	peopleService      application.PeopleService

	isPullComplete bool
	pageToken      *string
}

func NewGoogleContactSourcePuller(userId domain.UserID,
	contactSourceId domain.ContactSourceID,
	googleContactsRepo repository.IGoogleContacts,
	peopleService application.PeopleService) *GoogleContactSourcePuller {
	return &GoogleContactSourcePuller{contactSourceId: contactSourceId, userId: userId, googleContactsRepo: googleContactsRepo, peopleService: peopleService}
}

func (s *GoogleContactSourcePuller) Pull(ctx context.Context) (newContacts []domain.Contact, updatedContacts []domain.Contact, deletedContacts []domain.Contact, err error) {
	if s.isPullComplete {
		err = application.ErrPullCompleted
		return
	}

	var list *people.ListConnectionsResponse
	list, err = s.peopleService.List(s.pageToken)
	if err != nil {
		log.Println("failed to list connections", err)
		return
	}
	for _, connection := range list.Connections {
		contact := domain.GoogleContact{}
		contact.FromGooglePerson(connection)

		existing, getError := s.googleContactsRepo.Get(ctx, s.userId, s.contactSourceId, contact.ID.String())
		if status.Code(getError) == codes.NotFound {

			// new contact
			// create/update to contact source's contact list
			_, err = s.googleContactsRepo.Create(ctx, s.userId, s.contactSourceId, contact.ID.String(), contact)
			if err != nil {
				return
			}
			newContacts = append(newContacts, contact.MapToDomain())
		} else if getError != nil {
			// something else went wrong
			err = getError
			return
		}

		// check if any fields are updated
		if existing.IsContactDataUpdated(contact) {
			// write the update to source collection, so changes won't be added again on next sync
			err = s.googleContactsRepo.Update(ctx, s.userId, s.contactSourceId, contact.ID.String(), contact)
			if err != nil {
				return
			}
			updatedContacts = append(updatedContacts, contact.MapToDomain())
		}

	}
	log.Println("next page token: ", list.NextPageToken, "last batch count", len(list.Connections))
	if list.NextPageToken == "" || len(list.Connections) == 0 {
		// all contacts synced
		s.isPullComplete = true
		return
	}
	s.pageToken = &list.NextPageToken

	// TODO: add logic for deleted contacts

	return
}
