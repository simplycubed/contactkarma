package application

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/simplycubed/contactkarma/contacts/adapters/utils"
	"github.com/simplycubed/contactkarma/contacts/domain"
	"github.com/simplycubed/contactkarma/contacts/domain/repository"
	jobModels "github.com/simplycubed/contactkarma/contacts/gen-jobs/models"
	"github.com/simplycubed/contactkarma/contacts/gen/models"
	"golang.org/x/oauth2"
	googleOAuth2 "google.golang.org/api/oauth2/v1"
	people "google.golang.org/api/people/v1"
)

type IContactSource interface {
	// Update() updates contact documents in contacts subcollection under contact source and updates the remote
	Update(ctx context.Context, userId domain.UserID, sourceId domain.ContactSourceID, updates []domain.ContactSourceUpdate) (err error)

	// Delete() deletes a contact documents in contacts subcollection under contact source and deletes it from remote
	Delete(ctx context.Context, userId domain.UserID, sourceId domain.ContactSourceID, contactIds []domain.ContactID) (err error)

	// puller return a new contacts puller.
	Puller(ctx context.Context, userId domain.UserID, source domain.ContactSource) (puller IContactSourcePuller)

	// reader returns a reader instance to read all documents stored in the database (and not from the remote. ie google, outlook etc) for a source.
	Reader(ctx context.Context, userId domain.UserID, sourceId domain.ContactSourceID) IContactSourceReader

	// Remove() will remove contacts from database (and not from the remote, ie google, outlook etc). This is invoked when a contact source is removed.
	Remove(ctx context.Context, userId domain.UserID, sourceId domain.ContactSourceID, contactIds []domain.ContactID) (err error)
}

var ErrPullCompleted = errors.New("pull completed")

// puller returns contacts from a source. Contacts are converted to domain.Contact.
// Pull should be called continuosly until it returns err == ErrPullCompleted (which is the normal termination) or some other error.
type IContactSourcePuller interface {
	Pull(ctx context.Context) (newContacts []domain.Contact, updatedContacts []domain.Contact, deletedContacts []domain.Contact, err error)
}

var ErrSourceReadCompleted = errors.New("source read completed")

// reader returns contacts stored for a source from the database. Contacts are converted to domain.Contact.
// Read should be called continuosly until it returns err == ErrSourceReadCompleted (which is the normal termination) or some other error.
type IContactSourceReader interface {
	Read(ctx context.Context) (contacts []domain.Contact, err error)
}
type IContactSourceProvider interface {
	Get(source domain.Source) IContactSource
}

type GoogleOAuthService interface {
	GetRedirectUrl(ctx context.Context) (url string, err error)
	GetToken(ctx context.Context, code string) (token *oauth2.Token, err error)
	GetUserInfo(ctx context.Context, accessToken string) (info *googleOAuth2.Tokeninfo, err error)
	GetConfig() (config *oauth2.Config)
}

// TODO: move PeopleService to google contact source package
type PeopleService interface {
	List(pageToken *string) (resp *people.ListConnectionsResponse, err error)
	Update(personId string, person *people.Person) (updated *people.Person, err error)
	Get(personId string) (person *people.Person, err error)
	BatchGet(personIds []string) (response *people.GetPeopleResponse, err error)
	BatchUpdate(updates map[string]people.Person) (response *people.BatchUpdateContactsResponse, err error)
	BatchDelete(personIds []string) (err error)
}

type PeopleServiceFactory interface {
	New(ctx context.Context, accessToken string, refreshToken string, expiry time.Time) (service PeopleService)
}

type ContactSourceService struct {
	googleOAuthService            GoogleOAuthService
	contactSourceRepo             repository.IContactSource
	pullContactSourcePublisher    PullContactPublisher
	unifiedContactService         IUnifiedContactService
	userRepo                      repository.IUser
	contactSourceProvider         IContactSourceProvider
	contactSourceDeletedPublisher ContactSourceDeletedPublisher
	unifiedRepo                   repository.IUnified
}

type PullContactPublisher interface {
	Publish(ctx context.Context, job jobModels.PullContactsRequest) (err error)
}

