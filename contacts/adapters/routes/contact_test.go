package routes

import (
	"context"
	"strings"
	"testing"

	"github.com/simplycubed/contactkarma/contacts/adapters/api"
	"github.com/simplycubed/contactkarma/contacts/adapters/contactsource"
	"github.com/simplycubed/contactkarma/contacts/adapters/contactsource/defaultcontactsource"
	"github.com/simplycubed/contactkarma/contacts/adapters/contactsource/googlecontactsource"
	"github.com/simplycubed/contactkarma/contacts/adapters/firestore"
	"github.com/simplycubed/contactkarma/contacts/application"
	"github.com/simplycubed/contactkarma/contacts/conf"
	"github.com/simplycubed/contactkarma/contacts/domain"
	"github.com/simplycubed/contactkarma/contacts/gen/client/operations"
	"github.com/simplycubed/contactkarma/contacts/gen/mocks/mock_application"
	"github.com/simplycubed/contactkarma/contacts/gen/models"
	"github.com/simplycubed/contactkarma/contacts/test/testutils"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

type contactTestContext struct {
	*testing.T
	*testutils.Env
	client             operations.ClientService
	userRepo           *firestore.UserFirestore
	contactService     application.ContactService
	unifiedRepo        *firestore.UnifiedContactFirestore
	defaultContactRepo *firestore.DefaultContactsFirestore
	tagRepo            *firestore.TagFirestore
	noteRepo           *firestore.NoteFirestore
	contactLogRepo     *firestore.ContactLogFirestore
}

func newContactTestContext(t *testing.T, ctrl *gomock.Controller) *contactTestContext {

	//googleOAuthService := mock_application.NewMockGoogleOAuthService(ctrl)

	credetialsOption := option.WithCredentials(&google.Credentials{})
	dbpool, _ := firestore.NewFirestoreClient(conf.ProjectID(), credetialsOption)

	userFirestore := firestore.NewUserFirestore(dbpool)
	userService := application.NewUserService(userFirestore)
	defaultContactFirestore := firestore.NewDefaultContactsFirestore(dbpool)
	unifiedRepo := firestore.NewUnifiedContactFirestore(dbpool)
	linkSuggestionRepo := firestore.NewLinkSuggestionFirestore(dbpool)
	contactLogRepo := firestore.NewContactLogFirestore(dbpool)
	linkSuggestionService := application.NewLinkSuggestionService(unifiedRepo, linkSuggestionRepo)
	unifiedContactService := application.NewUnifiedContactService(unifiedRepo, linkSuggestionService, contactLogRepo)
	googleContactsFirestore := firestore.NewGoogleContactsFirestore(dbpool)
	peopleServiceFactory := mock_application.NewMockPeopleServiceFactory(ctrl)

	contactSourceFirestore := firestore.NewContactSourceFirestore(dbpool)
	defaultContactStore := firestore.NewDefaultContactsFirestore(dbpool)

	googleContactSource := googlecontactsource.NewGoogleContactSource(googleContactsFirestore, contactSourceFirestore, peopleServiceFactory, googleContactsFirestore)
	defaultContactSource := defaultcontactsource.NewDefaultContactSource(defaultContactStore)
	contactSourceProvider := contactsource.NewContactSourceProvider(defaultContactSource, googleContactSource)
	contactService := application.NewContactService(userFirestore, defaultContactFirestore, unifiedRepo, unifiedContactService, contactSourceProvider)
	csvImporter := application.NewCsvImporter(defaultContactFirestore, unifiedContactService)
	tagRepo := firestore.NewTagFirestore(dbpool)
	noteRepo := firestore.NewNoteFirestore(dbpool)
	testApi := api.Create()
	Contacts(testApi, contactService, csvImporter)

	server := api.CreateServer(testApi)

	// setup env and client
	env := testutils.NewEnv(server.GetHandler(), userService)
	env.ClearDB() // clear db

	testCtxt := &contactTestContext{}
	testCtxt.Env = env
	testCtxt.client = operations.New(env.Transport(), strfmt.Default)
	testCtxt.userRepo = userFirestore
	testCtxt.contactService = contactService
	testCtxt.unifiedRepo = unifiedRepo
	testCtxt.defaultContactRepo = defaultContactStore
	testCtxt.noteRepo = noteRepo
	testCtxt.tagRepo = tagRepo
	testCtxt.contactLogRepo = contactLogRepo
	return testCtxt
}

func (ctx *contactTestContext) setupTestUser() {
	err := ctx.AddTestUser()
	if err != nil {
		ctx.Fatal("Could not create user:", err)
	}
}

func TestCreateUserContactHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	testCtxt := newContactTestContext(t, ctrl)
	testCtxt.setupTestUser()
	defer t.Cleanup(func() { testCtxt.ClearDB() })
	params := &operations.CreateUserContactParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		Body: &models.CreateContactDto{
			Names: []*models.UserNames{{DisplayName: "John Doe", GivenName: "Johnny"}},
			EmailAddresses: []*models.EmailAddress{{
				Value: "john@gmail.com",
			}},
			Birthdays: []*models.Birthday{{
				Date: "27-06-2000",
			}},
			Genders: []*models.Gender{{
				Value: "male",
			}},
		},
	}
	contact, err := testCtxt.client.CreateUserContact(params)
	assert.Equal(t, err, nil)
	assert.Equal(t, true, contact.Payload.ID != "")

	contacts, err := testCtxt.contactService.GetContacts(context.Background(), domain.UserID(testCtxt.User.UserID), 10, nil)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(contacts))
	assert.Equal(t, true, contacts[0].ID != "") // id should be defined
	// should add to unified as well
	unified, err := testCtxt.unifiedRepo.GetContactByID(context.Background(), domain.UserID(testCtxt.User.UserID), domain.UnifiedId(contact.Payload.ID))
	assert.Equal(t, nil, err)
	assert.Equal(t, true, unified.ID != "")
	assert.Equal(t, "Johnny", unified.DisplayName)
	// should create a log
	logs, err := testCtxt.contactLogRepo.GetByUnifiedId(context.Background(), domain.UserID(testCtxt.User.UserID), unified.ID)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(logs))
	// should update quota
	user, err := testCtxt.userRepo.GetUserByID(context.Background(), domain.UserID(testCtxt.User.UserID))
	assert.Equal(t, err, nil)
	assert.Equal(t, true, user.Quota != nil)
	assert.Equal(t, int64(1), user.Quota.TotalContacts)
}

