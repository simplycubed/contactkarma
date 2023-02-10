package domain

import (
	"time"

	"github.com/simplycubed/contactkarma/contacts/gen/models"
	"github.com/go-openapi/strfmt"
)

type TagID string

func (c TagID) String() string {
	return string(c)
}

type Tag struct {

	// created at
	CreatedAt time.Time `firestore:"created_at" json:"created_at,omitempty"`

	// id
	ID TagID `firestore:"id" json:"id,omitempty"`

	// tag name
	TagName string `firestore:"tag_name" json:"tag_name,omitempty" fake:"{noun}"`
}

func (tag *Tag) FromDto(dto *models.Tag) {
	tag.CreatedAt = time.Time(dto.CreatedAt)
	tag.ID = TagID(dto.ID)
	tag.TagName = dto.TagName
}

func (tag *Tag) MapToDto() (dto *models.Tag) {
	dto = &models.Tag{
		CreatedAt: strfmt.DateTime(tag.CreatedAt),
		ID:        string(tag.ID),
		TagName:   tag.TagName,
	}
	return
}

type Tags []Tag

func (tags Tags) MapToDto() (dto []*models.Tag) {
	for _, contact := range tags {
		dto = append(dto, contact.MapToDto())
	}
	return
}
