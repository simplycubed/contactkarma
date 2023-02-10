package application

import (
	"context"
	"errors"
	"strings"

	"github.com/simplycubed/contactkarma/contacts/domain"
	"github.com/simplycubed/contactkarma/contacts/domain/repository"
	"github.com/simplycubed/contactkarma/contacts/gen/models"
)

type ILinkSuggestionService interface {

	// check for existing link suggestion and append to matches
	CheckAndAddToLinkSuggestion(ctx context.Context, userId domain.UserID, unified domain.Unified) (added bool, err error)

	// create new link suggestions if any
	CheckAndCreateLinkSuggestion(ctx context.Context, userId domain.UserID, unified domain.Unified) (added bool, err error)

	GetLinkSuggestions(ctx context.Context, uID domain.UserID, limit int, lastDocumentId *domain.LinkSuggestionID) ([]*models.LinkSuggestion, error)

	ApplyLinkSuggestion(ctx context.Context, uID domain.UserID, suggestionId domain.LinkSuggestionID, request []domain.UnifiedId) (err error)
}

type LinkSuggestionService struct {
	unifiedRepo        repository.IUnified
	linkSuggestionRepo repository.ILinkSuggestion
}

func NewLinkSuggestionService(unifiedRepo repository.IUnified, linkSuggestionRepo repository.ILinkSuggestion) *LinkSuggestionService {
	return &LinkSuggestionService{
		unifiedRepo:        unifiedRepo,
		linkSuggestionRepo: linkSuggestionRepo,
	}
}

// AddToExistingSuggestion checks if suggestion already exists for a key-value pair
// key could be email, phone or name
func (s *LinkSuggestionService) AddToExistingSuggestion(ctx context.Context, userId domain.UserID, match domain.LinkMatch, key domain.LinkSuggestionKey, value string) (addedToExisting bool, err error) {
	link, err := s.linkSuggestionRepo.GetByKeyValue(ctx, userId, key, value)
	if err != nil {
		if err == repository.ErrLinkSuggestionNotFound {
			return false, nil
		}
		return
	}
	// update existing suggestion
	link.Matches = append(link.Matches, match)
	err = s.linkSuggestionRepo.Save(ctx, userId, link.ID, link)
	if err != nil {
		return
	}
	addedToExisting = true
	return
}

func (s *LinkSuggestionService) CheckAndAddToLinkSuggestion(ctx context.Context, userId domain.UserID, unified domain.Unified) (added bool, err error) {

	match := domain.LinkMatch{
		UnifiedId:   unified.ID,
		DisplayName: unified.GetDisplayName(),
	}

	// Check if matching link suggestion exists

	if len(unified.EmailAddresses) > 0 {
		email := unified.EmailAddresses[0].Value
		if email != "" {
			added, err = s.AddToExistingSuggestion(ctx, userId, match, domain.KeyEmail, email)
			if err != nil || added {
				return
			}
		}
	}

	if len(unified.PhoneNumbers) > 0 {
		phone := unified.PhoneNumbers[0].Value
		if phone != "" {
			added, err = s.AddToExistingSuggestion(ctx, userId, match, domain.KeyPhone, phone)
			if err != nil || added {
				return
			}
		}
	}

	if len(unified.Names) > 0 {
		givenName := unified.Names[0].GivenName
		if givenName != "" {
			added, err = s.AddToExistingSuggestion(ctx, userId, match, domain.KeyName, givenName)
			if err != nil || added {
				return
			}
		}

		middleName := unified.Names[0].MiddleName
		if middleName != "" {
			added, err = s.AddToExistingSuggestion(ctx, userId, match, domain.KeyName, middleName)
			if err != nil || added {
				return
			}
		}

		familyName := unified.Names[0].FamilyName
		if familyName != "" {
			added, err = s.AddToExistingSuggestion(ctx, userId, match, domain.KeyName, familyName)
			if err != nil || added {
				return
			}
		}
	}
	return
}