type ContactSourceDeletedPublisher interface {
	Publish(ctx context.Context, job jobModels.ContactSourceDeleted) (err error)
}

func NewContactSourceService(
	googleOAuthService GoogleOAuthService,
	contactSourceRepo repository.IContactSource,
	pullContactSourcePublisher PullContactPublisher,
	unifiedContactService IUnifiedContactService,
	userRepo repository.IUser,
	contactSourceProvider IContactSourceProvider,
	unifiedRepo repository.IUnified,
	contactSourceDeletedPublisher ContactSourceDeletedPublisher,
) *ContactSourceService {
	return &ContactSourceService{
		googleOAuthService:            googleOAuthService,
		contactSourceRepo:             contactSourceRepo,
		pullContactSourcePublisher:    pullContactSourcePublisher,
		unifiedContactService:         unifiedContactService,
		userRepo:                      userRepo,
		contactSourceProvider:         contactSourceProvider,
		unifiedRepo:                   unifiedRepo,
		contactSourceDeletedPublisher: contactSourceDeletedPublisher,
	}
}

func (s *ContactSourceService) GetGoogleRedirectUrl(ctx context.Context) (url string, err error) {
	return s.googleOAuthService.GetRedirectUrl(ctx)
}

func (s *ContactSourceService) LinkGoogleContactSource(ctx context.Context, userId domain.UserID, role *domain.Role, code string) (err error) {

	// check limit
	user, err := s.userRepo.GetUserByID(ctx, userId)
	if err != nil {
		return
	}

	if user.Quota.ContactSources() >= role.MaxContactSources() {
		err = domain.ErrContactSourcesLimitReached
		return
	}

	token, err := s.googleOAuthService.GetToken(ctx, code)
	if err != nil {
		return
	}
	info, err := s.googleOAuthService.GetUserInfo(ctx, token.AccessToken)
	if err != nil {
		return
	}

	existing, err := s.contactSourceRepo.GetByEmail(ctx, userId, info.Email, domain.Google)
	if err != nil {
		return
	}
	if existing == nil {
		contactSource := domain.ContactSource{
			UserID:       userId,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Source:       domain.Google,
			Email:        info.Email,
			GoogleUserId: info.UserId,
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
			TokenExpiry:  token.Expiry,
		}
		var created *domain.ContactSource
		created, err = s.contactSourceRepo.Create(ctx, userId, contactSource)
		if err != nil {
			return
		}
		// update quota
		user.AddToTotalContactSources(1)
		err = s.userRepo.UpdateUser(ctx, userId, *user)
		if err != nil {
			return
		}
		// Publish sync job to start syncing right away
		err = s.pullContactSourcePublisher.Publish(ctx, jobModels.PullContactsRequest{
			ContactSourceID: string(created.ID),
			UserID:          string(userId),
		})
		if err != nil {
			log.Println("error publishing", err.Error())
		}
	} else {
		// contact source was already added
		err = errors.New("contact source has already been linked")
	}
	return
}

func (s *ContactSourceService) GetAllContactSources(ctx context.Context, userId domain.UserID) (sources []*models.ContactSource, err error) {
	contactSources, err := s.contactSourceRepo.GetAll(ctx, userId)
	if err != nil {
		return
	}
	return domain.ContactSources(contactSources).MapToDto(), nil
}

func (s *ContactSourceService) GetContactSources(ctx context.Context, userId domain.UserID, limit int, lastDocumentId *domain.ContactSourceID) (sources []*models.ContactSource, err error) {
	contactSources, err := s.contactSourceRepo.Get(ctx, userId, limit, lastDocumentId)
	if err != nil {
		return
	}
	return domain.ContactSources(contactSources).MapToDto(), nil
}

