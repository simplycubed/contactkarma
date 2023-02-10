package testutils

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"

	googlePubsub "cloud.google.com/go/pubsub"
	"github.com/simplycubed/contactkarma/contacts/adapters/firebase"
	"github.com/simplycubed/contactkarma/contacts/conf"
	"github.com/simplycubed/contactkarma/contacts/domain"
	"github.com/simplycubed/contactkarma/contacts/gen/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Env is an abstraction for running e2e tests
type UserRepository interface {
	SaveUser(context.Context, domain.UserID, *models.User) (*models.User, error)
}

type Env struct {
	transport *TestClientTransport
	User      firebase.UserInfo
	UserRepo  UserRepository
}

func NewEnv(handler http.Handler, userRepo UserRepository) (env *Env) {

	testTransport := NewTestClientTransport(handler)

	// add a default test user
	testUser := firebase.UserInfo{
		Name:   "John Doe",
		UserID: "test-user",
	}

	env = &Env{transport: testTransport, User: testUser, UserRepo: userRepo}

	return
}

func LoadEnvFile() {
	filePath := os.Getenv("TEST_ENV_PATH")
	if filePath != "" {
		err := godotenv.Load(filePath)
		if err != nil {
			log.Println("Ignoring error trying to load env file:", err.Error())
		}
	}
}

func (env *Env) Transport() runtime.ClientTransport {
	return env.transport
}

func (env *Env) ClientAuthInfoWriter() runtime.ClientAuthInfoWriterFunc {
	return func(r runtime.ClientRequest, _ strfmt.Registry) error {
		jsonEncoded, err := json.Marshal(env.User)
		if err != nil {
			return err
		}
		baseEncoded := base64.RawURLEncoding.EncodeToString(jsonEncoded)
		r.SetHeaderParam(firebase.GatewayUserInfoHeader, baseEncoded)
		return nil
	}
}

func (env *Env) GetXApigatewayAPIUserinfo() string {
	jsonEncoded, _ := json.Marshal(env.User)
	baseEncoded := base64.RawURLEncoding.EncodeToString(jsonEncoded)
	return baseEncoded
}

func (env *Env) AddTestUser() (err error) {
	_, err = env.UserRepo.SaveUser(context.Background(), domain.UserID(env.User.UserID), &models.User{
		Addresses:      []*models.Address{},
		Birthdays:      []*models.Birthday{},
		EmailAddresses: []*models.EmailAddress{},
		Genders:        []*models.Gender{},
		ID:             "test-user",
		Names:          []*models.UserNames{},
		Nicknames:      []*models.Nickname{},
		Occupations:    []*models.Occupation{},
		Organizations:  []*models.Organization{},
		PhoneNumbers:   []*models.PhoneNumber{},
		Photos:         []*models.Photo{},
		Relations:      []*models.Relation{},
		Urls:           []*models.URL{},
	})
	return
}

func (env *Env) SetRole(role domain.Role) (err error) {
	env.User.StripeRole = &role
	return
}

func (env *Env) RemoveRole(role domain.Role) (err error) {
	env.User.StripeRole = nil
	return
}

func ClearDB() (err error) {
	host := os.Getenv("FIRESTORE_EMULATOR_HOST")
	req, err := http.NewRequest(http.MethodDelete,
		fmt.Sprintf("http://%s/emulator/v1/projects/%s/databases/(default)/documents", host, conf.ProjectID()),
		nil,
	)
	if err != nil {
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if resp.StatusCode != 200 {
		err = errors.New("failed to clear local db")
	}
	return
}

func (env *Env) ClearDB() (err error) {
	return ClearDB()
}

func (env *Env) CreateTestTopic(t *testing.T, client *googlePubsub.Client, topicName string) {
	_, err := client.CreateTopic(context.Background(), topicName)
	if status.Code(err) == codes.AlreadyExists {
		log.Println("[env] topic already exists")
		return
	}
	if err != nil {
		t.Fatal("Could not create topic:", err)
	}
}

func (env *Env) DeleteTestTopic(t *testing.T, client *googlePubsub.Client, topicName string) {
	err := client.Topic(topicName).Delete(context.Background())
	if status.Code(err) == codes.NotFound {
		log.Println("[env] topic not found to delete")
		return
	}
	if err != nil {
		t.Fatal("Could not create topic:", err)
	}
}
