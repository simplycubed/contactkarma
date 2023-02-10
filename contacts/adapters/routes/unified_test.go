package routes

import (
	"context"
	"fmt"
	"sort"
	"testing"
	"time"

	"github.com/simplycubed/contactkarma/contacts/adapters/api"
	"github.com/simplycubed/contactkarma/contacts/adapters/firestore"
	"github.com/simplycubed/contactkarma/contacts/adapters/typesense"
	"github.com/simplycubed/contactkarma/contacts/application"
	"github.com/simplycubed/contactkarma/contacts/conf"
	"github.com/simplycubed/contactkarma/contacts/domain"
	"github.com/simplycubed/contactkarma/contacts/gen/client/operations"
	"github.com/simplycubed/contactkarma/contacts/gen/models"
	"github.com/simplycubed/contactkarma/contacts/test/testutils"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	tsense "github.com/typesense/typesense-go/typesense"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type unifiedTestContext struct {
	*testing.T
	*testutils.Env
	client             operations.ClientService
	userRepo           *firestore.UserFirestore
	unifiedRepo        *firestore.UnifiedContactFirestore
	linkSuggestionRepo *firestore.LinkSuggestionFirestore
	typesenseClient    *tsense.Client
	contactSearch      *typesense.ContactSearch
}

func newUnifiedTestContext(t *testing.T) *unifiedTestContext {
	// fake google auth credential
	credetialsOption := option.WithCredentials(&google.Credentials{})
	dbpool, _ := firestore.NewFirestoreClient(conf.ProjectID(), credetialsOption)

	userFirestore := firestore.NewUserFirestore(dbpool)
	userService := application.NewUserService(userFirestore)
	unifiedFirestore := firestore.NewUnifiedContactFirestore(dbpool)
	linkSuggestionRepo := firestore.NewLinkSuggestionFirestore(dbpool)
	linkSuggestionService := application.NewLinkSuggestionService(unifiedFirestore, linkSuggestionRepo)
	contactLogRepo := firestore.NewContactLogFirestore(dbpool)
	typesenseClient := typesense.New(conf.TypesenseHost(), conf.TypesenseApiKey())
	contactSearch := typesense.NewContactSearch(typesenseClient)
	unifiedService := application.NewUnifiedContactService(unifiedFirestore, linkSuggestionService, contactLogRepo)
	contactSearchService := application.NewContactSearchService(contactSearch, unifiedFirestore)
	testApi := api.Create()
	Unified(testApi, unifiedService, linkSuggestionService, contactSearchService)

	server := api.CreateServer(testApi)

	// setup env and client
	env := testutils.NewEnv(server.GetHandler(), userService)
	env.ClearDB() // clear db

	testCtxt := &unifiedTestContext{}
	testCtxt.Env = env
	testCtxt.client = operations.New(env.Transport(), strfmt.Default)
	testCtxt.userRepo = userFirestore
	testCtxt.unifiedRepo = unifiedFirestore
	testCtxt.linkSuggestionRepo = linkSuggestionRepo
	testCtxt.typesenseClient = typesenseClient
	testCtxt.contactSearch = contactSearch
	return testCtxt
}

func (ctx *unifiedTestContext) setupTestUser() {
	err := ctx.AddTestUser()
	if err != nil {
		ctx.Fatal("Could not create user:", err)
	}
}

func createTestLinkSuggestion(linkSuggestionRepo *firestore.LinkSuggestionFirestore, userId domain.UserID, modify func(v *domain.LinkSuggestion)) (created *domain.LinkSuggestion, err error) {
	testLinkSuggestion := domain.LinkSuggestion{
		Key:     domain.KeyEmail,
		Value:   "test@gmail.com",
		Matches: []domain.LinkMatch{},
	}

	// add any modification to model if any
	if modify != nil {
		modify(&testLinkSuggestion)
	}

	created, err = linkSuggestionRepo.Create(context.Background(), userId, testLinkSuggestion)
	return
}

func TestGetUnifiedContactHandler(t *testing.T) {
	testCtxt := newUnifiedTestContext(t)
	testCtxt.setupTestUser()
	defer t.Cleanup(func() { testCtxt.ClearDB() })
	userId := domain.UserID(testCtxt.User.UserID)

	_, err := createTestUnified(testCtxt.unifiedRepo, userId, nil)
	assert.Equal(t, err, nil)
	params := &operations.GetUnifiedContactsParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
	}
	response, err := testCtxt.client.GetUnifiedContacts(params)
	assert.Equal(t, err, nil)
	assert.Equal(t, 1, len(response.Payload))
}

