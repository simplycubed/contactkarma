package repository

import (
	"context"
	"errors"

	"github.com/simplycubed/contactkarma/contacts/domain"
)

var ErrLinkSuggestionNotFound = errors.New("link suggestion not found")

type ILinkSuggestion interface {
	Create(ctx context.Context, userId domain.UserID, suggestion domain.LinkSuggestion) (created *domain.LinkSuggestion, err error)
	Save(ctx context.Context, userId domain.UserID, id domain.LinkSuggestionID, suggestion domain.LinkSuggestion) (err error)
	GetById(ctx context.Context, userId domain.UserID, id domain.LinkSuggestionID) (link domain.LinkSuggestion, err error)
	GetByKeyValue(ctx context.Context, userId domain.UserID, key domain.LinkSuggestionKey, value string) (link domain.LinkSuggestion, err error)
	GetAll(ctx context.Context, id domain.UserID) (suggestions []domain.LinkSuggestion, err error)
	Get(ctx context.Context, id domain.UserID, limit int, lastDocumentId *domain.LinkSuggestionID) ([]domain.LinkSuggestion, error)
	Delete(ctx context.Context, userId domain.UserID, id domain.LinkSuggestionID) error
}
