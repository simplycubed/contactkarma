package auth

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/firestore"
)

// AuthEvent is the payload of a Firestore Auth event.
type AuthEvent struct {
	Email    string `json:"email"`
	Metadata struct {
		CreatedAt time.Time `json:"createdAt"`
	} `json:"metadata"`
	UID string `json:"uid"`
}

type AuthService struct {
	db *firestore.Client
}

func NewAuthService(db *firestore.Client) *AuthService {
	return &AuthService{db: db}
}

type User struct {
	ID             string         `firestore:"id" json:"id"`
	EmailAddresses []EmailAddress `firestore:"email_addresses" json:"email_addresses"`
	Quota          *Quota         `firestore:"quota" json:"quota"`
}

type EmailAddress struct {
	DisplayName string `firestore:"display_name,omitempty" json:"display_name,omitempty" fake:"{firstname}"`
	Type        string `firestore:"type,omitempty" json:"type,omitempty"`
	Value       string `firestore:"value,omitempty" json:"value,omitempty" fake:"{email}"`
}

type Quota struct {

	// counter to track total number of contacts added
	TotalContacts int64 `firestore:"total_contacts" json:"total_contacts"`

	// counter to track total number of contacts with category assigned
	TotalCategoryAssigned int64 `firestore:"total_category_assigned" json:"total_category_assigned"`

	// total contact sources
	TotalContactSources int64 `firestore:"total_contact_sources" json:"total_contact_sources"`
}

// OnCreateAuthUser
func (service *AuthService) OnCreateAuthUser(ctx context.Context, e AuthEvent) error {
	log.Printf("Function triggered by creation of user: %q %s", e.UID, e.Email)
	log.Printf("Created at: %v", e.Metadata.CreatedAt)
	if e.Email == "" {
		log.Println("email is empty")
		return nil
	}
	// create user in firestore
	usersCollection := "users"

	user := User{
		ID: e.UID,
		EmailAddresses: []EmailAddress{{
			Value: e.Email,
		}},
		Quota: &Quota{
			TotalContacts:         0,
			TotalCategoryAssigned: 0,
			TotalContactSources:   0,
		},
	}

	_, err := service.db.Collection(usersCollection).Doc(e.UID).Set(ctx, user)
	if err != nil {
		log.Printf("Created at: %v", e.Metadata.CreatedAt)
		return err
	}

	return nil
}