func TestGetUnifiedContactHandler_Pagination(t *testing.T) {
	testCtxt := newUnifiedTestContext(t)
	testCtxt.setupTestUser()
	defer t.Cleanup(func() { testCtxt.ClearDB() })
	userId := domain.UserID(testCtxt.User.UserID)

	// test pagination
	ids := []string{}
	for i := 0; i < 10; i++ {
		contact, err := createTestUnified(testCtxt.unifiedRepo, userId, func(v *domain.Unified) { v.Names[0].DisplayName = fmt.Sprintf("John %d", i) })
		assert.Equal(t, err, nil)
		ids = append(ids, string(contact.ID))
	}
	sort.StringSlice(ids).Sort()

	var limit int64 = 5
	params := &operations.GetUnifiedContactsParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		Limit:                  &limit,
	}
	firstBatchResponse, err := testCtxt.client.GetUnifiedContacts(params)
	assert.Equal(t, err, nil)
	assert.Equal(t, int(limit), len(firstBatchResponse.Payload))
	for i := 0; i < len(firstBatchResponse.Payload); i++ {
		assert.Equal(t, ids[i], firstBatchResponse.Payload[i].ID)
	}

	// next batch
	params = &operations.GetUnifiedContactsParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		Limit:                  &limit,
		LastDocumentID:         &firstBatchResponse.Payload[len(firstBatchResponse.Payload)-1].ID,
	}
	secondBatchResponse, err := testCtxt.client.GetUnifiedContacts(params)
	assert.Equal(t, err, nil)
	assert.Equal(t, int(limit), len(secondBatchResponse.Payload))
	for i := 0; i < len(secondBatchResponse.Payload); i++ {
		assert.Equal(t, ids[i+int(limit)], secondBatchResponse.Payload[i].ID)
	}

}

func TestPendingFollowUpsHandler(t *testing.T) {
	testCtxt := newUnifiedTestContext(t)
	testCtxt.setupTestUser()
	defer t.Cleanup(func() { testCtxt.ClearDB() })
	userId := domain.UserID(testCtxt.User.UserID)

	_, err := createTestUnified(testCtxt.unifiedRepo, userId, func(contact *domain.Unified) {
		now := time.Now().Add(24 * time.Hour)
		contact.NextContact = &now
	})
	assert.Equal(t, err, nil)

	_, err = createTestUnified(testCtxt.unifiedRepo, userId, func(contact *domain.Unified) {
		contact.Names = []*domain.UserNames{{
			DisplayName: "Jane Doe",
		}}
		contact.Birthdays = []domain.Birthday{{
			Date: "27-06-2000",
		}}
		contact.Genders = []domain.Gender{{
			Value: "female",
		}}
		now := time.Now().Add(-5 * time.Second)
		contact.NextContact = &now
	})
	assert.Equal(t, err, nil)

	params := &operations.GetPendingFollowUpsParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
	}
	followUps, err := testCtxt.client.GetPendingFollowUps(params)
	assert.Equal(t, err, nil)
	assert.Equal(t, 1, len(followUps.Payload))
	assert.Equal(t, "Jane Doe", followUps.Payload[0].Names[0].DisplayName)

}

