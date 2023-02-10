package application

import (
	"context"

	"github.com/simplycubed/contactkarma/contacts/domain"
	"github.com/simplycubed/contactkarma/contacts/domain/repository"
	"github.com/simplycubed/contactkarma/contacts/gen/models"
)

type noteService struct {
	noteRepo repository.INote
}

func NewNoteService(noteRepo repository.INote) *noteService {
	return &noteService{noteRepo: noteRepo}
}

func (a *noteService) SaveNote(ctx context.Context, uID domain.UserID, cID domain.UnifiedId, note *models.Note) (*models.Note, error) {
	noteToCreate := domain.Note{}
	noteToCreate.FromDto(note)
	created, err := a.noteRepo.SaveNote(ctx, uID, cID, noteToCreate)
	if err != nil {
		return nil, err
	}
	return created.MapToDto(), nil
}
func (a *noteService) UpdateNote(ctx context.Context, uID domain.UserID, cID domain.UnifiedId, nID domain.NoteID, note *models.Note) (*models.Note, error) {
	_, err := a.noteRepo.GetNoteByID(ctx, uID, cID, nID)
	if err != nil {
		return nil, err
	}
	noteToUpdate := domain.Note{}
	noteToUpdate.FromDto(note)
	if err := a.noteRepo.UpdateNote(ctx, uID, cID, nID, noteToUpdate); err != nil {
		return nil, err
	}
	updatedNote, err := a.noteRepo.GetNoteByID(ctx, uID, cID, nID)
	if err != nil {
		return nil, err
	}
	return updatedNote.MapToDto(), nil
}
func (a *noteService) GetAllNotes(ctx context.Context, uID domain.UserID, cID domain.UnifiedId) ([]*models.Note, error) {
	notes, err := a.noteRepo.GetAllNotes(ctx, uID, cID)
	if err != nil {
		return nil, err
	}
	return domain.Notes(notes).MapToDto(), nil
}

func (a *noteService) GetNotes(ctx context.Context, uID domain.UserID, cID domain.UnifiedId, limit int, lastDocumentId *domain.NoteID) ([]*models.Note, error) {
	notes, err := a.noteRepo.GetNotes(ctx, uID, cID, limit, lastDocumentId)
	if err != nil {
		return nil, err
	}
	return domain.Notes(notes).MapToDto(), nil
}

func (a *noteService) DeleteNote(ctx context.Context, uID domain.UserID, cID domain.UnifiedId, nID domain.NoteID) error {
	return a.noteRepo.DeleteNote(ctx, uID, cID, nID)
}