func (s *ContactSourceService) DeleteContactSource(ctx context.Context, userId domain.UserID, sourceId domain.ContactSourceID, removeContactsFromUnified bool) (err error) {

	source, err := s.contactSourceRepo.GetById(ctx, userId, sourceId)
	if err != nil {
		return
	}

	err = s.contactSourceRepo.Delete(ctx, userId, sourceId)
	if err != nil {
		log.Printf("DeleteContactSource: %v", err)
		return
	}
	user, err := s.userRepo.GetUserByID(ctx, userId)
	if err != nil {
		return
	}
	// restore quota: TODO: discuss add and delete scenario
	user.AddToTotalContactSources(-1)
	err = s.userRepo.UpdateUser(ctx, userId, *user)
	if err != nil {
		return
	}

	err = s.contactSourceDeletedPublisher.Publish(ctx, jobModels.ContactSourceDeleted{
		ContactSourceID:           sourceId.String(),
		RemoveContactsFromUnified: removeContactsFromUnified,
		Source:                    source.Source.String(),
		UserID:                    userId.String(),
	})
	return
}

// SyncContacts pulls contacts from a contact source and does following as part of the execution:
// a. syncs unified collection
//		Each contact from the source should have a corresponding entry in unified collection.
//		This could be a merged contact as well, in which case 2 or more contacts would have same entry in the unified collection.
//		Reference to source contacts are added in a field called 'origins' in unified contact document.
//		'origins' is boolean map where keys are of the form $sourceId:$contactId
//
// b. generates link suggestions if any.
//		When a new contact is added to unified collection and a link suggestion exists
//		for a matching key-value pair (suggestion.key == key && suggestion.value == value),
//		the new contacts is appended to that suggestion.
//		If no link suggestion exists, it is checked for possible matches
//		with other contacts in the unified collection and link suggestions are generated.

func (s *ContactSourceService) SyncContacts(ctx context.Context, userId domain.UserID, sourceId domain.ContactSourceID) (err error) {
	contactSource, err := s.contactSourceRepo.GetById(ctx, userId, sourceId)
	if err != nil {
		return
	}
	puller := s.contactSourceProvider.Get(contactSource.Source).Puller(ctx, userId, contactSource)

	totalContactAdded := 0
	totalContactDeleted := 0

	defer func() {
		// update quota
		quotaUpdateError := s.updateQuota(ctx, userId, totalContactAdded-totalContactDeleted)
		if err == nil {
			err = quotaUpdateError
		}
	}()

	for {
		var newContacts, updatedContacts, deletedContacts []domain.Contact
		newContacts, updatedContacts, deletedContacts, err = puller.Pull(ctx)
		if err == ErrPullCompleted {
			// pull complete
			return nil
		}
		if err != nil {
			return
		}

		for _, contact := range newContacts {
			//create  unified contact
			_, err = s.unifiedContactService.Add(ctx, userId, contactSource.Source, sourceId, domain.ContactID(contact.ID), contact)
			if err != nil {
				return
			}
			totalContactAdded++
		}

		err = s.propogateUpdates(ctx, userId, contactSource, sourceId, updatedContacts)
		if err != nil {
			return
		}

		err = s.propogateDeletes(ctx, userId, contactSource, sourceId, deletedContacts)
		if err != nil {
			return
		}
		totalContactDeleted += len(deletedContacts)
	}
}

func (s *ContactSourceService) updateQuota(ctx context.Context, userId domain.UserID, change int) (err error) {
	// TODO: use transaction to update quota
	user, err := s.userRepo.GetUserByID(ctx, userId)
	user.AddToTotalContacts(int64(change))
	err = s.userRepo.UpdateUser(ctx, userId, *user)
	if err != nil {
		return
	}
	return
}

