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

type noteTestContext struct {
	*testing.T
	*testutils.Env
	client      operations.ClientService
	userRepo    *firestore.UserFirestore
	unifiedRepo *firestore.UnifiedContactFirestore
	noteRepo    *firestore.NoteFirestore
}

func newNoteTestContext(t *testing.T) *noteTestContext {
	// fake google auth credential
	credetialsOption := option.WithCredentials(&google.Credentials{})
	dbpool, _ := firestore.NewFirestoreClient(conf.ProjectID(), credetialsOption)

	userFirestore := firestore.NewUserFirestore(dbpool)
	userService := application.NewUserService(userFirestore)
	unifiedFirestore := firestore.NewUnifiedContactFirestore(dbpool)
	noteRepo := firestore.NewNoteFirestore(dbpool)
	noteService := application.NewNoteService(noteRepo)
	testApi := api.Create()
	Notes(testApi, noteService)

	server := api.CreateServer(testApi)

	// setup env and client
	env := testutils.NewEnv(server.GetHandler(), userService)
	env.ClearDB() // clear db

	testCtxt := &noteTestContext{}
	testCtxt.Env = env
	testCtxt.client = operations.New(env.Transport(), strfmt.Default)
	testCtxt.userRepo = userFirestore
	testCtxt.unifiedRepo = unifiedFirestore
	testCtxt.noteRepo = noteRepo
	return testCtxt
}

func (ctx *noteTestContext) setupTestUser() {
	err := ctx.AddTestUser()
	if err != nil {
		ctx.Fatal("Could not create user:", err)
	}
}

func TestPostContactNoteHandler(t *testing.T) {
	testCtxt := newNoteTestContext(t)
	testCtxt.setupTestUser()
	defer t.Cleanup(func() { testCtxt.ClearDB() })
	userId := domain.UserID(testCtxt.User.UserID)
	contact, err := createTestUnified(testCtxt.unifiedRepo, userId, nil)
	assert.Equal(t, nil, err)

	params := &operations.PostContactNoteParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		Body: &models.Note{
			CreatedAt: strfmt.DateTime(time.Now()),
			Note:      "test note",
		},
		UnifiedID: contact.ID.String(),
	}
	_, err = testCtxt.client.PostContactNote(params)
	assert.Equal(t, nil, err)
}

func TestPatchContactNoteHandler(t *testing.T) {
	testCtxt := newNoteTestContext(t)
	testCtxt.setupTestUser()
	defer t.Cleanup(func() { testCtxt.ClearDB() })

	userId := domain.UserID(testCtxt.User.UserID)
	contact, err := createTestUnified(testCtxt.unifiedRepo, userId, nil)
	assert.Equal(t, nil, err)

	note, err := createTestContactNote(testCtxt.noteRepo, userId, domain.UnifiedId(contact.ID), nil)
	assert.Equal(t, nil, err)

	params := &operations.PatchContactNoteParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		Body: &models.Note{
			CreatedAt: strfmt.DateTime(time.Now()),
			Note:      "updated note",
		},
		UnifiedID: contact.ID.String(),
		NoteID:    note.ID.String(),
	}
	updated, err := testCtxt.client.PatchContactNote(params)
	assert.Equal(t, nil, err)
	assert.Equal(t, "updated note", updated.Payload.Note)
}

func TestDeleteContactNoteHandler(t *testing.T) {
	testCtxt := newNoteTestContext(t)
	testCtxt.setupTestUser()
	defer t.Cleanup(func() { testCtxt.ClearDB() })

	userId := domain.UserID(testCtxt.User.UserID)
	contact, err := createTestUnified(testCtxt.unifiedRepo, userId, nil)
	assert.Equal(t, nil, err)

	note, err := createTestContactNote(testCtxt.noteRepo, userId, domain.UnifiedId(contact.ID), nil)
	assert.Equal(t, nil, err)

	params := &operations.DeleteContactNoteParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		UnifiedID:              contact.ID.String(),
		NoteID:                 note.ID.String(),
	}
	_, err = testCtxt.client.DeleteContactNote(params)
	assert.Equal(t, nil, err)
}

func TestGetContactNotesHandler(t *testing.T) {
	testCtxt := newNoteTestContext(t)
	testCtxt.setupTestUser()
	defer t.Cleanup(func() { testCtxt.ClearDB() })

	userId := domain.UserID(testCtxt.User.UserID)
	contact, err := createTestUnified(testCtxt.unifiedRepo, userId, nil)
	assert.Equal(t, nil, err)

	_, err = createTestContactNote(testCtxt.noteRepo, userId, domain.UnifiedId(contact.ID), nil)
	assert.Equal(t, nil, err)

	params := &operations.GetContactNotesParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		UnifiedID:              contact.ID.String(),
	}
	notes, err := testCtxt.client.GetContactNotes(params)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(notes.Payload))
}

func TestGetContactNotesHandler_Pagination(t *testing.T) {
	testCtxt := newNoteTestContext(t)
	testCtxt.setupTestUser()
	defer t.Cleanup(func() { testCtxt.ClearDB() })

	userId := domain.UserID(testCtxt.User.UserID)
	contact, err := createTestUnified(testCtxt.unifiedRepo, userId, nil)
	assert.Equal(t, nil, err)

	ids := []string{}
	for i := 0; i < 10; i++ {
		created, err := createTestContactNote(testCtxt.noteRepo, userId, domain.UnifiedId(contact.ID), func(v *domain.Note) { v.Note = fmt.Sprintf("Note %d", i) })
		assert.Equal(t, err, nil)
		ids = append(ids, created.ID.String())
	}
	sort.StringSlice(ids).Sort()

	var limit int64 = 5
	params := &operations.GetContactNotesParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		UnifiedID:              contact.ID.String(),
		Limit:                  &limit,
	}
	firstBatchResponse, err := testCtxt.client.GetContactNotes(params)
	assert.Equal(t, err, nil)
	assert.Equal(t, int(limit), len(firstBatchResponse.Payload))
	for i := 0; i < len(firstBatchResponse.Payload); i++ {
		assert.Equal(t, ids[i], firstBatchResponse.Payload[i].ID)
	}

	// next batch
	params = &operations.GetContactNotesParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		UnifiedID:              contact.ID.String(),
		Limit:                  &limit,
		LastDocumentID:         &firstBatchResponse.Payload[len(firstBatchResponse.Payload)-1].ID,
	}
	secondBatchResponse, err := testCtxt.client.GetContactNotes(params)
	assert.Equal(t, err, nil)
	assert.Equal(t, int(limit), len(secondBatchResponse.Payload))
	for i := 0; i < len(secondBatchResponse.Payload); i++ {
		assert.Equal(t, ids[i+int(limit)], secondBatchResponse.Payload[i].ID)
	}
}
