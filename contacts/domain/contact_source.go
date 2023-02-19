package domain

import (
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/simplycubed/contactkarma/contacts/gen/models"
)

type Source string

func (c Source) String() string {
	return string(c)
}

const Google Source = "google"
const Default Source = "default"

type ContactSourceID string

func (c ContactSourceID) String() string {
	return string(c)
}

type ContactSource struct {
	ID        ContactSourceID `firestore:"id" json:"id"`
	UserID    UserID          `firestore:"user_id" json:"user_id"`
	CreatedAt time.Time       `firestore:"created_at" json:"created_at"`
	UpdatedAt time.Time       `firestore:"updated_at" json:"updated_at"`

	Source       Source    `firestore:"source" json:"source"`
	Email        string    `firestore:"email" json:"email"`
	GoogleUserId string    `firestore:"google_user_id" json:"google_user_id"`
	AccessToken  string    `firestore:"access_token" json:"access_token"`
	RefreshToken string    `firestore:"refresh_token" json:"refresh_token"`
	TokenExpiry  time.Time `firestore:"token_expiry" json:"token_expiry"`

	NextSyncAt time.Time `firestore:"next_sync_at" json:"next_sync_at"`
}

func (doc ContactSource) MapToDto() (dto *models.ContactSource) {
	dto = &models.ContactSource{}
	dto.ID = string(doc.ID)
	dto.UserID = string(doc.UserID)
	dto.CreatedAt = strfmt.DateTime(doc.CreatedAt)
	dto.UpdatedAt = strfmt.DateTime(doc.UpdatedAt)
	dto.Source = string(doc.Source)
	dto.Email = doc.Email
	return
}

type ContactSources []ContactSource

func (docs ContactSources) MapToDto() (dto []*models.ContactSource) {
	for _, doc := range docs {
		dto = append(dto, doc.MapToDto())
	}
	return
}

type ContactSourceUpdate struct {
	ContactId ContactID
	Unified   Unified
}
