package jobs

import (
	"context"
	"encoding/json"
	"os"
	"sort"
	"testing"
	"time"

	googleFirestore "cloud.google.com/go/firestore"
	googlePubsub "cloud.google.com/go/pubsub"
	"github.com/go-openapi/strfmt"
	"github.com/golang/mock/gomock"
	"github.com/simplycubed/contactkarma/contacts/adapters"
	"github.com/simplycubed/contactkarma/contacts/adapters/contactsource/defaultcontactsource"
	"github.com/simplycubed/contactkarma/contacts/adapters/firestore"
	"github.com/simplycubed/contactkarma/contacts/adapters/pubsub"
	"github.com/simplycubed/contactkarma/contacts/application"
	"github.com/simplycubed/contactkarma/contacts/conf"
	"github.com/simplycubed/contactkarma/contacts/domain"
	"github.com/simplycubed/contactkarma/contacts/gen-jobs/client/operations"
	"github.com/simplycubed/contactkarma/contacts/gen-jobs/models"
	"github.com/simplycubed/contactkarma/contacts/gen/mocks/mock_application"
	"github.com/simplycubed/contactkarma/contacts/test/testutils"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type testContext struct {
	*testing.T
	*testutils.Env
	client                    operations.ClientService
	dbClient                  *googleFirestore.Client
	googleOAuthService        *mock_application.MockGoogleOAuthService
	peopleServiceFactory      *mock_application.MockPeopleServiceFactory
	googleContactsFirestore   *firestore.GoogleContactsFirestore
	contactSourceRepo         *firestore.ContactSourceFirestore
	pubsubClient              *googlePubsub.Client
	unifiedRepo               *firestore.UnifiedContactFirestore
	unifiedService            *application.UnifiedContactService
	linkSuggestionRepo        *firestore.LinkSuggestionFirestore
	mockContactSourceProvider *mock_application.MockIContactSourceProvider
	defaultContactFirestore   *firestore.DefaultContactsFirestore
}

func newTestContext(t *testing.T, ctrl *gomock.Controller) *testContext {

	googleOAuthService := mock_application.NewMockGoogleOAuthService(ctrl)
	credetialsOption := option.WithCredentials(&google.Credentials{})
	dbClient := adapters.NewFirestore(credetialsOption)

	userFirestore := firestore.NewUserFirestore(dbClient)

	contactSourceFirestore := firestore.NewContactSourceFirestore(dbClient)
	googleContactsFirestore := firestore.NewGoogleContactsFirestore(dbClient)
	peopleServiceFactory := mock_application.NewMockPeopleServiceFactory(ctrl)
	pubsubClient := pubsub.New(conf.ProjectID(), credetialsOption)
	pullContactPublisher := pubsub.NewPullContactSourcePublisher(pubsubClient.Topic(conf.PullContactsSourceTopic()))
	contactSourceDeletedPublisher := pubsub.NewContactSourceDeletedPublisher(pubsubClient.Topic(conf.ContactSourceDeletedTopic()))
	unifiedRepo := firestore.NewUnifiedContactFirestore(dbClient)
	linkSuggestionRepo := firestore.NewLinkSuggestionFirestore(dbClient)
	linkSuggestionService := application.NewLinkSuggestionService(unifiedRepo, linkSuggestionRepo)
	contactLogRepo := firestore.NewContactLogFirestore(dbClient)
	unifiedService := application.NewUnifiedContactService(unifiedRepo, linkSuggestionService, contactLogRepo)
	defaultContactFirestore := firestore.NewDefaultContactsFirestore(dbClient)
	mockContactSourceProvider := mock_application.NewMockIContactSourceProvider(ctrl)
	contactSourceService := application.NewContactSourceService(googleOAuthService, contactSourceFirestore, pullContactPublisher, unifiedService, userFirestore, mockContactSourceProvider, unifiedRepo, contactSourceDeletedPublisher)

	userService := application.NewUserService(userFirestore)

	api := CreateApi()
	Routes(api, contactSourceService)

	server := CreateServer(api)
	transport := testutils.NewTestClientTransport(server.GetHandler())
	client := operations.New(transport, strfmt.Default)
	env := testutils.NewEnv(server.GetHandler(), userService)
	return &testContext{
		T:                         t,
		Env:                       env,
		client:                    client,
		googleOAuthService:        googleOAuthService,
		peopleServiceFactory:      peopleServiceFactory,
		googleContactsFirestore:   googleContactsFirestore,
		contactSourceRepo:         contactSourceFirestore,
		pubsubClient:              pubsubClient,
		dbClient:                  dbClient,
		unifiedRepo:               unifiedRepo,
		unifiedService:            unifiedService,
		linkSuggestionRepo:        linkSuggestionRepo,
		mockContactSourceProvider: mockContactSourceProvider,
		defaultContactFirestore:   defaultContactFirestore,
	}
}

func TestMain(m *testing.M) {
	testutils.LoadEnvFile()
	// run tests
	exitVal := m.Run()
	os.Exit(exitVal)
}

func (ctx *testContext) setupTestUser() {
	err := ctx.AddTestUser()
	if err != nil {
		ctx.Fatal("Could not create user:", err)
	}
}

func (ctx *testContext) createTestContactSource(modify func(v *domain.ContactSource)) (created *domain.ContactSource, err error) {
	testSource := domain.ContactSource{
		UserID:       domain.UserID(ctx.User.UserID),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Source:       domain.Google,
		Email:        ctx.User.Email,
		GoogleUserId: "google-user-id",
		AccessToken:  "access-token",
		RefreshToken: "refresh-token",
		TokenExpiry:  time.Now().Add(5 * time.Second),
	}

	// add any modification to model if any
	if modify != nil {
		modify(&testSource)
	}

	created, err = ctx.contactSourceRepo.Create(context.Background(), domain.UserID(ctx.User.UserID), testSource)
	if err != nil {
		ctx.Fatal("Could not create source:", err)
	}
	return
}

func TestPullContactSourceHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testContext := newTestContext(t, ctrl)
	testContext.setupTestUser()
	t.Cleanup(func() { testContext.ClearDB() })

	source, err := testContext.createTestContactSource(func(v *domain.ContactSource) {
		v.Source = domain.Default
	})
	assert.Equal(t, err, nil)

	mockContactSource := mock_application.NewMockIContactSource(ctrl)
	mockPuller := mock_application.NewMockIContactSourcePuller(ctrl)
	mockContactSource.EXPECT().Puller(gomock.Any(), domain.UserID(testContext.User.UserID), gomock.Any()).Return(mockPuller)
	gomock.InOrder(
		mockPuller.EXPECT().Pull(gomock.Any()).Return([]domain.Contact{{
			EmailAddresses: []domain.EmailAddress{
				{
					DisplayName: "John Doe",
					Value:       "johndoe@gmail.com",
				},
			},
			ID: "test-contact",
		}}, nil, nil, nil),
		mockPuller.EXPECT().Pull(gomock.Any()).Return(nil, nil, nil, application.ErrPullCompleted),
	)

	testContext.mockContactSourceProvider.EXPECT().Get(domain.Default).Return(mockContactSource)

	pullContactsRequest := models.PullContactsRequest{
		UserID:          string(testContext.User.UserID),
		ContactSourceID: string(source.ID),
	}

	messageData, err := json.Marshal(pullContactsRequest)
	assert.Equal(t, err, nil)
	params := &operations.PullContactSourceParams{
		Body: &models.PubsubMessage{
			Message: &models.Message{
				Data:        messageData,
				MessageID:   new(string),
				PublishTime: strfmt.DateTime(time.Now()),
			},
		},
	}
	_, err = testContext.client.PullContactSource(params)
	assert.Equal(t, nil, err)
	ctx := context.Background()

	// contact should be synced to unified
	unifiedContacts, err := testContext.unifiedRepo.GetAllContacts(ctx, domain.UserID(testContext.User.UserID))
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(unifiedContacts))
	assert.Equal(t, "johndoe@gmail.com", unifiedContacts[0].EmailAddresses[0].Value)
	origin := domain.NewContactOrigin(domain.Google, source.ID, domain.ContactID("test-contact")).String()
	assert.Equal(t, unifiedContacts[0].Origins[origin], true)
}

func TestPullContactsHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testContext := newTestContext(t, ctrl)
	testContext.setupTestUser()
	t.Cleanup(func() { testContext.ClearDB() })

	testContext.CreateTestTopic(t, testContext.pubsubClient, conf.PullContactsSourceTopic())

	_, err := testContext.createTestContactSource(func(v *domain.ContactSource) { v.NextSyncAt = time.Now().Add(-5 * time.Minute) })
	assert.Equal(t, err, nil)
	_, err = testContext.createTestContactSource(func(v *domain.ContactSource) { v.NextSyncAt = time.Now().Add(5 * time.Minute) })
	assert.Equal(t, err, nil)

	params := &operations.PullContactsParams{
		Body: &models.PubsubMessage{
			Message: &models.Message{
				MessageID:   new(string),
				PublishTime: strfmt.DateTime(time.Now()),
			},
		},
	}
	_, err = testContext.client.PullContacts(params)
	assert.Equal(t, nil, err)
}

func TestUnifiedSync(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testContext := newTestContext(t, ctrl)
	testContext.setupTestUser()
	t.Cleanup(func() { testContext.ClearDB() })

	ctx := context.Background()
	userId := domain.UserID(testContext.User.UserID)

	contact := domain.Contact{
		EmailAddresses: []domain.EmailAddress{
			{
				Value: "johndoe@gmail.com",
			},
		},
		Names: []*domain.UserNames{
			{
				GivenName:  "John",
				FamilyName: "Doe",
			},
		},
		PhoneNumbers: []domain.PhoneNumber{
			{
				Value: "+919567123456",
			},
		},
	}

	sourceId := domain.ContactSourceID("test-source-1")
	contactId := domain.ContactID("test-contact-1")

	_, err := testContext.unifiedService.Add(ctx, userId, domain.Default, sourceId, contactId, contact)
	assert.Equal(t, nil, err)
	createContact, err := testContext.unifiedRepo.GetContactByOrigin(ctx, userId, domain.NewContactOrigin(domain.Default, sourceId, contactId))
	assert.Equal(t, nil, err)
	expectedTerms := sort.StringSlice([]string{"johndoe@gmail.com", "+919567123456", "john doe"})
	expectedTerms.Sort()
	results := sort.StringSlice(createContact.SearchTerms)
	results.Sort()
	assert.Equal(t, expectedTerms, results)

	suggestions, err := testContext.linkSuggestionRepo.GetAll(ctx, userId)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, len(suggestions))

	// new contact with same email should create a link suggestion
	contact = domain.Contact{
		EmailAddresses: []domain.EmailAddress{
			{
				Value: "johndoe@gmail.com",
			},
		},
		Names: []*domain.UserNames{
			{
				GivenName:  "Johnny",
				FamilyName: "",
			},
		},
		PhoneNumbers: []domain.PhoneNumber{
			{
				Value: "+919567000001",
			},
		},
	}
	_, err = testContext.unifiedService.Add(ctx, userId, domain.Default, sourceId, domain.ContactID("test-contact-2"), contact)
	assert.Equal(t, nil, err)

	suggestions, err = testContext.linkSuggestionRepo.GetAll(ctx, userId)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(suggestions))
	assert.Equal(t, domain.KeyEmail, suggestions[0].Key)
	assert.Equal(t, "johndoe@gmail.com", suggestions[0].Value)
	assert.Equal(t, 2, len(suggestions[0].Matches))

	// new contact with same email should append the contact to existing link suggestion
	contact = domain.Contact{
		EmailAddresses: []domain.EmailAddress{
			{
				Value: "johndoe@gmail.com",
			},
		},
		Names: []*domain.UserNames{
			{
				GivenName:  "Jo",
				FamilyName: "",
			},
		},
		PhoneNumbers: []domain.PhoneNumber{
			{
				Value: "+919567000002",
			},
		},
	}
	_, err = testContext.unifiedService.Add(ctx, userId, domain.Default, sourceId, domain.ContactID("test-contact-3"), contact)
	assert.Equal(t, nil, err)

	suggestions, err = testContext.linkSuggestionRepo.GetAll(ctx, userId)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(suggestions))
	assert.Equal(t, domain.KeyEmail, suggestions[0].Key)
	assert.Equal(t, "johndoe@gmail.com", suggestions[0].Value)
	assert.Equal(t, 3, len(suggestions[0].Matches))
}