func TestPendingFollowUps_Pagination(t *testing.T) {
	testCtxt := newUnifiedTestContext(t)
	testCtxt.setupTestUser()
	defer t.Cleanup(func() { testCtxt.ClearDB() })
	userId := domain.UserID(testCtxt.User.UserID)

	// test pagination
	ids := []string{}
	for i := 0; i < 10; i++ {
		contact, err := createTestUnified(testCtxt.unifiedRepo, userId, func(v *domain.Unified) {
			v.Names[0].DisplayName = fmt.Sprintf("John %d", i)
			// set past date
			time := time.Date(2020, 1, 1, i, 0, 0, 0, time.UTC)
			v.NextContact = &time
		})
		assert.Equal(t, err, nil)
		ids = append(ids, string(contact.ID))
	}
	var limit int64 = 5
	params := &operations.GetPendingFollowUpsParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		Limit:                  &limit,
	}
	firstBatchResponse, err := testCtxt.client.GetPendingFollowUps(params)
	assert.Equal(t, err, nil)
	assert.Equal(t, int(limit), len(firstBatchResponse.Payload))
	for i := 0; i < len(firstBatchResponse.Payload); i++ {
		assert.Equal(t, ids[i], firstBatchResponse.Payload[i].ID)
	}

	// next batch
	params = &operations.GetPendingFollowUpsParams{
		XApigatewayAPIUserinfo:  testCtxt.GetXApigatewayAPIUserinfo(),
		Limit:                   &limit,
		LastDocumentID:          &firstBatchResponse.Payload[len(firstBatchResponse.Payload)-1].ID,
		LastDocumentNextContact: &firstBatchResponse.Payload[len(firstBatchResponse.Payload)-1].NextContact,
	}
	secondBatchResponse, err := testCtxt.client.GetPendingFollowUps(params)
	assert.Equal(t, err, nil)
	assert.Equal(t, int(limit), len(secondBatchResponse.Payload))
	for i := 0; i < len(secondBatchResponse.Payload); i++ {
		assert.Equal(t, ids[i+int(limit)], secondBatchResponse.Payload[i].ID)
	}

}

func TestGetRecentContactsHandler(t *testing.T) {
	testCtxt := newUnifiedTestContext(t)
	testCtxt.setupTestUser()
	defer t.Cleanup(func() { testCtxt.ClearDB() })
	userId := domain.UserID(testCtxt.User.UserID)

	_, err := createTestUnified(testCtxt.unifiedRepo, userId, func(contact *domain.Unified) {
		now := time.Now().Add(-4 * 24 * time.Hour)
		contact.LastContact = &now
	})
	assert.Equal(t, err, nil)

	_, err = createTestUnified(testCtxt.unifiedRepo, userId, func(contact *domain.Unified) {
		contact.Names[0].DisplayName = "Mark Doe"
		now := time.Now().Add(-2 * 24 * time.Hour)
		contact.LastContact = &now // contacted 2 days before
	})
	assert.Equal(t, err, nil)

	_, err = createTestUnified(testCtxt.unifiedRepo, userId, func(contact *domain.Unified) {
		contact.Names[0].DisplayName = "Mevin Doe"
		now := time.Now().Add(-1 * 24 * time.Hour)
		contact.LastContact = &now // contacted 1 day before
	})
	assert.Equal(t, err, nil)

	var maxDays int64 = 3
	params := &operations.GetRecentContactsParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		MaxDays:                &maxDays,
	}
	recent, err := testCtxt.client.GetRecentContacts(params)
	assert.Equal(t, err, nil)
	assert.Equal(t, 2, len(recent.Payload))
	assert.Equal(t, "Mevin Doe", recent.Payload[0].Names[0].DisplayName) // mevin is most recent contact
	assert.Equal(t, "Mark Doe", recent.Payload[1].Names[0].DisplayName)
}

