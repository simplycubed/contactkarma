package routes

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/simplycubed/contactkarma/contacts/adapters/api"
	"github.com/simplycubed/contactkarma/contacts/adapters/firestore"
	"github.com/simplycubed/contactkarma/contacts/adapters/pubsub"
	"github.com/simplycubed/contactkarma/contacts/application"
	"github.com/simplycubed/contactkarma/contacts/conf"
	"github.com/simplycubed/contactkarma/contacts/domain"
	"github.com/simplycubed/contactkarma/contacts/gen/client/operations"
	"github.com/simplycubed/contactkarma/contacts/gen/mocks/mock_application"
	"github.com/simplycubed/contactkarma/contacts/test/testutils"
	"github.com/go-openapi/strfmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	googleOAuth2 "google.golang.org/api/oauth2/v1"
	"google.golang.org/api/option"
)

type contactSourceTestContext struct {
	*testing.T
	*testutils.Env
	client               operations.ClientService
	userRepo             *firestore.UserFirestore
	unifiedRepo          *firestore.UnifiedContactFirestore
	contactSourceRepo    *firestore.ContactSourceFirestore
	googleOAuthService   *mock_application.MockGoogleOAuthService
	contactSourceService *application.ContactSourceService
}

func newContactSourceTestContext(t *testing.T, ctrl *gomock.Controller) *contactSourceTestContext {
	// fake google auth credential

	credetialsOption := option.WithCredentials(&google.Credentials{})
	dbpool, _ := firestore.NewFirestoreClient(conf.ProjectID(), credetialsOption)

	googleOAuthService := mock_application.NewMockGoogleOAuthService(ctrl)
	userFirestore := firestore.NewUserFirestore(dbpool)
	userService := application.NewUserService(userFirestore)
	unifiedFirestore := firestore.NewUnifiedContactFirestore(dbpool)
	contactSourceFirestore := firestore.NewContactSourceFirestore(dbpool)
	pubsubClient := pubsub.New(conf.ProjectID(), credetialsOption)
	pullContactPublisher := pubsub.NewPullContactSourcePublisher(pubsubClient.Topic(conf.PullContactsSourceTopic()))
	contactSourceDeletedPublisher := pubsub.NewContactSourceDeletedPublisher(pubsubClient.Topic(conf.ContactSourceDeletedTopic()))
	linkSuggestionRepo := firestore.NewLinkSuggestionFirestore(dbpool)
	linkSuggestionService := application.NewLinkSuggestionService(unifiedFirestore, linkSuggestionRepo)
	contactLogRepo := firestore.NewContactLogFirestore(dbpool)
	unifiedContactService := application.NewUnifiedContactService(unifiedFirestore, linkSuggestionService, contactLogRepo)

	mockContactSourceProvider := mock_application.NewMockIContactSourceProvider(ctrl)
	contactSourceService := application.NewContactSourceService(googleOAuthService, contactSourceFirestore, pullContactPublisher, unifiedContactService, userFirestore, mockContactSourceProvider, unifiedFirestore, contactSourceDeletedPublisher)
	testApi := api.Create()
	ContactSource(testApi, contactSourceService)

	server := api.CreateServer(testApi)

	// setup env and client
	env := testutils.NewEnv(server.GetHandler(), userService)
	env.ClearDB() // clear db

	testCtxt := &contactSourceTestContext{}
	testCtxt.Env = env
	testCtxt.client = operations.New(env.Transport(), strfmt.Default)
	testCtxt.userRepo = userFirestore
	testCtxt.unifiedRepo = unifiedFirestore
	testCtxt.contactSourceRepo = contactSourceFirestore
	testCtxt.googleOAuthService = googleOAuthService
	testCtxt.contactSourceService = contactSourceService
	return testCtxt
}

func (ctx *contactSourceTestContext) setupTestUser() {
	err := ctx.AddTestUser()
	if err != nil {
		ctx.Fatal("Could not create user:", err)
	}
}