// propogateUpdates applies update to unified documents, then to other sources linked to them
func (s *ContactSourceService) propogateUpdates(ctx context.Context, userId domain.UserID, updatedContactSource domain.ContactSource, updatedSourceId domain.ContactSourceID, updatedContacts []domain.Contact) (err error) {
	updateGroups := map[domain.ContactSourceID][]domain.ContactSourceUpdate{}
	for _, contact := range updatedContacts {
		contactOrigin := domain.NewContactOrigin(updatedContactSource.Source, updatedSourceId, contact.ID)
		// apply update to unified document
		unifiedContact, updateError := s.updateUnifiedContact(ctx, userId, contactOrigin, contact)
		if updateError != nil {
			return updateError
		}
		// group updates to other sources
		for origin := range unifiedContact.Origins {
			currentOrigin := domain.ContactOrigin(origin)
			if currentOrigin == contactOrigin {
				// ignore this origin
				continue
			}
			updates, ok := updateGroups[currentOrigin.SourceId()]
			if ok {
				updateGroups[currentOrigin.SourceId()] = append(updates, domain.ContactSourceUpdate{ContactId: currentOrigin.ContactId(), Unified: unifiedContact})
			} else {
				updateGroups[currentOrigin.SourceId()] = []domain.ContactSourceUpdate{{ContactId: currentOrigin.ContactId(), Unified: unifiedContact}}
			}
		}
	}

	for sourceId, updates := range updateGroups {
		contactSource, contactSourceError := s.contactSourceRepo.GetById(ctx, userId, sourceId)
		if contactSourceError != nil {
			return contactSourceError
		}
		// apply updates to source and remotes
		err = s.contactSourceProvider.Get(contactSource.Source).Update(ctx, userId, sourceId, updates)
		if err != nil {
			return
		}
	}
	return
}

func (s *ContactSourceService) propogateDeletes(ctx context.Context, userId domain.UserID, updatedContactSource domain.ContactSource, updatedSourceId domain.ContactSourceID, deletedContacts []domain.Contact) (err error) {
	deleteGroups := map[domain.ContactSourceID][]domain.ContactID{}
	unifiedIdsToDelete := []domain.UnifiedId{}
	for _, contact := range deletedContacts {
		contactOrigin := domain.NewContactOrigin(updatedContactSource.Source, updatedSourceId, contact.ID)
		// apply update to unified document
		unifiedContact, contactErr := s.unifiedRepo.GetContactByOrigin(ctx, userId, contactOrigin)
		if contactErr != nil {
			return contactErr
		}
		unifiedIdsToDelete = append(unifiedIdsToDelete, unifiedContact.ID)
		// group updates to other sources
		for origin := range unifiedContact.Origins {
			currentOrigin := domain.ContactOrigin(origin)
			if currentOrigin == contactOrigin {
				// ignore this origin
				continue
			}
			updates, ok := deleteGroups[currentOrigin.SourceId()]
			if ok {
				deleteGroups[currentOrigin.SourceId()] = append(updates, currentOrigin.ContactId())
			} else {
				deleteGroups[currentOrigin.SourceId()] = []domain.ContactID{currentOrigin.ContactId()}
			}
		}
	}
	for sourceId, updates := range deleteGroups {
		contactSource, contactSourceError := s.contactSourceRepo.GetById(ctx, userId, sourceId)
		if contactSourceError != nil {
			return contactSourceError
		}
		err = s.contactSourceProvider.Get(contactSource.Source).Delete(ctx, userId, sourceId, updates)
		if err != nil {
			return
		}
	}
	err = s.unifiedRepo.BulkDeleteContacts(ctx, userId, unifiedIdsToDelete)
	return
}

func (s *ContactSourceService) updateUnifiedContact(ctx context.Context, userId domain.UserID, origin domain.ContactOrigin, contact domain.Contact) (unifiedContact domain.Unified, err error) {
	unifiedContact, err = s.unifiedRepo.GetContactByOrigin(ctx, userId, origin)
	if err == repository.ErrContactNotFound {
		return
	}
	unifiedContact.SetUpdates(contact)
	err = s.unifiedRepo.UpdateContact(ctx, userId, unifiedContact.ID, &unifiedContact)
	return
}

func (s *ContactSourceService) PullContacts(ctx context.Context) (err error) {
	now := time.Now()
	// get all contacts that due for update
	contactSources, err := s.contactSourceRepo.GetByNextUpdateAt(ctx, now)
	if err != nil {
		return
	}

	if len(contactSources) == 0 {
		log.Println("no contact sources found to update")
		return
	}

	for _, source := range contactSources {

		// publish for update
		err = s.pullContactSourcePublisher.Publish(ctx, jobModels.PullContactsRequest{
			ContactSourceID: string(source.ID),
			UserID:          string(source.UserID),
		})
		if err != nil {
			log.Println("error publishing", err.Error())
			return
		}

		// TODO: update next sync at based for user's plan
		updates := map[string]interface{}{
			"next_sync_at": now.Add(24 * time.Hour), // 24 hours
		}
		err = s.contactSourceRepo.UpdateByMap(ctx, source.UserID, source.ID, updates)
		if err != nil {
			return
		}
	}

	log.Printf("published %d job", len(contactSources))

	return
}

