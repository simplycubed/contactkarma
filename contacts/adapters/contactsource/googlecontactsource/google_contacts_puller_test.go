package googlecontactsource

import (
	"context"
	"testing"

	"github.com/simplycubed/contactkarma/contacts/adapters/firestore"
	"github.com/simplycubed/contactkarma/contacts/application"
	"github.com/simplycubed/contactkarma/contacts/conf"
	"github.com/simplycubed/contactkarma/contacts/gen/mocks/mock_application"
	"github.com/simplycubed/contactkarma/contacts/test/testutils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/people/v1"
)

func TestGoogleContactsPuller(t *testing.T) {
	testutils.LoadEnvFile()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	defer testutils.ClearDB()

	credetialsOption := option.WithCredentials(&google.Credentials{})
	dbpool, _ := firestore.NewFirestoreClient(conf.ProjectID(), credetialsOption)
	googleContactsFirestore := firestore.NewGoogleContactsFirestore(dbpool)

	mockPeopleService := mock_application.NewMockPeopleService(ctrl)
	// expectations
	listResponse := people.ListConnectionsResponse{
		Connections: []*people.Person{
			{
				ResourceName: "people/12345",
				EmailAddresses: []*people.EmailAddress{
					{
						DisplayName: "John Doe",
						Value:       "johndoe@gmail.com",
					},
				},
			},
		},
		NextPageToken:   "",
		NextSyncToken:   "",
		TotalItems:      0,
		TotalPeople:     0,
		ServerResponse:  googleapi.ServerResponse{},
		ForceSendFields: []string{},
		NullFields:      []string{},
	}
	mockPeopleService.EXPECT().List(nil).Return(&listResponse, nil)

	puller := NewGoogleContactSourcePuller("testUser", "testSource", googleContactsFirestore, mockPeopleService)
	contacts, _, _, err := puller.Pull(context.Background())
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(contacts))

	_, _, _, err = puller.Pull(context.Background())
	assert.Equal(t, application.ErrPullCompleted, err)
}
