package application

import (
	"context"
	"time"

	"github.com/simplycubed/contactkarma/contacts/domain"
	"github.com/simplycubed/contactkarma/contacts/domain/repository"
	"github.com/simplycubed/contactkarma/contacts/gen/models"
)

type IUnifiedSyncer interface {
	SyncContactToUnified(ctx context.Context, userId domain.UserID, source domain.Source, sourceId domain.ContactSourceID, contactId domain.ContactID, contact domain.Contact) (unified *domain.Unified, err error)
}

type IUnifiedContactService interface {
	GetContacts(ctx context.Context, id domain.UserID, limit int, lastDocumentId *domain.UnifiedId) ([]*models.Unified, error)
	GetPendingContacts(ctx context.Context, id domain.UserID, limit int, lastDocumentInstant *time.Time, lastDocumentId *domain.UnifiedId) ([]*models.Unified, error)
	GetRecentContacts(ctx context.Context, uid domain.UserID, maxDays *int64, limit int, lastDocumentInstant *time.Time, lastDocumentId *domain.UnifiedId) ([]*models.Unified, error)
}
type UnifiedContactService struct {
	unifiedRepo           repository.IUnified
	linkSuggestionService ILinkSuggestionService
	contactLogRepo        repository.IContactLog
}

func NewUnifiedContactService(unifiedRepo repository.IUnified, linkSuggestionService ILinkSuggestionService, contactLogRepo repository.IContactLog) *UnifiedContactService {
	return &UnifiedContactService{
		unifiedRepo:           unifiedRepo,
		linkSuggestionService: linkSuggestionService,
		contactLogRepo:        contactLogRepo,
	}
}

func (a *UnifiedContactService) GetContacts(ctx context.Context, uID domain.UserID, limit int, lastDocumentId *domain.UnifiedId) ([]*models.Unified, error) {
	contacts, err := a.unifiedRepo.GetContacts(ctx, uID, limit, lastDocumentId)
	if err != nil {
		return nil, err
	}
	return domain.UnifiedContacts(contacts).MapToDto(), nil
}

func (a *UnifiedContactService) GetPendingContacts(ctx context.Context, uID domain.UserID, limit int, lastDocumentInstant *time.Time, lastDocumentId *domain.UnifiedId) (pending []*models.Unified, err error) {
	contacts, err := a.unifiedRepo.GetContactsByNextContact(ctx, uID, time.Now(), limit, lastDocumentInstant, lastDocumentId)
	if err != nil {
		return
	}
	pending = domain.UnifiedContacts(contacts).MapToDto()
	return
}

func (a *UnifiedContactService) GetRecentContacts(ctx context.Context, uID domain.UserID, maxDays *int64, limit int, lastDocumentInstant *time.Time, lastDocumentId *domain.UnifiedId) (pending []*models.Unified, err error) {
	contactedAfter := time.Now().Add(-14 * 24 * time.Hour) // default of 14 days
	if maxDays != nil {
		contactedAfter = time.Now().Add(-1 * time.Duration(*maxDays) * 24 * time.Hour)
	}
	contacts, err := a.unifiedRepo.GetContactsByLastContact(ctx, uID, contactedAfter, limit, lastDocumentInstant, lastDocumentId)
	if err != nil {
		return
	}
	pending = domain.UnifiedContacts(contacts).MapToDto()
	return
}

func (s *UnifiedContactService) SyncContactToUnified(ctx context.Context, userId domain.UserID, source domain.Source, sourceId domain.ContactSourceID, contactId domain.ContactID, contact domain.Contact) (createdUnified *domain.Unified, err error) {
	origin := domain.NewContactOrigin(source, sourceId, contactId)
	_, err = s.unifiedRepo.GetContactByOrigin(ctx, userId, origin)
	if err == repository.ErrContactNotFound {
		err = nil
		// contact has not been synced before
		unified := &domain.Unified{}
		unified.FromContact(contact, origin)
		now := time.Now()
		unified.NextContact = &now
		createdUnified, err = s.unifiedRepo.SaveContact(ctx, userId, *unified)
		if err != nil {
			return
		}

		// create log
		contactLog := domain.NewContactLog(*createdUnified, domain.Create)
		_, err = s.contactLogRepo.Create(ctx, userId, contactLog)
		if err != nil {
			return
		}

		// check existing link suggestion for same key-value pair
		var added bool
		added, err = s.linkSuggestionService.CheckAndAddToLinkSuggestion(ctx, userId, *createdUnified)
		if err != nil || added {
			return
		}
		// check and create new link suggestions
		_, err = s.linkSuggestionService.CheckAndCreateLinkSuggestion(ctx, userId, *createdUnified)
		if err != nil {
			return
		}
		return
	}
	return
}
