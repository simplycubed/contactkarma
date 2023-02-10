package firestore

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/simplycubed/contactkarma/contacts/domain"
	"google.golang.org/api/iterator"
)

type GoogleContactsFirestore struct {
	client *firestore.Client
}

func NewGoogleContactsFirestore(c *firestore.Client) *GoogleContactsFirestore {
	return &GoogleContactsFirestore{
		client: c,
	}
}

func (u *GoogleContactsFirestore) collectionRef(id domain.UserID, sourceId domain.ContactSourceID) *firestore.CollectionRef {
	return u.client.Collection(GetUserCollection()).Doc(id.String()).
		Collection(GetContactSourceCollection()).Doc(string(sourceId)).
		Collection(GetGoogleContactsCollection())
}

func (u *GoogleContactsFirestore) Create(ctx context.Context, id domain.UserID, sourceId domain.ContactSourceID, personId string, contact domain.GoogleContact) (created *domain.GoogleContact, err error) {
	ref := u.collectionRef(id, sourceId).Doc(personId)
	contact.ID = domain.ContactID(ref.ID)
	_, err = ref.Set(ctx, contact)
	if err != nil {
		return
	}
	created = &contact
	return
}

func (u *GoogleContactsFirestore) Update(ctx context.Context, id domain.UserID, sourceId domain.ContactSourceID, personId string, contact domain.GoogleContact) (err error) {
	ref := u.collectionRef(id, sourceId).Doc(personId)
	contact.ID = domain.ContactID(ref.ID)
	_, err = ref.Set(ctx, contact)
	if err != nil {
		return
	}
	return
}

func (u *GoogleContactsFirestore) Get(ctx context.Context, id domain.UserID, sourceId domain.ContactSourceID, personId string) (contact domain.GoogleContact, err error) {
	doc, err := u.collectionRef(id, sourceId).Doc(personId).Get(ctx)
	if err != nil {
		return
	}
	err = doc.DataTo(&contact)
	return
}

func (u *GoogleContactsFirestore) GetAll(ctx context.Context, uID domain.UserID, sourceId domain.ContactSourceID) ([]domain.GoogleContact, error) {
	iter := u.collectionRef(uID, sourceId).Documents(ctx)
	contacts := []domain.GoogleContact{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		contact := domain.GoogleContact{}
		err = doc.DataTo(&contact)
		if err != nil {
			log.Printf("DataTo contact: %s user %s, %v", doc.Ref.ID, uID, err)
			return nil, err
		}
		contacts = append(contacts, contact)
	}
	return contacts, nil
}

func (u *GoogleContactsFirestore) List(ctx context.Context, id domain.UserID, sourceId domain.ContactSourceID, limit int, lastDocumentId *string) (contacts []domain.GoogleContact, err error) {
	query := u.collectionRef(id, sourceId).Limit(limit).OrderBy("id", firestore.Asc)
	if lastDocumentId != nil {
		query = query.StartAfter(lastDocumentId)
	}
	iter := query.Documents(ctx)
	contacts = []domain.GoogleContact{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		contact := domain.GoogleContact{}
		err = doc.DataTo(&contact)
		if err != nil {
			log.Printf("DataTo contact: %s user %s, %v", doc.Ref.ID, id, err)
			return nil, err
		}
		contacts = append(contacts, contact)
	}
	return contacts, nil
}

func (store *GoogleContactsFirestore) BulkDeleteContacts(ctx context.Context, uID domain.UserID, sourceId domain.ContactSourceID, contactIds []domain.ContactID) (err error) {
	batch := store.client.Batch()
	currentSize := 0
	batchSize := 250
	for _, contactId := range contactIds {
		ref := store.collectionRef(uID, sourceId).Doc(contactId.String())
		batch.Delete(ref)
		currentSize++
		if currentSize == batchSize {
			_, err = batch.Commit(ctx)
			if err != nil {
				return
			}
			// reset batch
			currentSize = 0
			batch = store.client.Batch()
		}
	}
	if currentSize > 0 {
		_, err = batch.Commit(ctx)
		if err != nil {
			return
		}
	}
	return
}
