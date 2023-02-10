package routes

import (
	"context"
	"os"
	"testing"

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

func TestMain(m *testing.M) {
	testutils.LoadEnvFile()
	// run tests
	exitVal := m.Run()
	os.Exit(exitVal)
}

type userTestContext struct {
	*testing.T
	*testutils.Env
	client   operations.ClientService
	userRepo *firestore.UserFirestore
}

func newUserTestContext(t *testing.T) *userTestContext {
	// fake google auth credential
	credetialsOption := option.WithCredentials(&google.Credentials{})
	dbpool, _ := firestore.NewFirestoreClient(conf.ProjectID(), credetialsOption)

	userFirestore := firestore.NewUserFirestore(dbpool)
	userService := application.NewUserService(userFirestore)

	testApi := api.Create()
	Users(testApi, userService)

	server := api.CreateServer(testApi)

	// setup env and client
	env := testutils.NewEnv(server.GetHandler(), userService)
	env.ClearDB() // clear db

	testCtxt := &userTestContext{}
	testCtxt.Env = env
	testCtxt.client = operations.New(env.Transport(), strfmt.Default)
	testCtxt.userRepo = userFirestore
	return testCtxt
}

func (ctx *userTestContext) setupTestUser() {
	err := ctx.AddTestUser()
	if err != nil {
		ctx.Fatal("Could not create user:", err)
	}
}

func TestCreateUserHandler(t *testing.T) {
	testCtxt := newUserTestContext(t)
	defer t.Cleanup(func() { testCtxt.ClearDB() })
	params := &operations.CreateUserParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		Body:                   &models.User{Names: []*models.UserNames{{DisplayName: "John Doe"}}},
	}
	user, err := testCtxt.client.CreateUser(params)
	assert.Equal(t, nil, err)
	assert.Equal(t, "John Doe", user.Payload.Names[0].DisplayName)
	assert.Equal(t, string(testCtxt.User.UserID), user.Payload.ID)

}

func TestUpdateUserHandler(t *testing.T) {
	testCtxt := newUserTestContext(t)
	defer t.Cleanup(func() { testCtxt.ClearDB() })
	_, err := testCtxt.userRepo.SaveUser(context.Background(), domain.UserID(testCtxt.User.UserID), domain.User{
		EmailAddresses: []domain.EmailAddress{{
			Value: "john@gmail.com",
		}},
		ID: "test-user",
		Names: []domain.UserNames{{
			DisplayName: "John Doe",
		}},
		Genders: []domain.Gender{},
	})
	assert.Equal(t, nil, err)

	params := &operations.UpdateUserParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		Body: &models.User{
			Names:   []*models.UserNames{{DisplayName: "Sam Doe"}},
			Genders: []*models.Gender{{Value: "male"}},
		},
	}
	user, err := testCtxt.client.UpdateUser(params)
	assert.Equal(t, nil, err)
	assert.Equal(t, "Sam Doe", user.Payload.Names[0].DisplayName)
	assert.Equal(t, "john@gmail.com", user.Payload.EmailAddresses[0].Value) // email shouldn't be replaced as we didn't pass it.
	assert.Equal(t, "male", user.Payload.Genders[0].Value)

	params = &operations.UpdateUserParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		Body:                   &models.User{},
	}
	_, err = testCtxt.client.UpdateUser(params)
	assert.Equal(t, true, err != nil)
	assert.ErrorContains(t, err, "no fields found")
}

func TestGetUserHandler(t *testing.T) {
	testCtxt := newUserTestContext(t)
	testCtxt.setupTestUser()
	defer t.Cleanup(func() { testCtxt.ClearDB() })
	params := &operations.GetUserParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		UserID:                 string(testCtxt.User.UserID),
	}
	user, err := testCtxt.client.GetUser(params)
	assert.Equal(t, err, nil)
	assert.Equal(t, string(testCtxt.User.UserID), user.Payload.ID)

	// should fail when user id don't match
	params = &operations.GetUserParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		UserID:                 "random-id",
	}
	_, err = testCtxt.client.GetUser(params)
	assert.Equal(t, true, err != nil)
	_, isUnAuthorized := err.(*operations.GetUserUnauthorized)
	assert.Equal(t, true, isUnAuthorized)
}

func TestDeleteUserHandler(t *testing.T) {
	testCtxt := newUserTestContext(t)
	testCtxt.setupTestUser()
	//defer t.Cleanup(func() { testCtxt.ClearDB() })
	params := &operations.DeleteUserParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		UserID:                 string(testCtxt.User.UserID),
	}
	_, err := testCtxt.client.DeleteUser(params)
	assert.Equal(t, err, nil)

	// should fail when user id don't match
	params = &operations.DeleteUserParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		UserID:                 "random-id",
	}
	_, err = testCtxt.client.DeleteUser(params)
	assert.Equal(t, true, err != nil)
	_, isUnAuthorized := err.(*operations.DeleteUserUnauthorized)
	assert.Equal(t, true, isUnAuthorized)

	// TODO: fix
	// it should delete all related users as well
	// testCtxt.setupTestUser()

	// _, err = createTestContact(t, nil)
	// assert.Equal(t, err, nil)

	// params = &operations.DeleteUserParams{
	// 	XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
	// 	UserID:                 string(testCtxt.User.UserID),
	// }
	// _, err = testCtxt.client.DeleteUser(params)
	// assert.Equal(t, err, nil)

	// ctx := context.Background()
	// contacts, err := testCtxt.App.GetContacts(ctx, domain.UserID(testCtxt.User.UserID), 10, nil)
	// assert.Equal(t, err, nil)
	// assert.Equal(t, 0, len(contacts))
}
