package application

import (
	"context"
	"log"

	"github.com/simplycubed/contactkarma/contacts/adapters/firestore"
	"github.com/simplycubed/contactkarma/contacts/domain"
	"github.com/simplycubed/contactkarma/contacts/domain/repository"
	"github.com/simplycubed/contactkarma/contacts/gen/models"
)

type contactService struct {
	userRepo                 repository.IUser
	defaultSourceContactRepo repository.IContact
	unifiedRepo              repository.IUnified
	unifiedContactSyncer     IUnifiedSyncer
	contactSourceProvider    IContactSourceProvider
}

func NewContactService(
	userRepo repository.IUser,
	defaultSourceContactRepo repository.IContact,
	unifiedRepo repository.IUnified,
	unifiedContactSyncer IUnifiedSyncer,
	contactSourceProvider IContactSourceProvider,
) *contactService {
	return &contactService{
		userRepo:                 userRepo,
		defaultSourceContactRepo: defaultSourceContactRepo,
		unifiedRepo:              unifiedRepo,
		unifiedContactSyncer:     unifiedContactSyncer,
		contactSourceProvider:    contactSourceProvider,
	}
}

func (a *contactService) SaveContact(ctx context.Context, uID domain.UserID, contact *models.CreateContactDto) (updated *models.Unified, err error) {
	user, err := a.userRepo.GetUserByID(ctx, uID)
	if err != nil {
		return nil, err
	}

	contactToCreate := domain.Contact{}
	contactToCreate.FromDto(contact)

	created, err := a.defaultSourceContactRepo.SaveContact(ctx, uID, contactToCreate)
	if err != nil {
		return nil, err
	}

	unified, err := a.unifiedContactSyncer.SyncContactToUnified(ctx, uID, domain.Default, firestore.DefaultContactSourceId, domain.ContactID(created.ID), *created)
	if err != nil {
		return
	}

	// add to quota
	user.AddToTotalContacts(1)
	err = a.userRepo.UpdateUser(ctx, uID, *user)
	if err != nil {
		return
	}

	return unified.MapToDto(), nil
}

func (a *contactService) GetContacts(ctx context.Context, uID domain.UserID, limit int, lastDocumentId *domain.ContactID) ([]*models.Contact, error) {
	contacts, err := a.defaultSourceContactRepo.GetContacts(ctx, uID, limit, lastDocumentId)
	if err != nil {
		return nil, err
	}
	return domain.Contacts(contacts).MapToDto(), nil
}

func (a *contactService) GetContactByID(ctx context.Context, uID domain.UserID, cID domain.UnifiedId) (*models.Unified, error) {
	contact, err := a.unifiedRepo.GetContactByID(ctx, uID, cID)
	if err != nil {
		return nil, err
	}
	return contact.MapToDto(), nil
}

func (a *contactService) DeleteContact(ctx context.Context, uID domain.UserID, cID domain.UnifiedId) (err error) {

	user, err := a.userRepo.GetUserByID(ctx, uID)
	if err != nil {
		return
	}

	unified, err := a.unifiedRepo.GetContactByID(ctx, uID, cID)
	if err != nil {
		return
	}

	// TODO: delete all from sources this contact is linked to
	for origin, _ := range unified.Origins {
		sourceId := domain.ContactOrigin(origin).SourceId()
		contactId := domain.ContactOrigin(origin).ContactId()

		if sourceId == firestore.DefaultContactSourceId {
			err = a.defaultSourceContactRepo.DeleteContact(ctx, uID, contactId)
			if err != nil {
				return
			}
		} else {
			// TODO: query contact source and sync based source type eg Google.
		}
	}

	if err := a.unifiedRepo.DeleteContact(ctx, uID, cID); err != nil {
		log.Printf("DeleteContact: %v", err)
		return err
	}

	// update quota
	user.AddToTotalContacts(-1)
	err = a.userRepo.UpdateUser(ctx, uID, *user)
	if err != nil {
		return
	}

	// delete all tags and notes added for this contact
	// TODO: do this operation asynchronously
	// TODO: tags and notes are now part of unified
	// err := a.noteRepo.DeleteAllNotes(ctx, uID, cID)
	// if err != nil {
	// 	log.Printf("DeleteAllNotes: %v", err)
	// 	return err
	// }
	// err = a.tagRepo.DeleteAllTags(ctx, uID, cID)
	// if err != nil {
	// 	log.Printf("DeleteAllTags: %v", err)
	// 	return err
	// }
	return nil
}

func (a *contactService) UpdateContact(ctx context.Context, uID domain.UserID, cID domain.UnifiedId, req *models.UpdateUnifiedDto) (dto *models.Unified, err error) {

	// Apply update to unified
	unified, err := a.unifiedRepo.GetContactByID(ctx, uID, cID)
	if err != nil {
		return
	}

	unified.FromUpdateDto(req)

	err = a.unifiedRepo.UpdateContact(ctx, uID, cID, unified)
	if err != nil {
		return
	}

	updated, err := a.unifiedRepo.GetContactByID(ctx, uID, cID)
	if err != nil {
		return
	}

	// updates are applied to all contact sources this contact is linked to
	for origin, _ := range unified.Origins {
		source := domain.ContactOrigin(origin).Source()
		sourceId := domain.ContactOrigin(origin).SourceId()
		contactId := domain.ContactOrigin(origin).ContactId()

		err = a.contactSourceProvider.Get(source).Update(ctx, uID, sourceId, contactId, *updated)
		if err != nil {
			return
		}

	}

	return updated.MapToDto(), nil
}

func (a *contactService) UpdateCategory(ctx context.Context, userId domain.UserID, role *domain.Role, unifiedId domain.UnifiedId, category domain.ContactCatgeory) (unifiedResponse *models.Unified, err error) {
	user, err := a.userRepo.GetUserByID(ctx, userId)
	if err != nil {
		return
	}

	// Apply update to unified
	unified, err := a.unifiedRepo.GetContactByID(ctx, userId, unifiedId)
	if err != nil {
		return
	}

	// limit check
	isNewCategoryAssigment := unified.Category == ""
	if isNewCategoryAssigment {
		if user.Quota.Contacts() >= role.MaxCategoryAssignable() {
			err = domain.ErrCategoryAssignableLimitReached
			return
		}
	}

	unified.Category = category
	err = a.unifiedRepo.UpdateContact(ctx, userId, unifiedId, unified)
	if err != nil {
		return
	}

	// update quota
	user.AddToTotalCategoryAssigned(1)
	err = a.userRepo.UpdateUser(ctx, userId, *user)
	if err != nil {
		return
	}

	return unified.MapToDto(), nil
}
