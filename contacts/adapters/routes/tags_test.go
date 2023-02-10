package routes

import (
	"fmt"
	"sort"
	"testing"
	"time"

	"github.com/simplycubed/contactkarma/contacts/adapters/api"
	"github.com/simplycubed/contactkarma/contacts/adapters/firestore"
	"github.com/simplycubed/contactkarma/contacts/application"
	"github.com/simplycubed/contactkarma/contacts/conf"
	"github.com/simplycubed/contactkarma/contacts/domain"
	"github.com/simplycubed/contactkarma/contacts/gen/client/operations"
	"github.com/simplycubed/contactkarma/contacts/gen/models"
	"github.com/simplycubed/contactkarma/contacts/test/testutils"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

type tagTestContext struct {
	*testing.T
	*testutils.Env
	client      operations.ClientService
	userRepo    *firestore.UserFirestore
	unifiedRepo *firestore.UnifiedContactFirestore
	tagRepo     *firestore.TagFirestore
}

func newTagTestContext(t *testing.T) *tagTestContext {
	// fake google auth credential
	credetialsOption := option.WithCredentials(&google.Credentials{})
	dbpool, _ := firestore.NewFirestoreClient(conf.ProjectID(), credetialsOption)

	userFirestore := firestore.NewUserFirestore(dbpool)
	userService := application.NewUserService(userFirestore)
	unifiedFirestore := firestore.NewUnifiedContactFirestore(dbpool)
	tagFirestore := firestore.NewTagFirestore(dbpool)
	tagService := application.NewTagService(tagFirestore)
	testApi := api.Create()
	Tags(testApi, tagService)

	server := api.CreateServer(testApi)

	// setup env and client
	env := testutils.NewEnv(server.GetHandler(), userService)
	env.ClearDB() // clear db

	testCtxt := &tagTestContext{}
	testCtxt.Env = env
	testCtxt.client = operations.New(env.Transport(), strfmt.Default)
	testCtxt.userRepo = userFirestore
	testCtxt.unifiedRepo = unifiedFirestore
	testCtxt.tagRepo = tagFirestore
	return testCtxt
}

func (ctx *tagTestContext) setupTestUser() {
	err := ctx.AddTestUser()
	if err != nil {
		ctx.Fatal("Could not create user:", err)
	}
}

func TestPostContactTagHandler(t *testing.T) {
	testCtxt := newTagTestContext(t)
	testCtxt.setupTestUser()
	defer t.Cleanup(func() { testCtxt.ClearDB() })

	userId := domain.UserID(testCtxt.User.UserID)
	contact, err := createTestUnified(testCtxt.unifiedRepo, userId, nil)
	assert.Equal(t, nil, err)

	params := &operations.PostContactTagParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		Body: &models.Tag{
			CreatedAt: strfmt.DateTime(time.Now()),
			TagName:   "test tag",
		},
		UnifiedID: contact.ID.String(),
	}
	_, err = testCtxt.client.PostContactTag(params)
	assert.Equal(t, nil, err)
}

func TestPatchContactTagHandler(t *testing.T) {
	testCtxt := newTagTestContext(t)
	testCtxt.setupTestUser()
	defer t.Cleanup(func() { testCtxt.ClearDB() })

	userId := domain.UserID(testCtxt.User.UserID)
	contact, err := createTestUnified(testCtxt.unifiedRepo, userId, nil)
	assert.Equal(t, nil, err)

	tag, err := createTestContactTag(testCtxt.tagRepo, userId, domain.UnifiedId(contact.ID), nil)
	assert.Equal(t, nil, err)

	params := &operations.PatchContactTagParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		Body: &models.Tag{
			CreatedAt: strfmt.DateTime(time.Now()),
			TagName:   "updated tag",
		},
		UnifiedID: contact.ID.String(),
		TagID:     tag.ID.String(),
	}
	updated, err := testCtxt.client.PatchContactTag(params)
	assert.Equal(t, nil, err)
	assert.Equal(t, "updated tag", updated.Payload.TagName)
}

func TestDeleteContactTagHandler(t *testing.T) {
	testCtxt := newTagTestContext(t)
	testCtxt.setupTestUser()
	defer t.Cleanup(func() { testCtxt.ClearDB() })

	userId := domain.UserID(testCtxt.User.UserID)
	contact, err := createTestUnified(testCtxt.unifiedRepo, userId, nil)
	assert.Equal(t, nil, err)

	tag, err := createTestContactTag(testCtxt.tagRepo, userId, domain.UnifiedId(contact.ID), nil)
	assert.Equal(t, nil, err)

	params := &operations.DeleteContactTagParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		UnifiedID:              contact.ID.String(),
		TagID:                  tag.ID.String(),
	}
	_, err = testCtxt.client.DeleteContactTag(params)
	assert.Equal(t, nil, err)

	//

}

func TestGetContactTagsHandler(t *testing.T) {
	testCtxt := newTagTestContext(t)
	testCtxt.setupTestUser()
	defer t.Cleanup(func() { testCtxt.ClearDB() })

	userId := domain.UserID(testCtxt.User.UserID)
	contact, err := createTestUnified(testCtxt.unifiedRepo, userId, nil)
	assert.Equal(t, nil, err)

	_, err = createTestContactTag(testCtxt.tagRepo, userId, domain.UnifiedId(contact.ID), nil)
	assert.Equal(t, nil, err)

	params := &operations.GetContactTagsParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		UnifiedID:              contact.ID.String(),
	}
	tags, err := testCtxt.client.GetContactTags(params)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(tags.Payload))
}

func TestGetContactTagsHandler_Pagination(t *testing.T) {
	testCtxt := newTagTestContext(t)
	testCtxt.setupTestUser()
	defer t.Cleanup(func() { testCtxt.ClearDB() })

	userId := domain.UserID(testCtxt.User.UserID)
	contact, err := createTestUnified(testCtxt.unifiedRepo, userId, nil)
	assert.Equal(t, nil, err)

	ids := []string{}
	for i := 0; i < 10; i++ {
		created, err := createTestContactTag(testCtxt.tagRepo, userId, domain.UnifiedId(contact.ID), func(v *domain.Tag) { v.TagName = fmt.Sprintf("Tag %d", i) })
		assert.Equal(t, err, nil)
		ids = append(ids, created.ID.String())
	}
	sort.StringSlice(ids).Sort()

	var limit int64 = 5
	params := &operations.GetContactTagsParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		UnifiedID:              contact.ID.String(),
		Limit:                  &limit,
	}
	firstBatchResponse, err := testCtxt.client.GetContactTags(params)
	assert.Equal(t, err, nil)
	assert.Equal(t, int(limit), len(firstBatchResponse.Payload))
	for i := 0; i < len(firstBatchResponse.Payload); i++ {
		assert.Equal(t, ids[i], firstBatchResponse.Payload[i].ID)
	}

	// next batch
	params = &operations.GetContactTagsParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		UnifiedID:              contact.ID.String(),
		Limit:                  &limit,
		LastDocumentID:         &firstBatchResponse.Payload[len(firstBatchResponse.Payload)-1].ID,
	}
	secondBatchResponse, err := testCtxt.client.GetContactTags(params)
	assert.Equal(t, err, nil)
	assert.Equal(t, int(limit), len(secondBatchResponse.Payload))
	for i := 0; i < len(secondBatchResponse.Payload); i++ {
		assert.Equal(t, ids[i+int(limit)], secondBatchResponse.Payload[i].ID)
	}
}
