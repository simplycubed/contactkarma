package repository

import (
	"context"

	"github.com/simplycubed/contactkarma/contacts/domain"
)

type INote interface {
	//Get
	GetNoteByID(ctx context.Context, uID domain.UserID, cID domain.UnifiedId, nID domain.NoteID) (*domain.Note, error)
	GetAllNotes(context.Context, domain.UserID, domain.UnifiedId) ([]domain.Note, error)
	GetNotes(ctx context.Context, userId domain.UserID, contactId domain.UnifiedId, limit int, lastDocumentId *domain.NoteID) ([]domain.Note, error)

	//Patch
	UpdateNote(context.Context, domain.UserID, domain.UnifiedId, domain.NoteID, domain.Note) error

	//Save
	SaveNote(context.Context, domain.UserID, domain.UnifiedId, domain.Note) (note *domain.Note, err error)

	//Delete
	DeleteNote(context.Context, domain.UserID, domain.UnifiedId, domain.NoteID) error
	DeleteAllNotes(context.Context, domain.UserID, domain.UnifiedId) (err error)
}
