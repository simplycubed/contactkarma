package firestore

import (
	"context"
	"errors"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/simplycubed/contactkarma/contacts/domain"
)

type UserFirestore struct {
	client *firestore.Client
}

func NewUserFirestore(c *firestore.Client) *UserFirestore {
	return &UserFirestore{
		client: c,
	}
}

func (u *UserFirestore) collectionRef() *firestore.CollectionRef {
	return u.client.Collection(GetUserCollection())
}

func (u UserFirestore) SaveUser(ctx context.Context, uID domain.UserID, user domain.User) (created *domain.User, err error) {
	ref := u.collectionRef().Doc(uID.String())
	user.ID = uID
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	_, err = ref.Set(ctx, user)
	if err != nil {
		return
	}
	created = &user
	return
}

func (u UserFirestore) getFirestoreUpdates(request domain.User) (updates []firestore.Update) {
	if request.Names != nil {
		updates = append(updates, firestore.Update{
			Path:  "names",
			Value: request.Names,
		})
	}
	if request.Addresses != nil {
		updates = append(updates, firestore.Update{
			Path:  "addresses",
			Value: request.Addresses,
		})
	}
	if request.Birthdays != nil {
		updates = append(updates, firestore.Update{
			Path:  "birthdays",
			Value: request.Birthdays,
		})
	}
	if request.EmailAddresses != nil {
		updates = append(updates, firestore.Update{
			Path:  "email_addresses",
			Value: request.EmailAddresses,
		})
	}
	if request.Genders != nil {
		updates = append(updates, firestore.Update{
			Path:  "genders",
			Value: request.Genders,
		})
	}
	if request.Nicknames != nil {
		updates = append(updates, firestore.Update{
			Path:  "nicknames",
			Value: request.Nicknames,
		})
	}
	if request.Occupations != nil {
		updates = append(updates, firestore.Update{
			Path:  "occupations",
			Value: request.Occupations,
		})
	}
	if request.Organizations != nil {
		updates = append(updates, firestore.Update{
			Path:  "organizations",
			Value: request.Organizations,
		})
	}
	if request.PhoneNumbers != nil {
		updates = append(updates, firestore.Update{
			Path:  "phone_numbers",
			Value: request.PhoneNumbers,
		})
	}
	if request.Photos != nil {
		updates = append(updates, firestore.Update{
			Path:  "photos",
			Value: request.Photos,
		})
	}
	if request.Relations != nil {
		updates = append(updates, firestore.Update{
			Path:  "relations",
			Value: request.Relations,
		})
	}
	if request.Urls != nil {
		updates = append(updates, firestore.Update{
			Path:  "urls",
			Value: request.Urls,
		})
	}
	if request.Quota != nil {
		updates = append(updates, firestore.Update{
			Path:  "quota",
			Value: request.Quota,
		})
	}
	if len(updates) > 0 {
		updates = append(updates, firestore.Update{
			Path:  "updated_at",
			Value: time.Now(),
		})
	}
	return
}

//Patch
func (u UserFirestore) UpdateUser(ctx context.Context, uID domain.UserID, update domain.User) (err error) {
	updates := u.getFirestoreUpdates(update)
	if len(updates) == 0 {
		err = errors.New("no fields found in user update request")
		return
	}
	_, err = u.collectionRef().Doc(uID.String()).Update(ctx, updates)
	if err != nil {
		log.Printf("Unable to update user with userID:%s, %v", uID, err)
		return err
	}
	return nil
}

func (u UserFirestore) GetUserByID(ctx context.Context, uID domain.UserID) (user *domain.User, err error) {
	userDoc, err := u.collectionRef().Doc(uID.String()).Get(ctx)
	if err != nil {
		return nil, err
	}
	user = &domain.User{}
	if err := userDoc.DataTo(user); err != nil {
		log.Printf("DataTo user: %s, %v", uID, err)
		return nil, err
	}
	return user, nil
}

func (u UserFirestore) DeleteUser(ctx context.Context, uID domain.UserID) error {
	_, err := u.collectionRef().Doc(uID.String()).Delete(ctx)
	if err != nil {
		log.Printf("Unable to delete user: %s, %v", uID, err)
		return err
	}
	return nil
}