func (s *ContactSourceService) OnDeleteContactSource(ctx context.Context, userId domain.UserID, sourceId domain.ContactSourceID, source domain.Source, removeContactsFromUnified bool) (err error) {
	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()

	var errcList []<-chan error

	contactc, errc := s.getContactsFromSource(ctx, userId, sourceId, source)
	errcList = append(errcList, errc)
	// Read and remove contacts in batches
	errc = s.processContactsForSourceRemoval(ctx, userId, source, sourceId, removeContactsFromUnified, contactc)
	errcList = append(errcList, errc)

	errs := utils.MergeErrors(errcList...)

	// wait for errors if any
	for err := range errs {
		if err != nil {
			return err
		}
	}

	return
}

func (s *ContactSourceService) processContactsForSourceRemoval(ctx context.Context, userId domain.UserID, source domain.Source, sourceId domain.ContactSourceID, removeContactsFromUnified bool, contactC <-chan []domain.Contact) <-chan error {
	contactSource := s.contactSourceProvider.Get(source)
	errC := make(chan error, 1)

	go func() {
		defer close(errC)
		for contacts := range contactC {
			// remove/update unified contacts
			contactIds := domain.Contacts(contacts).MapToIds()
			err := s.updateUnifiedForSourceRemoval(ctx, userId, source, sourceId, removeContactsFromUnified, contactIds)
			if err != nil {
				errC <- err
				return
			}
			// remove from source's contact list
			err = contactSource.Remove(ctx, userId, sourceId, contactIds)
			if err != nil {
				errC <- err
				return
			}
		}
	}()

	return errC
}

func (s *ContactSourceService) getContactsFromSource(ctx context.Context, userId domain.UserID, sourceId domain.ContactSourceID, source domain.Source) (<-chan []domain.Contact, <-chan error) {
	contactSource := s.contactSourceProvider.Get(source)
	reader := contactSource.Reader(ctx, userId, sourceId)
	contactC := make(chan []domain.Contact, 10)
	errC := make(chan error, 1)
	go func() {
		defer close(contactC)
		defer close(errC)
		for {
			contacts, err := reader.Read(ctx)
			if err == ErrSourceReadCompleted {
				// read complete
				return
			}
			if err != nil {
				errC <- err
				return
			}
			contactC <- contacts
		}
	}()

	return contactC, errC
}

// if unified contact's only source is the one being deleted and removeContactsFromUnified flag is on, unified contact is deleted
// otherwise source is removed from the origin
func (s *ContactSourceService) updateUnifiedForSourceRemoval(ctx context.Context, userId domain.UserID, source domain.Source, sourceId domain.ContactSourceID, removeContactsFromUnified bool, contactIds []domain.ContactID) (err error) {
	toDelete := []domain.UnifiedId{}
	updates := map[domain.UnifiedId]*domain.Unified{}
	for _, contactId := range contactIds {
		origin := domain.NewContactOrigin(source, sourceId, contactId)
		unified, err := s.unifiedRepo.GetContactByOrigin(ctx, userId, origin)
		if err != nil {
			return err
		}
		origins := unified.Origins
		delete(origins, origin.String())
		if len(origins) == 0 && removeContactsFromUnified {
			toDelete = append(toDelete, unified.ID)
		} else {
			updates[unified.ID] = &domain.Unified{Origins: origins}
		}
	}
	if len(toDelete) > 0 {
		s.unifiedRepo.BulkDeleteContacts(ctx, userId, toDelete)
	}
	if len(updates) > 0 {
		s.unifiedRepo.BulkUpdateContacts(ctx, userId, updates)
	}

	return
}