func TestGetRecentContacts_Pagination(t *testing.T) {
	testCtxt := newUnifiedTestContext(t)
	testCtxt.setupTestUser()
	defer t.Cleanup(func() { testCtxt.ClearDB() })
	userId := domain.UserID(testCtxt.User.UserID)
	// test pagination
	ids := []string{}
	for i := 0; i < 10; i++ {
		contact, err := createTestUnified(testCtxt.unifiedRepo, userId, func(v *domain.Unified) {
			v.Names[0].DisplayName = fmt.Sprintf("John %d", i)
			now := time.Now()
			time := time.Date(now.Year(), now.Month(), now.Day(), now.Hour()+i, 0, 0, 0, time.UTC)
			v.LastContact = &time
		})
		assert.Equal(t, err, nil)
		ids = append([]string{string(contact.ID)}, ids...)
	}
	var limit int64 = 5
	params := &operations.GetRecentContactsParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		Limit:                  &limit,
	}
	firstBatchResponse, err := testCtxt.client.GetRecentContacts(params)
	assert.Equal(t, err, nil)
	assert.Equal(t, int(limit), len(firstBatchResponse.Payload))
	for i := 0; i < len(firstBatchResponse.Payload); i++ {
		assert.Equal(t, ids[i], firstBatchResponse.Payload[i].ID)
	}

	// next batch
	params = &operations.GetRecentContactsParams{
		XApigatewayAPIUserinfo:  testCtxt.GetXApigatewayAPIUserinfo(),
		Limit:                   &limit,
		LastDocumentID:          &firstBatchResponse.Payload[len(firstBatchResponse.Payload)-1].ID,
		LastDocumentLastContact: &firstBatchResponse.Payload[len(firstBatchResponse.Payload)-1].LastContact,
	}
	secondBatchResponse, err := testCtxt.client.GetRecentContacts(params)
	assert.Equal(t, err, nil)
	assert.Equal(t, int(limit), len(secondBatchResponse.Payload))
	for i := 0; i < len(secondBatchResponse.Payload); i++ {
		assert.Equal(t, ids[i+int(limit)], secondBatchResponse.Payload[i].ID)
	}

}

func TestGetLinkSuggestionsHandler(t *testing.T) {
	testCtxt := newUnifiedTestContext(t)
	testCtxt.setupTestUser()
	defer t.Cleanup(func() { testCtxt.ClearDB() })
	userId := domain.UserID(testCtxt.User.UserID)

	_, err := createTestLinkSuggestion(testCtxt.linkSuggestionRepo, userId, nil)
	assert.Equal(t, err, nil)
	params := &operations.GetLinkSuggestionsParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
	}
	response, err := testCtxt.client.GetLinkSuggestions(params)
	assert.Equal(t, err, nil)
	assert.Equal(t, 1, len(response.Payload))
}

func TestGetLinkSuggestionsHandler_Pagination(t *testing.T) {
	testCtxt := newUnifiedTestContext(t)
	testCtxt.setupTestUser()
	defer t.Cleanup(func() { testCtxt.ClearDB() })
	userId := domain.UserID(testCtxt.User.UserID)

	// test pagination
	ids := []string{}
	for i := 0; i < 10; i++ {
		contact, err := createTestLinkSuggestion(testCtxt.linkSuggestionRepo, userId, func(v *domain.LinkSuggestion) { v.Value = fmt.Sprintf("John %d", i) })
		assert.Equal(t, err, nil)
		ids = append(ids, string(contact.ID))
	}
	sort.StringSlice(ids).Sort()

	var limit int64 = 5
	params := &operations.GetLinkSuggestionsParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		Limit:                  &limit,
	}
	firstBatchResponse, err := testCtxt.client.GetLinkSuggestions(params)
	assert.Equal(t, err, nil)
	assert.Equal(t, int(limit), len(firstBatchResponse.Payload))
	for i := 0; i < len(firstBatchResponse.Payload); i++ {
		assert.Equal(t, ids[i], firstBatchResponse.Payload[i].ID)
	}

	// next batch
	params = &operations.GetLinkSuggestionsParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		Limit:                  &limit,
		LastDocumentID:         &firstBatchResponse.Payload[len(firstBatchResponse.Payload)-1].ID,
	}
	secondBatchResponse, err := testCtxt.client.GetLinkSuggestions(params)
	assert.Equal(t, err, nil)
	assert.Equal(t, int(limit), len(secondBatchResponse.Payload))
	for i := 0; i < len(secondBatchResponse.Payload); i++ {
		assert.Equal(t, ids[i+int(limit)], secondBatchResponse.Payload[i].ID)
	}

}

