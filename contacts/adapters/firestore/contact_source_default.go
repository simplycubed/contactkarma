package firestore

import (
	"context"
	"errors"

	"log"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/simplycubed/contactkarma/contacts/adapters/utils"
	"github.com/simplycubed/contactkarma/contacts/domain"
	"google.golang.org/api/iterator"
)

type DefaultContactsFirestore struct {
	client *firestore.Client
}

func NewDefaultContactsFirestore(c *firestore.Client) *DefaultContactsFirestore {
	return &DefaultContactsFirestore{
		client: c,
	}
}

func (u *DefaultContactsFirestore) collectionRef(id domain.UserID) *firestore.CollectionRef {
	return u.client.Collection(GetUserCollection()).Doc(id.String()).
		Collection(GetContactSourceCollection()).Doc(string(DefaultContactSourceId)).
		Collection(GetContactCollection())
}

func (u DefaultContactsFirestore) SaveContact(ctx context.Context, uID domain.UserID, user domain.Contact) (created *domain.Contact, err error) {
	ref := u.collectionRef(uID).NewDoc()
	user.ID = domain.ContactID(ref.ID)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	_, err = ref.Set(ctx, user)
	if err != nil {
		return
	}
	created = &user
	return
}

func (u DefaultContactsFirestore) GetContactByID(ctx context.Context, uID domain.UserID, cID domain.ContactID) (user *domain.Contact, err error) {
	userDoc, err := u.collectionRef(uID).Doc(cID.String()).Get(ctx)
	if err != nil {
		return nil, err
	}

	if err := userDoc.DataTo(&user); err != nil {
		log.Printf("DataTo user: %s, %v", uID, err)
		return nil, err
	}

	return user, nil
}

func (u DefaultContactsFirestore) GetAllContacts(ctx context.Context, uID domain.UserID) ([]domain.Contact, error) {
	iter := u.collectionRef(uID).Documents(ctx)
	contacts := []domain.Contact{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		contact := domain.Contact{}
		err = doc.DataTo(&contact)
		if err != nil {
			log.Printf("DataTo contact: %s user %s, %v", doc.Ref.ID, uID, err)
			return nil, err
		}
		contacts = append(contacts, contact)
	}
	return contacts, nil
}

func (u DefaultContactsFirestore) GetContacts(ctx context.Context, uID domain.UserID, limit int, lastDocumentId *domain.ContactID) ([]domain.Contact, error) {
	query := u.collectionRef(uID).Limit(limit).OrderBy("id", firestore.Asc)
	if lastDocumentId != nil {
		query = query.StartAfter(lastDocumentId)
	}
	iter := query.Documents(ctx)
	contacts := []domain.Contact{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		contact := domain.Contact{}
		err = doc.DataTo(&contact)
		if err != nil {
			log.Printf("DataTo contact: %s user %s, %v", doc.Ref.ID, uID, err)
			return nil, err
		}
		contacts = append(contacts, contact)
	}
	return contacts, nil
}

func (u DefaultContactsFirestore) DeleteContact(ctx context.Context, uID domain.UserID, contactID domain.ContactID) error {
	_, err := u.collectionRef(uID).Doc(contactID.String()).Delete(ctx)
	if err != nil {
		log.Printf("Unable to delete contact: %s, %v", contactID, err)
		return err
	}
	return nil
}

func (store *DefaultContactsFirestore) BulkDeleteContacts(ctx context.Context, uID domain.UserID, contactIds []domain.ContactID) (err error) {
	batch := store.client.Batch()
	currentSize := 0
	batchSize := 250
	for _, contactId := range contactIds {
		ref := store.collectionRef(uID).Doc(contactId.String())
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

func (u DefaultContactsFirestore) getUpdates(request *domain.Contact) (updates []firestore.Update) {
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
	if len(updates) > 0 {
		updates = append(updates, firestore.Update{
			Path:  "updated_at",
			Value: time.Now(),
		})
	}
	return
}

func (u DefaultContactsFirestore) UpdateContact(ctx context.Context, uID domain.UserID, cID domain.ContactID, contact *domain.Contact) (err error) {
	updates := u.getUpdates(contact)
	if len(updates) == 0 {
		err = errors.New("no fields found in user update request")
		return
	}
	_, err = u.collectionRef(uID).Doc(cID.String()).Update(ctx, updates)

	if err != nil {
		log.Printf("Unable to update contact with userID:%s & contactID: %s, %v", cID, cID, err)
		return err
	}
	return nil
}

func (u *DefaultContactsFirestore) GetContactsByNextContact(ctx context.Context, uID domain.UserID, before time.Time, limit int, lastDocumentInstant *time.Time, lastDocumentId *domain.ContactID) ([]domain.Contact, error) {
	query := u.collectionRef(uID).
		WherePath(firestore.FieldPath{"next_contact"}, "<", before).
		Limit(limit).
		OrderBy("next_contact", firestore.Asc).
		OrderBy("id", firestore.Asc)
	if lastDocumentId != nil && lastDocumentInstant != nil {
		query = query.StartAfter(lastDocumentInstant, lastDocumentId)
	}
	iter := query.Documents(ctx)
	contacts := []domain.Contact{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		contact := domain.Contact{}
		err = doc.DataTo(&contact)
		if err != nil {
			log.Printf("DataTo contact: %s user %s, %v", doc.Ref.ID, uID, err)
			return nil, err
		}
		contacts = append(contacts, contact)
	}
	return contacts, nil
}

func (u *DefaultContactsFirestore) GetContactsByLastContact(ctx context.Context, uID domain.UserID, after time.Time, limit int, lastDocumentInstant *time.Time, lastDocumentId *domain.ContactID) ([]domain.Contact, error) {
	query := u.collectionRef(uID).
		WherePath(firestore.FieldPath{"last_contact"}, ">", after).
		Limit(limit).
		OrderBy("last_contact", firestore.Desc).
		OrderBy("id", firestore.Asc)
	if lastDocumentId != nil && lastDocumentInstant != nil {
		query = query.StartAfter(lastDocumentInstant, lastDocumentId)
	}
	iter := query.Documents(ctx)
	contacts := []domain.Contact{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		contact := domain.Contact{}
		err = doc.DataTo(&contact)
		if err != nil {
			log.Printf("DataTo contact: %s user %s, %v", doc.Ref.ID, uID, err)
			return nil, err
		}
		contacts = append(contacts, contact)
	}
	return contacts, nil
}

func (u DefaultContactsFirestore) getAllContacts(ctx context.Context, uID domain.UserID) (<-chan domain.Contact, <-chan error) {
	contacts := make(chan domain.Contact, 10)
	errC := make(chan error, 1)

	go func() {
		defer close(contacts)
		defer close(errC)
		iter := u.collectionRef(uID).Documents(ctx)
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				errC <- err
				return
			}
			contact := domain.Contact{}
			err = doc.DataTo(&contact)
			if err != nil {
				errC <- err
				return
			}
			contacts <- contact
		}
	}()

	return contacts, errC
}

