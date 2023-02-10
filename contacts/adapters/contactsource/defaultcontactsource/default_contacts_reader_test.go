package defaultcontactsource

import (
	"context"
	"strconv"
	"testing"

	"github.com/simplycubed/contactkarma/contacts/adapters/firestore"
	"github.com/simplycubed/contactkarma/contacts/application"
	"github.com/simplycubed/contactkarma/contacts/conf"
	"github.com/simplycubed/contactkarma/contacts/domain"
	"github.com/simplycubed/contactkarma/contacts/test/testutils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

func TestDefaultContactsReader(t *testing.T) {
	testutils.LoadEnvFile()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	defer testutils.ClearDB()

	credetialsOption := option.WithCredentials(&google.Credentials{})
	dbpool, _ := firestore.NewFirestoreClient(conf.ProjectID(), credetialsOption)
	contactsFirestore := firestore.NewDefaultContactsFirestore(dbpool)

	userId := domain.UserID("userId")

	// write bunch of contacts
	for i := 0; i < 8; i++ {
		_, err := contactsFirestore.SaveContact(context.Background(), userId, domain.Contact{
			ID:    domain.ContactID("contact-" + strconv.Itoa(i)),
			Names: []*domain.UserNames{{GivenName: "contact-" + strconv.Itoa(i)}},
		})
		assert.Equal(t, nil, err)
	}

	reader := NewDefaultContactsReader(contactsFirestore, userId, 5)
	contacts, err := reader.Read(context.Background())
	assert.Equal(t, nil, err)
	assert.Equal(t, 5, len(contacts))

	contacts, err = reader.Read(context.Background())
	assert.Equal(t, nil, err)
	assert.Equal(t, 3, len(contacts))

	_, err = reader.Read(context.Background())
	assert.Equal(t, application.ErrSourceReadCompleted, err)
}