func TestApplyLinkSuggestionHandler(t *testing.T) {
	testCtxt := newUnifiedTestContext(t)
	testCtxt.setupTestUser()
	defer t.Cleanup(func() { testCtxt.ClearDB() })
	userId := domain.UserID(testCtxt.User.UserID)

	contact1, err := createTestUnified(testCtxt.unifiedRepo, userId, func(v *domain.Unified) { v.EmailAddresses = []domain.EmailAddress{{Value: "john@gmail.com"}} })
	assert.Equal(t, err, nil)
	contact2, err := createTestUnified(testCtxt.unifiedRepo, userId, func(v *domain.Unified) { v.EmailAddresses = []domain.EmailAddress{{Value: "john@gmail.com"}} })
	assert.Equal(t, err, nil)
	suggestion, err := createTestLinkSuggestion(testCtxt.linkSuggestionRepo, userId, func(v *domain.LinkSuggestion) {
		v.Matches = []domain.LinkMatch{{UnifiedId: contact1.ID}, {UnifiedId: contact2.ID}}
	})
	assert.Equal(t, err, nil)
	params := &operations.ApplyLinkSuggestionParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		SuggestionID:           string(suggestion.ID),
		Body: operations.ApplyLinkSuggestionBody{
			UnifiedIds: []string{string(contact1.ID), string(contact2.ID)},
		},
	}
	ctx := context.Background()
	_, err = testCtxt.client.ApplyLinkSuggestion(params)
	assert.Equal(t, err, nil)
	_, err = testCtxt.linkSuggestionRepo.GetById(ctx, domain.UserID(testCtxt.User.UserID), suggestion.ID)
	assert.Equal(t, true, err != nil)
	assert.Equal(t, true, status.Code(err) == codes.NotFound)

	_, err = testCtxt.unifiedRepo.GetContactByID(ctx, domain.UserID(testCtxt.User.UserID), contact1.ID)
	assert.Equal(t, err, nil)

	_, err = testCtxt.unifiedRepo.GetContactByID(ctx, domain.UserID(testCtxt.User.UserID), contact2.ID)
	assert.Equal(t, true, err != nil)
	assert.Equal(t, true, status.Code(err) == codes.NotFound)

}

func TestSearchContactsHandler(t *testing.T) {
	testCtxt := newUnifiedTestContext(t)
	testCtxt.setupTestUser()
	defer t.Cleanup(func() { testCtxt.ClearDB() })
	userId := domain.UserID(testCtxt.User.UserID)

	unified, err := createTestUnified(testCtxt.unifiedRepo, userId, nil)
	assert.Equal(t, err, nil)

	err = testCtxt.contactSearch.CreateCollection()
	assert.Equal(t, true, err == nil)
	defer func() {
		testCtxt.typesenseClient.Collection(typesense.GetContactCollection()).Delete()
	}()
	unifiedSearch := domain.UnifiedSearch{
		ID:          domain.UnifiedSearchId(string(unified.ID) + "_" + string(userId)),
		UnifiedId:   unified.ID,
		UserId:      userId,
		DisplayName: "John Doe",
		Category:    domain.A,
		Score:       100,
	}
	_, err = testCtxt.contactSearch.Upsert(userId, unifiedSearch)
	assert.Equal(t, true, err == nil)

	params := &operations.SearchUserContactParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		Body: &models.SearchContactDto{
			Query: "joh",
			Sort: []*models.SearchSort{{
				Field: "score",
				Order: "desc",
			}},
			Filters: []*models.SearchFilter{{
				Field:    "category",
				Operator: "=",
				Value:    "A",
			}},
		},
	}
	response, err := testCtxt.client.SearchUserContact(params)
	assert.Equal(t, err, nil)
	assert.Equal(t, 1, len(response.Payload))
}