// delete pipeline to delete all contacts and relates notes and tags
func (store DefaultContactsFirestore) DeleteAllContacts(ctx context.Context, uID domain.UserID) error {

	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()

	var errcList []<-chan error

	contacts, errc := store.getAllContacts(ctx, domain.UserID(uID))
	errcList = append(errcList, errc)

	refs, errc := store.getAllContactRefs(ctx, uID, contacts)
	errcList = append(errcList, errc)

	errc = store.deleteRefs(ctx, refs)
	errcList = append(errcList, errc)

	errs := utils.MergeErrors(errcList...)

	// wait for errors if any
	for err := range errs {
		if err != nil {
			return err
		}
	}

	return nil
}

// get all related refs for a contact
func (store DefaultContactsFirestore) getAllContactRefs(ctx context.Context, uID domain.UserID, contacts <-chan domain.Contact) (<-chan *firestore.DocumentRef, <-chan error) {
	refs := make(chan *firestore.DocumentRef, 100)
	errC := make(chan error, 1)

	go func() {
		defer close(refs)
		defer close(errC)
		for contact := range contacts {

			err := store.getAllTagRefs(ctx, uID, domain.ContactID(contact.ID), refs)
			if err != nil {
				errC <- err
				return
			}
			err = store.getAllNoteRefs(ctx, uID, domain.ContactID(contact.ID), refs)
			if err != nil {
				errC <- err
				return
			}

			contactRef := store.collectionRef(uID).Doc(string(contact.ID))
			refs <- contactRef
		}
	}()

	return refs, errC
}

// TODO: move to unified
func (store DefaultContactsFirestore) getAllTagRefs(ctx context.Context, uID domain.UserID, cID domain.ContactID, refs chan *firestore.DocumentRef) error {
	ref := store.collectionRef(uID).Doc(cID.String()).Collection(GetTagCollection())
	batchSize := 100
	for {
		iter := ref.Limit(batchSize).Documents(ctx)
		numRead := 0
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}
			refs <- doc.Ref
			numRead++
		}
		if numRead == 0 {
			return nil
		}
	}
}

// TODO: move to unified
func (store DefaultContactsFirestore) getAllNoteRefs(ctx context.Context, uID domain.UserID, cID domain.ContactID, refs chan *firestore.DocumentRef) error {
	ref := store.collectionRef(uID).Doc(cID.String()).Collection(GetNoteCollection())
	batchSize := 100
	for {
		iter := ref.Limit(batchSize).Documents(ctx)
		numRead := 0
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}
			refs <- doc.Ref
			numRead++
		}
		if numRead == 0 {
			return nil
		}
	}
}

func (store DefaultContactsFirestore) deleteRefs(ctx context.Context, refs <-chan *firestore.DocumentRef) <-chan error {
	batch := store.client.Batch()
	currentSize := 0
	batchSize := 250
	errC := make(chan error, 1)
	go func() {
		defer close(errC)
		for ref := range refs {
			batch.Delete(ref)
			currentSize++
			if currentSize == batchSize {
				_, err := batch.Commit(ctx)
				if err != nil {
					errC <- err
					return
				}
				// reset batch
				currentSize = 0
				batch = store.client.Batch()
			}
		}
		if currentSize > 0 {
			_, err := batch.Commit(ctx)
			if err != nil {
				errC <- err
				return
			}
		}
	}()
	return errC
}