func createTestContactSource(contactSourceRepo *firestore.ContactSourceFirestore, userId domain.UserID, modify func(v *domain.ContactSource)) (created *domain.ContactSource, err error) {

	testSource := domain.ContactSource{
		UserID:       userId,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Source:       domain.Google,
		Email:        "test@gmail.com",
		GoogleUserId: "google-user-id",
		AccessToken:  "access-token",
		RefreshToken: "refresh-token",
		TokenExpiry:  time.Now().Add(5 * time.Second),
	}

	// add any modification to model if any
	if modify != nil {
		modify(&testSource)
	}

	created, err = contactSourceRepo.Create(context.Background(), userId, testSource)
	return
}

func TestInitGoogleContactSourceHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	testCtxt := newContactSourceTestContext(t, ctrl)
	testCtxt.setupTestUser()
	defer t.Cleanup(func() { testCtxt.ClearDB() })

	testCtxt.googleOAuthService.EXPECT().GetRedirectUrl(gomock.Any()).Return("http://test-url", nil)

	params := &operations.InitGoogleContactSourceParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
	}
	resp, err := testCtxt.client.InitGoogleContactSource(params)
	assert.Equal(t, true, err == nil)
	assert.Equal(t, "http://test-url", resp.Payload.URL)
	t.Cleanup(func() { testCtxt.ClearDB() })
}
func TestLinkGoogleContactSourceHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	testCtxt := newContactSourceTestContext(t, ctrl)
	testCtxt.setupTestUser()
	defer t.Cleanup(func() { testCtxt.ClearDB() })

	credetialsOption := option.WithCredentials(&google.Credentials{})
	pubsubClient := pubsub.New(conf.ProjectID(), credetialsOption)
	testCtxt.CreateTestTopic(t, pubsubClient, conf.PullContactsSourceTopic())

	// set mock expectations
	token := &oauth2.Token{
		AccessToken:  "test-access-token",
		TokenType:    "",
		RefreshToken: "",
		Expiry:       time.Now(),
	}
	testCtxt.googleOAuthService.EXPECT().GetToken(gomock.Any(), "test-auth-code").Return(token, nil)

	info := &googleOAuth2.Tokeninfo{
		Email:         "test@gmail.com",
		EmailVerified: true,
		UserId:        "test-user-id",
		VerifiedEmail: false,
	}
	testCtxt.googleOAuthService.EXPECT().GetUserInfo(gomock.Any(), "test-access-token").Return(info, nil)

	params := &operations.LinkGoogleContactSourceParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		Body: operations.LinkGoogleContactSourceBody{
			AuthCode: "test-auth-code",
		},
	}
	_, err := testCtxt.client.LinkGoogleContactSource(params)
	assert.Equal(t, nil, err)

	updatedUser, err := testCtxt.userRepo.GetUserByID(context.Background(), domain.UserID(testCtxt.User.UserID))
	assert.Equal(t, nil, err)
	assert.Equal(t, int64(1), updatedUser.Quota.TotalContactSources)
	t.Cleanup(func() { testCtxt.ClearDB() })
}

func TestLinkGoogleContactSourceHandler_LimitCheck(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	testCtxt := newContactSourceTestContext(t, ctrl)
	testCtxt.setupTestUser()
	defer t.Cleanup(func() { testCtxt.ClearDB() })
	testCtxt.SetRole(domain.RoleContactSync)
	testCtxt.userRepo.UpdateUser(context.Background(), domain.UserID(testCtxt.User.UserID), domain.User{
		Quota: &domain.Quota{
			TotalContactSources: 10,
		},
	})
	params := &operations.LinkGoogleContactSourceParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		Body: operations.LinkGoogleContactSourceBody{
			AuthCode: "test-auth-code",
		},
	}
	_, err := testCtxt.client.LinkGoogleContactSource(params)
	assert.Equal(t, true, err != nil)
	assert.Equal(t, true, strings.Contains(err.Error(), domain.ErrContactSourcesLimitReached.Error()))
	t.Cleanup(func() { testCtxt.ClearDB() })
}

func TestGetContactSourcesHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	testCtxt := newContactSourceTestContext(t, ctrl)
	testCtxt.setupTestUser()
	defer t.Cleanup(func() { testCtxt.ClearDB() })

	userId := domain.UserID(testCtxt.User.UserID)
	_, err := createTestContactSource(testCtxt.contactSourceRepo, userId, nil)
	assert.Equal(t, err, nil)

	_, err = createTestContactSource(testCtxt.contactSourceRepo, userId, func(v *domain.ContactSource) {
		v.Email = "secondary-email@gmail.com"
	})
	assert.Equal(t, err, nil)

	params := &operations.GetContactSourcesParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
	}
	response, err := testCtxt.client.GetContactSources(params)
	assert.Equal(t, err, nil)
	assert.Equal(t, 2, len(response.Payload))
}

func TestGetContactSourcesHandler_Pagination(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	testCtxt := newContactSourceTestContext(t, ctrl)
	testCtxt.setupTestUser()
	defer t.Cleanup(func() { testCtxt.ClearDB() })
	userId := domain.UserID(testCtxt.User.UserID)

	ids := []string{}
	for i := 0; i < 10; i++ {
		created, err := createTestContactSource(testCtxt.contactSourceRepo, userId, func(v *domain.ContactSource) { v.Email = fmt.Sprintf("tagm%d@gmail.co", i) })
		assert.Equal(t, err, nil)
		ids = append(ids, string(created.ID))
	}
	sort.StringSlice(ids).Sort()

	var limit int64 = 5
	params := &operations.GetContactSourcesParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		Limit:                  &limit,
	}
	firstBatchResponse, err := testCtxt.client.GetContactSources(params)
	assert.Equal(t, err, nil)
	assert.Equal(t, int(limit), len(firstBatchResponse.Payload))
	for i := 0; i < len(firstBatchResponse.Payload); i++ {
		assert.Equal(t, ids[i], firstBatchResponse.Payload[i].ID)
	}

	// next batch
	params = &operations.GetContactSourcesParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		Limit:                  &limit,
		LastDocumentID:         &firstBatchResponse.Payload[len(firstBatchResponse.Payload)-1].ID,
	}
	secondBatchResponse, err := testCtxt.client.GetContactSources(params)
	assert.Equal(t, err, nil)
	assert.Equal(t, int(limit), len(secondBatchResponse.Payload))
	for i := 0; i < len(secondBatchResponse.Payload); i++ {
		assert.Equal(t, ids[i+int(limit)], secondBatchResponse.Payload[i].ID)
	}
}

func TestDeleteContactSourceHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	testCtxt := newContactSourceTestContext(t, ctrl)
	testCtxt.setupTestUser()
	defer t.Cleanup(func() { testCtxt.ClearDB() })
	userId := domain.UserID(testCtxt.User.UserID)

	credetialsOption := option.WithCredentials(&google.Credentials{})
	pubsubClient := pubsub.New(conf.ProjectID(), credetialsOption)
	testCtxt.CreateTestTopic(t, pubsubClient, conf.ContactSourceDeletedTopic())

	source, err := createTestContactSource(testCtxt.contactSourceRepo, userId, nil)
	assert.Equal(t, err, nil)

	testCtxt.userRepo.UpdateUser(context.Background(), userId, domain.User{Quota: &domain.Quota{TotalContactSources: 1}})

	params := &operations.DeleteContactSourceParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		SourceID:               string(source.ID),
		Body: operations.DeleteContactSourceBody{
			RemoveFromUnified: false,
		},
	}
	_, err = testCtxt.client.DeleteContactSource(params)
	assert.Equal(t, err, nil)

	user, err := testCtxt.userRepo.GetUserByID(context.Background(), userId)
	assert.Equal(t, err, nil)
	assert.Equal(t, int64(0), user.Quota.TotalContactSources)
}
