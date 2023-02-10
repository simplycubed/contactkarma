package domain

import (
	"time"

	"github.com/simplycubed/contactkarma/contacts/gen/models"
	"github.com/go-openapi/strfmt"
)

type NoteID string

func (c NoteID) String() string {
	return string(c)
}

type Note struct {

	// created at
	CreatedAt time.Time `firestore:"created_at" json:"created_at,omitempty"`

	// id
	ID NoteID `firestore:"id" json:"id,omitempty"`

	// is updated
	IsUpdated bool `firestore:"is_updated" json:"is_updated,omitempty"`

	// note
	Note string `firestore:"note" json:"note,omitempty" fake:"{sentence:5}"`
}

func (user *Note) FromDto(dto *models.Note) {
	user.CreatedAt = time.Time(dto.CreatedAt)
	user.ID = NoteID(dto.ID)
	user.IsUpdated = dto.IsUpdated
	user.Note = dto.Note
}

func (note *Note) MapToDto() (dto *models.Note) {
	dto = &models.Note{
		CreatedAt: strfmt.DateTime(note.CreatedAt),
		ID:        string(note.ID),
		IsUpdated: note.IsUpdated,
		Note:      note.Note,
	}
	return
}

type Notes []Note

func (notes Notes) MapToDto() (dto []*models.Note) {
	for _, note := range notes {
		dto = append(dto, note.MapToDto())
	}
	return
}