// TODO: fix this test using mock people service
// func TestContactUpdatePropogationToOtherSources(t *testing.T) {
// 	setupTestUser(t)
// 	defer t.Cleanup(func() { testCtxt.ClearDB() })
// 	ctx := context.Background()
// 	userId := domain.UserID(testCtxt.User.UserID)
// 	defaultContact, err := createTestContact(t, func(v *models.Contact) {
// 		v.Names = []*models.UserNames{{DisplayName: "Jane"}}
// 		v.EmailAddresses = []*models.EmailAddress{{Value: "jane@gmail.com"}}
// 	})
// 	assert.Equal(t, err, nil)
// 	source, err := createTestContactSource(t, func(v *domain.ContactSource) { v.Source = domain.Default })
// 	assert.Equal(t, err, nil)
// 	googleContact, err := testCtxt.googleContactRepo.Create(ctx, userId, source.ID, "test-person", domain.Contact{
// 		Names:          []*domain.UserNames{{DisplayName: "Jannet"}},
// 		EmailAddresses: []domain.EmailAddress{{Value: "jane@gmail.com"}},
// 	})
// 	assert.Equal(t, err, nil)

// 	// sync google contact to unified
// 	googleUnifiedContact, err := testCtxt.unifiedContactService.SyncContactToUnified(ctx, userId, domain.Google, source.ID, "test-person", *googleContact)
// 	assert.Equal(t, err, nil)
// 	// Should generate a link suggestion
// 	suggestions, err := testCtxt.linkSuggestionRepo.GetAll(ctx, userId)
// 	assert.Equal(t, err, nil)
// 	assert.Equal(t, 1, len(suggestions))