func TestDeleteContactSourceJob_RemoveUnifiedFalse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testContext := newTestContext(t, ctrl)
	testContext.setupTestUser()
	t.Cleanup(func() { testContext.ClearDB() })

	ctx := context.Background()
	userId := domain.UserID(testContext.User.UserID)

	contact := domain.Contact{
		EmailAddresses: []domain.EmailAddress{
			{
				Value: "johndoe@gmail.com",
			},
		},
		Names: []*domain.UserNames{
			{
				GivenName:  "John",
				FamilyName: "Doe",
			},
		},
		PhoneNumbers: []domain.PhoneNumber{
			{
				Value: "+919567123456",
			},
		},
	}

	sourceId := domain.ContactSourceID("test-source-1")

	createdContact, err := testContext.defaultContactFirestore.SaveContact(ctx, userId, contact)
	assert.Equal(t, nil, err)

	createUnified, err := testContext.unifiedService.Add(ctx, userId, domain.Default, sourceId, createdContact.ID, contact)
	assert.Equal(t, nil, err)

	request := models.ContactSourceDeleted{
		ContactSourceID:           string(sourceId),
		RemoveContactsFromUnified: false,
		Source:                    string(domain.Default),
		UserID:                    string(testContext.User.UserID),
	}
	defaultContactSource := defaultcontactsource.NewDefaultContactSource(testContext.defaultContactFirestore)
	testContext.mockContactSourceProvider.EXPECT().Get(gomock.Any()).AnyTimes().Return(defaultContactSource)

	messageData, err := json.Marshal(request)
	assert.Equal(t, err, nil)

	params := &operations.ContactSourceCleanUpParams{
		Body: &models.PubsubMessage{
			Message: &models.Message{
				Data:        messageData,
				MessageID:   new(string),
				PublishTime: strfmt.DateTime(time.Now()),
			},
		},
	}
	_, err = testContext.client.ContactSourceCleanUp(params)
	assert.Equal(t, nil, err)

	// should not remove unified contact
	unified, err := testContext.unifiedRepo.GetContactByID(ctx, userId, createUnified.ID)
	assert.Equal(t, nil, err)
	assert.Equal(t, false, unified.Origins[string(domain.NewContactOrigin(domain.Default, sourceId, createdContact.ID))])

	// contact should be remove
	_, err = testContext.defaultContactFirestore.GetContactByID(ctx, userId, createdContact.ID)
	assert.Equal(t, true, status.Code(err) == codes.NotFound)
}

