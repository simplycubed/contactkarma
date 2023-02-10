package defaultcontactsource

import (
	"context"

	"github.com/simplycubed/contactkarma/contacts/application"
	"github.com/simplycubed/contactkarma/contacts/domain"
)

type DefaultContactSourcePuller struct {
}

func NewDefaultContactSourcePuller() *DefaultContactSourcePuller {
	return &DefaultContactSourcePuller{}
}

func (s *DefaultContactSourcePuller) Pull(ctx context.Context) (newContacts []domain.Contact, updatedContacts []domain.Contact, deletedContacts []domain.Contact, err error) {
	// nothing to pull for default source
	err = application.ErrPullCompleted
	return
}