// 	// apply linking
// 	err = testCtxt.linkSuggestionService.ApplyLinkSuggestion(ctx, userId, suggestions[0].ID, []domain.UnifiedId{domain.UnifiedId(defaultContact.ID), domain.UnifiedId(googleUnifiedContact.ID)})
// 	assert.Equal(t, err, nil)

// 	// updating contact should update both sources
// 	params := &operations.UpdateUserContactParams{
// 		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
// 		Body: &models.Contact{
// 			Names: []*models.UserNames{{DisplayName: "Jane Doe"}},
// 			Birthdays: []*models.Birthday{{
// 				Date: "27-06-2020",
// 			}},
// 			Genders: []*models.Gender{{
// 				Value: "female",
// 			}},
// 		},
// 		UnifiedID: defaultContact.ID,
// 	}
// 	_, err = testCtxt.client.UpdateUserContact(params)
// 	assert.Equal(t, err, nil)

// 	// verify updates
// 	user, err := testCtxt.unifiedRepo.GetContactByID(context.Background(), domain.UserID(testCtxt.User.UserID), domain.UnifiedId(defaultContact.ID))
// 	assert.Equal(t, err, nil)
// 	assert.Equal(t, "Jane Doe", user.Names[0].DisplayName)
// 	assert.Equal(t, "27-06-2020", user.Birthdays[0].Date)
// 	assert.Equal(t, "female", user.Genders[0].Value)
// 	assert.Equal(t, "jane@gmail.com", user.EmailAddresses[0].Value) // email shouldn't change

// 	updateGoogleContact, err := testCtxt.googleContactRepo.Get(ctx, userId, source.ID, "test-person")
// 	assert.Equal(t, err, nil)
// 	assert.Equal(t, "Jane Doe", updateGoogleContact.Names[0].DisplayName)
// 	assert.Equal(t, "27-06-2020", updateGoogleContact.Birthdays[0].Date)
// 	assert.Equal(t, "female", updateGoogleContact.Genders[0].Value)
// 	assert.Equal(t, "jane@gmail.com", updateGoogleContact.EmailAddresses[0].Value)

// 	allDefaultContacts, err := testCtxt.defaultContactRepo.GetAllContacts(ctx, userId)
// 	assert.Equal(t, err, nil)
// 	updatedDefaultContact := allDefaultContacts[0]
// 	assert.Equal(t, "Jannet", updatedDefaultContact.Names[0].DisplayName)
// 	assert.Equal(t, "27-06-2020", updatedDefaultContact.Birthdays[0].Date)
// 	assert.Equal(t, "female", updatedDefaultContact.Genders[0].Value)
// 	assert.Equal(t, "jane@gmail.com", updatedDefaultContact.EmailAddresses[0].Value)
// }

func TestUpdateUserContact(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	testCtxt := newContactTestContext(t, ctrl)
	testCtxt.setupTestUser()
	defer t.Cleanup(func() { testCtxt.ClearDB() })

	_, err := createTestContact(testCtxt.defaultContactRepo, domain.UserID(testCtxt.User.UserID), func(v *domain.Contact) { v.EmailAddresses = []domain.EmailAddress{{Value: "jane@gmail.com"}} })
	assert.Equal(t, err, nil)
	unified, err := createTestUnified(testCtxt.unifiedRepo, domain.UserID(testCtxt.User.UserID), func(v *domain.Unified) { v.EmailAddresses = []domain.EmailAddress{{Value: "jane@gmail.com"}} })
	assert.Equal(t, err, nil)

	// update contact
	params := &operations.UpdateUserContactParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		Body: &models.UpdateUnifiedDto{
			Names: []*models.UserNames{{DisplayName: "Jane Doe"}},
			Birthdays: []*models.Birthday{{
				Date: "27-06-2020",
			}},
			Genders: []*models.Gender{{
				Value: "female",
			}},
			DisplayName: "Jan Doe",
		},
		UnifiedID: unified.ID.String(),
	}
	_, err = testCtxt.client.UpdateUserContact(params)
	assert.Equal(t, err, nil)
	// verify updates
	user, err := testCtxt.unifiedRepo.GetContactByID(context.Background(), domain.UserID(testCtxt.User.UserID), domain.UnifiedId(unified.ID))
	assert.Equal(t, err, nil)
	assert.Equal(t, "Jane Doe", user.Names[0].DisplayName, nil)
	assert.Equal(t, "27-06-2020", user.Birthdays[0].Date, nil)
	assert.Equal(t, "female", user.Genders[0].Value, nil)
	assert.Equal(t, "jane@gmail.com", user.EmailAddresses[0].Value, nil) // email shouldn't change
	assert.Equal(t, "Jan Doe", user.DisplayName, nil)
	// not found 404
	params = &operations.UpdateUserContactParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		Body: &models.UpdateUnifiedDto{
			Names: []*models.UserNames{{
				DisplayName: "Non-Existant User",
			}},
		},
		UnifiedID: "non-existant-id",
	}
	_, err = testCtxt.client.UpdateUserContact(params)
	_, isNotFound := err.(*operations.UpdateUserContactNotFound)
	assert.Equal(t, true, isNotFound)
}

func TestUpdateContactCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	testCtxt := newContactTestContext(t, ctrl)
	testCtxt.setupTestUser()
	defer t.Cleanup(func() { testCtxt.ClearDB() })

	unified, err := createTestUnified(testCtxt.unifiedRepo, domain.UserID(testCtxt.User.UserID), func(v *domain.Unified) { v.EmailAddresses = []domain.EmailAddress{{Value: "jane@gmail.com"}} })
	assert.Equal(t, err, nil)

	// // update contact
	params := &operations.UpdateContactCategoryParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		Body: &models.UpdateCategoryDto{
			Category: models.ContactCategory(domain.A),
		},
		UnifiedID: unified.ID.String(),
	}
	_, err = testCtxt.client.UpdateContactCategory(params)
	assert.Equal(t, err, nil)
	//verify updates
	user, err := testCtxt.unifiedRepo.GetContactByID(context.Background(), domain.UserID(testCtxt.User.UserID), domain.UnifiedId(unified.ID))
	assert.Equal(t, err, nil)
	assert.Equal(t, string(domain.A), string(user.Category), nil)

	// not found 404
	params = &operations.UpdateContactCategoryParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		Body: &models.UpdateCategoryDto{
			Category: models.ContactCategory(domain.A),
		},
		UnifiedID: "non-existant-id",
	}
	_, err = testCtxt.client.UpdateContactCategory(params)
	_, isNotFound := err.(*operations.UpdateContactCategoryNotFound)
	assert.Equal(t, true, isNotFound)
}

func TestGetUserContactByIDHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	testCtxt := newContactTestContext(t, ctrl)
	testCtxt.setupTestUser()
	defer t.Cleanup(func() { testCtxt.ClearDB() })

	created, err := createTestUnified(testCtxt.unifiedRepo, domain.UserID(testCtxt.User.UserID), nil)
	assert.Equal(t, err, nil)
	params := &operations.GetUserContactByIDParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		UnifiedID:              created.ID.String(),
	}
	response, err := testCtxt.client.GetUserContactByID(params)
	assert.Equal(t, err, nil)
	assert.Equal(t, "John Doe", response.Payload.Names[0].DisplayName)

	// contact id not found 404
	params = &operations.GetUserContactByIDParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		UnifiedID:              "non-existant-id",
	}
	_, err = testCtxt.client.GetUserContactByID(params)
	_, notFound := err.(*operations.GetUserContactByIDNotFound)
	assert.Equal(t, true, notFound)
}

func TestDeleteUserContactHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	testCtxt := newContactTestContext(t, ctrl)
	testCtxt.setupTestUser()

	testCtxt.userRepo.UpdateUser(context.Background(), domain.UserID(testCtxt.User.UserID), domain.User{Quota: &domain.Quota{
		TotalContacts:         10,
		TotalCategoryAssigned: 8,
	}})
	defer t.Cleanup(func() { testCtxt.ClearDB() })

	created, err := createTestUnified(testCtxt.unifiedRepo, domain.UserID(testCtxt.User.UserID), func(v *domain.Unified) {})
	assert.Equal(t, err, nil)

	// delete contact
	params := &operations.DeleteUserContactParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		UnifiedID:              created.ID.String(),
	}
	_, err = testCtxt.client.DeleteUserContact(params)
	assert.Equal(t, err, nil)

	// should update quota
	user, err := testCtxt.userRepo.GetUserByID(context.Background(), domain.UserID(testCtxt.User.UserID))
	assert.Equal(t, err, nil)
	assert.Equal(t, true, user.Quota != nil)
	assert.Equal(t, int64(9), user.Quota.TotalContacts)
	assert.Equal(t, int64(8), user.Quota.TotalCategoryAssigned)
	// delete contact should delete tags and notes
	userId := domain.UserID(testCtxt.User.UserID)
	created, err = createTestUnified(testCtxt.unifiedRepo, userId, nil)
	assert.Equal(t, err, nil)
	_, err = createTestContactTag(testCtxt.tagRepo, userId, domain.UnifiedId(created.ID), nil)
	assert.Equal(t, err, nil)
	_, err = createTestContactNote(testCtxt.noteRepo, userId, domain.UnifiedId(created.ID), nil)
	assert.Equal(t, err, nil)

	params = &operations.DeleteUserContactParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		UnifiedID:              created.ID.String(),
	}
	_, err = testCtxt.client.DeleteUserContact(params)
	assert.Equal(t, err, nil)

}

func TestUploadContactCsv(t *testing.T) {
	content := `display_name,display_name_last_first,family_name,given_name,honorific_prefix,honorific_suffix,middle_name,phonetic_family_name,phonetic_full_name,phonetic_given_name,phonetic_honorific_prefix,phonetic_honorific_suffix,phonetic_middle_name,unstructured_name,nickname,city,country,country_code,extended_address,po_box,postal_code,region,street_address,address_type,birth_date,birth_text,email_display_name,email_type,email,address_me_as,gender,occupation,phone_type,phone,photo_default,photo_url,relation,relation_type,url_type,url,department,domain,end_date,job_description,location,name,phonetic_name,start_date,symbol,title,organization_type,is_current`
	content += "\nJohn,Doe,Doe,John,Mr,,Middle,Doe,John Doe,,,,,,jonny,Dubai,address_country,address_country_code,address_extended_address,address_po_box,address_postal_code,address_region,address_street_address,address_type,birth_date,birth_text,email_display_name,email_type,email_value,gender_address_me_as,gender_value,occupation_value,phone_type,phone_value,photo_default,photo_url,relation_person,relation_type,url_type,url_value,organization_department,organization_domain,organization_end_date,organization_job_description,organization_location,organization_name,organization_phonetic_name,organization_start_date,organization_symbol,organization_title,organization_type,organization_is_current"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	testCtxt := newContactTestContext(t, ctrl)
	testCtxt.setupTestUser()
	defer t.Cleanup(func() { testCtxt.ClearDB() })

	params := &operations.UploadContactsCsvParams{
		XApigatewayAPIUserinfo: testCtxt.GetXApigatewayAPIUserinfo(),
		File:                   runtime.NamedReader("file", strings.NewReader(content)),
	}
	_, err := testCtxt.client.UploadContactsCsv(params)
	assert.Equal(t, err, nil)

	ctx := context.Background()
	contacts, err := testCtxt.defaultContactRepo.GetAllContacts(ctx, domain.UserID(testCtxt.User.UserID))
	assert.Equal(t, err, nil)
	assert.Equal(t, 1, len(contacts))
	contact := contacts[0]
	assert.Equal(t, "John", contact.Names[0].DisplayName)
	assert.Equal(t, "Doe", contact.Names[0].DisplayNameLastFirst)
	assert.Equal(t, "John", contact.Names[0].GivenName)
	assert.Equal(t, "Doe", contact.Names[0].FamilyName)
	assert.Equal(t, "Dubai", contact.Addresses[0].City)
}