func TestDeleteContactSourceJob_RemoveUnifiedTrue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testContext := newTestContext(t, ctrl)
	testContext.setupTestUser()
	t.Cleanup(func() { testContext.ClearDB() })

	ctx := context.Background()
	userId := domain.UserID(testContext.User.UserID)

	contact := domain.Contact{
		EmailAddresses: []domain.EmailAddress{
			{
				Value: "johndoe@gmail.com",
			},
		},
		Names: []*domain.UserNames{
			{
				GivenName:  "John",
				FamilyName: "Doe",
			},
		},
		PhoneNumbers: []domain.PhoneNumber{
			{
				Value: "+919567123456",
			},
		},
	}

	sourceId := domain.ContactSourceID("test-source-1")

	createdContact, err := testContext.defaultContactFirestore.SaveContact(ctx, userId, contact)
	assert.Equal(t, nil, err)

	createUnified, err := testContext.unifiedService.Add(ctx, userId, domain.Default, sourceId, createdContact.ID, contact)
	assert.Equal(t, nil, err)

	request := models.ContactSourceDeleted{
		ContactSourceID:           string(sourceId),
		RemoveContactsFromUnified: true,
		Source:                    string(domain.Default),
		UserID:                    string(testContext.User.UserID),
	}
	defaultContactSource := defaultcontactsource.NewDefaultContactSource(testContext.defaultContactFirestore)
	testContext.mockContactSourceProvider.EXPECT().Get(gomock.Any()).AnyTimes().Return(defaultContactSource)

	messageData, err := json.Marshal(request)
	assert.Equal(t, err, nil)

	params := &operations.ContactSourceCleanUpParams{
		Body: &models.PubsubMessage{
			Message: &models.Message{
				Data:        messageData,
				MessageID:   new(string),
				PublishTime: strfmt.DateTime(time.Now()),
			},
		},
	}
	_, err = testContext.client.ContactSourceCleanUp(params)
	assert.Equal(t, nil, err)

	_, err = testContext.unifiedRepo.GetContactByID(ctx, userId, createUnified.ID)
	assert.Equal(t, true, status.Code(err) == codes.NotFound)

	// contact should be removed
	_, err = testContext.defaultContactFirestore.GetContactByID(ctx, userId, createdContact.ID)
	assert.Equal(t, true, status.Code(err) == codes.NotFound)
}
