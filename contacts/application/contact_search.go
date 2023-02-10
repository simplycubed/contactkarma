package application

import (
	"context"

	"github.com/simplycubed/contactkarma/contacts/domain"
	"github.com/simplycubed/contactkarma/contacts/domain/repository"
	"github.com/simplycubed/contactkarma/contacts/gen/models"
)

type ContactSearch interface {
	Search(userId domain.UserID, searchParams domain.SearchParams) (docs []domain.UnifiedSearch, err error)
}

type ContactSearchService interface {
	SearchContacts(ctx context.Context, userId domain.UserID, searchDto models.SearchContactDto) (unified []*models.Unified, err error)
}

type contactSearchService struct {
	contactSearch ContactSearch
	unifiedRepo   repository.IUnified
}

func NewContactSearchService(contactSearch ContactSearch, unifiedRepo repository.IUnified) *contactSearchService {
	return &contactSearchService{contactSearch: contactSearch, unifiedRepo: unifiedRepo}
}

func (s *contactSearchService) SearchContacts(ctx context.Context, userId domain.UserID, searchDto models.SearchContactDto) (unified []*models.Unified, err error) {
	params := domain.SearchParams{}
	params.FromModel(&searchDto)
	results, err := s.contactSearch.Search(userId, params)
	if err != nil {
		return
	}

	contacts, err := s.unifiedRepo.GetByIDs(ctx, userId, domain.UnifiedSearchResults(results).MapToUnifiedIds())
	if err != nil {
		return
	}
	return domain.UnifiedContacts(contacts).MapToDto(), nil
}
