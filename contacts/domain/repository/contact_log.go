package repository

import (
	"context"

	"github.com/simplycubed/contactkarma/contacts/domain"
)

type IContactLog interface {
	Create(ctx context.Context, userId domain.UserID, log domain.ContactLog) (created *domain.ContactLog, err error)
	Save(ctx context.Context, userId domain.UserID, id domain.ContactLogId, update domain.ContactLog) (err error)
	GetById(ctx context.Context, userId domain.UserID, id domain.ContactLogId) (log domain.ContactLog, err error)
	GetByUnifiedId(ctx context.Context, userId domain.UserID, unifiedId domain.UnifiedId) (log []domain.ContactLog, err error)
	Get(ctx context.Context, id domain.UserID, limit int, lastDocumentId *domain.ContactLogId) ([]domain.ContactLog, error)
	Delete(ctx context.Context, userId domain.UserID, id domain.ContactLogId) error
}