// Check for possible duplicates in unified collection and create suggestions if any
func (s *LinkSuggestionService) CheckAndCreateLinkSuggestion(ctx context.Context, userId domain.UserID, unified domain.Unified) (added bool, err error) {
	// lookup contacts by search terms
	if len(unified.SearchTerms) == 0 {
		return
	}
	contacts, err := s.unifiedRepo.GetContactBySearchTerms(ctx, userId, unified.SearchTerms)
	if err != nil {
		return
	}
	matches := []domain.Unified{}
	for _, contact := range contacts {
		if contact.ID == unified.ID {
			// ignore same contact
			continue
		}
		matches = append(matches, contact)
	}
	// group email matches
	if len(unified.EmailAddresses) > 0 {
		email := strings.ToLower(unified.EmailAddresses[0].Value)
		if email != "" {
			emailMatches := domain.UnifiedContacts(matches).FilterByEmail(email)
			if len(emailMatches) > 0 {
				_, err = s.CreateLinkSuggestion(ctx, userId, unified, emailMatches, domain.KeyEmail, email)
				if err != nil {
					return
				}
			}
		}
	}

	if len(unified.PhoneNumbers) > 0 {
		phone := strings.ToLower(unified.PhoneNumbers[0].Value)
		if phone != "" {
			phoneMatches := domain.UnifiedContacts(matches).FilterByPhone(phone)
			if len(phoneMatches) > 0 {
				_, err = s.CreateLinkSuggestion(ctx, userId, unified, phoneMatches, domain.KeyPhone, phone)
				if err != nil {
					return
				}
			}
		}
	}

	if len(unified.Names) > 0 {
		givenName := strings.ToLower(unified.Names[0].GivenName)
		if givenName != "" {
			givenNameMatches := domain.UnifiedContacts(matches).FilterByGivenName(givenName)
			if len(givenNameMatches) > 0 {
				_, err = s.CreateLinkSuggestion(ctx, userId, unified, givenNameMatches, domain.KeyName, givenName)
				if err != nil {
					return
				}
			}
		}

		middleName := strings.ToLower(unified.Names[0].MiddleName)
		if middleName != "" {
			middleNameMatches := domain.UnifiedContacts(matches).FilterByGivenName(givenName)
			if len(middleNameMatches) > 0 {
				_, err = s.CreateLinkSuggestion(ctx, userId, unified, middleNameMatches, domain.KeyName, middleName)
				if err != nil {
					return
				}
			}
		}

		familyName := strings.ToLower(unified.Names[0].FamilyName)
		if familyName != "" {
			familyNameMatches := domain.UnifiedContacts(matches).FilterByGivenName(givenName)
			if len(familyNameMatches) > 0 {
				_, err = s.CreateLinkSuggestion(ctx, userId, unified, familyNameMatches, domain.KeyName, familyName)
				if err != nil {
					return
				}
			}
		}

	}
	return
}

func (s *LinkSuggestionService) CreateLinkSuggestion(ctx context.Context, userId domain.UserID, unified domain.Unified, matches []domain.Unified, key domain.LinkSuggestionKey, value string) (created *domain.LinkSuggestion, err error) {
	linkMatches := []domain.LinkMatch{}

	linkMatches = append(linkMatches, domain.LinkMatch{
		UnifiedId:   unified.ID,
		DisplayName: unified.GetDisplayName(),
	})

	for _, match := range matches {
		linkMatches = append(linkMatches, domain.LinkMatch{
			UnifiedId:   match.ID,
			DisplayName: match.GetDisplayName(),
		})
	}

	linkSuggestion := domain.LinkSuggestion{
		Key:     key,
		Value:   value,
		Matches: linkMatches,
	}
	created, err = s.linkSuggestionRepo.Create(ctx, userId, linkSuggestion)
	return
}

func (a *LinkSuggestionService) GetLinkSuggestions(ctx context.Context, uID domain.UserID, limit int, lastDocumentId *domain.LinkSuggestionID) ([]*models.LinkSuggestion, error) {
	suggestions, err := a.linkSuggestionRepo.Get(ctx, uID, limit, lastDocumentId)
	if err != nil {
		return nil, err
	}
	return domain.LinkSuggestions(suggestions).MapToDto(), nil
}

func (a *LinkSuggestionService) ApplyLinkSuggestion(ctx context.Context, uID domain.UserID, suggestionId domain.LinkSuggestionID, ids []domain.UnifiedId) (err error) {

	if len(ids) <= 1 {
		err = errors.New("2 or more ids are required to apply link suggestion")
		return
	}
	unified, err := a.unifiedRepo.GetByIDs(ctx, uID, ids)
	if err != nil {
		return
	}

	// values are deduped and merged into a single unified document
	destination := unified[0]
	for _, contact := range unified[1:] {
		destination.Merge(contact)
	}
	err = a.unifiedRepo.UpdateContact(ctx, uID, destination.ID, &destination)
	if err != nil {
		return
	}

	// delete merged contacts
	for _, contact := range unified[1:] {
		err = a.unifiedRepo.DeleteContact(ctx, uID, contact.ID)
		if err != nil {
			return
		}
		// TODO: discuss if we should delete this unified id from other link suggestions
	}

	// contacts linked. delete link
	err = a.linkSuggestionRepo.Delete(ctx, uID, suggestionId)
	return
}
