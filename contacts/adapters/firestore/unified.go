package firestore

import (
	"context"
	"errors"

	"log"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/simplycubed/contactkarma/contacts/adapters/utils"
	"github.com/simplycubed/contactkarma/contacts/domain"
	"github.com/simplycubed/contactkarma/contacts/domain/repository"
	"google.golang.org/api/iterator"
)

type UnifiedContactFirestore struct {
	client *firestore.Client
}

func NewUnifiedContactFirestore(c *firestore.Client) *UnifiedContactFirestore {
	return &UnifiedContactFirestore{
		client: c,
	}
}

func (u *UnifiedContactFirestore) collectionRef(id domain.UserID) *firestore.CollectionRef {
	return u.client.Collection(GetUserCollection()).Doc(id.String()).Collection(GetUnifiedContactsCollection())
}

func (u UnifiedContactFirestore) SaveContact(ctx context.Context, uID domain.UserID, user domain.Unified) (created *domain.Unified, err error) {
	ref := u.collectionRef(uID).NewDoc()
	user.ID = domain.UnifiedId(ref.ID)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	_, err = ref.Set(ctx, user)
	if err != nil {
		return
	}
	created = &user
	return
}

func (u UnifiedContactFirestore) GetContactByID(ctx context.Context, uID domain.UserID, cID domain.UnifiedId) (user *domain.Unified, err error) {
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

func (u UnifiedContactFirestore) GetByIDs(ctx context.Context, uID domain.UserID, unifiedIds []domain.UnifiedId) (users []domain.Unified, err error) {
	refs := []*firestore.DocumentRef{}
	for _, id := range unifiedIds {
		refs = append(refs, u.collectionRef(uID).Doc(id.String()))
	}
	docs, err := u.client.GetAll(ctx, refs)
	if err != nil {
		return
	}
	for _, doc := range docs {
		user := domain.Unified{}
		if err := doc.DataTo(&user); err != nil {
			log.Printf("DataTo user: %s, %v", uID, err)
			return nil, err
		}
		users = append(users, user)
	}
	return
}

func (u *UnifiedContactFirestore) GetContactByOrigin(ctx context.Context, uID domain.UserID, origin domain.ContactOrigin) (unified domain.Unified, err error) {
	docs, err := u.collectionRef(uID).WherePath(firestore.FieldPath{"origins", origin.String()}, "==", true).Limit(1).Documents(ctx).GetAll()
	if err != nil {
		return
	}
	if len(docs) == 0 {
		err = repository.ErrContactNotFound
		return
	}
	err = docs[0].DataTo(&unified)
	return
}

func (u *UnifiedContactFirestore) GetContactBySearchTerms(ctx context.Context, uID domain.UserID, terms []string) (contacts []domain.Unified, err error) {
	docs, err := u.collectionRef(uID).Where("search_terms", "array-contains-any", terms).Documents(ctx).GetAll()
	if err != nil {
		return
	}
	contacts = []domain.Unified{}
	for _, doc := range docs {
		contact := domain.Unified{}
		err = doc.DataTo(&contact)
		if err != nil {
			return
		}
		contacts = append(contacts, contact)
	}
	return
}

func (u UnifiedContactFirestore) GetAllContacts(ctx context.Context, uID domain.UserID) ([]domain.Unified, error) {
	iter := u.collectionRef(uID).Documents(ctx)
	contacts := []domain.Unified{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		contact := domain.Unified{}
		err = doc.DataTo(&contact)
		if err != nil {
			log.Printf("DataTo contact: %s user %s, %v", doc.Ref.ID, uID, err)
			return nil, err
		}
		contacts = append(contacts, contact)
	}
	return contacts, nil
}

func (u UnifiedContactFirestore) GetContacts(ctx context.Context, uID domain.UserID, limit int, lastDocumentId *domain.UnifiedId) ([]domain.Unified, error) {
	query := u.collectionRef(uID).Limit(limit).OrderBy("id", firestore.Asc)
	if lastDocumentId != nil {
		query = query.StartAfter(lastDocumentId)
	}
	iter := query.Documents(ctx)
	contacts := []domain.Unified{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		contact := domain.Unified{}
		err = doc.DataTo(&contact)
		if err != nil {
			log.Printf("DataTo contact: %s user %s, %v", doc.Ref.ID, uID, err)
			return nil, err
		}
		contacts = append(contacts, contact)
	}
	return contacts, nil
}

func (u UnifiedContactFirestore) DeleteContact(ctx context.Context, uID domain.UserID, contactID domain.UnifiedId) error {
	_, err := u.collectionRef(uID).Doc(contactID.String()).Delete(ctx)
	if err != nil {
		log.Printf("Unable to delete contact: %s, %v", contactID, err)
		return err
	}
	return nil
}

func (u UnifiedContactFirestore) getUpdates(request *domain.Unified) (updates []firestore.Update) {
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
	if request.Origins != nil {
		updates = append(updates, firestore.Update{
			Path:  "origins",
			Value: request.Origins,
		})
	}
	if request.Category != "" {
		updates = append(updates, firestore.Update{
			Path:  "category",
			Value: request.Category,
		})
	}
	if len(request.SearchTerms) > 0 {
		updates = append(updates, firestore.Update{
			Path:  "search_terms",
			Value: request.SearchTerms,
		})
	}
	if request.DisplayName != "" {
		updates = append(updates, firestore.Update{
			Path:  "display_name",
			Value: request.DisplayName,
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

func (u UnifiedContactFirestore) UpdateContact(ctx context.Context, uID domain.UserID, cID domain.UnifiedId, contact *domain.Unified) (err error) {
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

func (u *UnifiedContactFirestore) GetContactsByNextContact(ctx context.Context, uID domain.UserID, before time.Time, limit int, lastDocumentInstant *time.Time, lastDocumentId *domain.UnifiedId) ([]domain.Unified, error) {
	query := u.collectionRef(uID).
		WherePath(firestore.FieldPath{"next_contact"}, "<", before).
		Limit(limit).
		OrderBy("next_contact", firestore.Asc).
		OrderBy("id", firestore.Asc)
	if lastDocumentId != nil && lastDocumentInstant != nil {
		query = query.StartAfter(lastDocumentInstant, lastDocumentId)
	}
	iter := query.Documents(ctx)
	contacts := []domain.Unified{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		contact := domain.Unified{}
		err = doc.DataTo(&contact)
		if err != nil {
			log.Printf("DataTo contact: %s user %s, %v", doc.Ref.ID, uID, err)
			return nil, err
		}
		contacts = append(contacts, contact)
	}
	return contacts, nil
}

func (u *UnifiedContactFirestore) GetContactsByLastContact(ctx context.Context, uID domain.UserID, after time.Time, limit int, lastDocumentInstant *time.Time, lastDocumentId *domain.UnifiedId) ([]domain.Unified, error) {
	query := u.collectionRef(uID).
		WherePath(firestore.FieldPath{"last_contact"}, ">", after).
		Limit(limit).
		OrderBy("last_contact", firestore.Desc).
		OrderBy("id", firestore.Asc)
	if lastDocumentId != nil && lastDocumentInstant != nil {
		query = query.StartAfter(lastDocumentInstant, lastDocumentId)
	}
	iter := query.Documents(ctx)
	contacts := []domain.Unified{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		contact := domain.Unified{}
		err = doc.DataTo(&contact)
		if err != nil {
			log.Printf("DataTo contact: %s user %s, %v", doc.Ref.ID, uID, err)
			return nil, err
		}
		contacts = append(contacts, contact)
	}
	return contacts, nil
}

func (u UnifiedContactFirestore) getAllContacts(ctx context.Context, uID domain.UserID) (<-chan domain.Unified, <-chan error) {
	contacts := make(chan domain.Unified, 10)
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
			contact := domain.Unified{}
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

func (u UnifiedContactFirestore) getAllContactsByIds(ctx context.Context, uID domain.UserID, contactIDs []domain.UnifiedId) (<-chan domain.Unified, <-chan error) {
	contacts := make(chan domain.Unified, 10)
	errC := make(chan error, 1)
	go func() {
		defer close(contacts)
		defer close(errC)
		docs, err := u.GetByIDs(ctx, uID, contactIDs)
		if err != nil {
			errC <- err
			return
		}
		for _, doc := range docs {
			contacts <- doc
		}
	}()

	return contacts, errC
}

// delete pipeline to delete all contacts and relates notes and tags
func (store UnifiedContactFirestore) DeleteAllContacts(ctx context.Context, uID domain.UserID) error {

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

// delete pipeline to delete all contacts and relates notes and tags
func (store UnifiedContactFirestore) BulkDeleteContacts(ctx context.Context, uID domain.UserID, contactIDs []domain.UnifiedId) error {

	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()

	var errcList []<-chan error

	contacts, errc := store.getAllContactsByIds(ctx, domain.UserID(uID), contactIDs)
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

// delete pipeline to delete all contacts and relates notes and tags

func (store *UnifiedContactFirestore) BulkUpdateContacts(ctx context.Context, uID domain.UserID, updates map[domain.UnifiedId]*domain.Unified) (err error) {
	batch := store.client.Batch()
	currentSize := 0
	batchSize := 250
	for contactId, update := range updates {
		ref := store.collectionRef(uID).Doc(contactId.String())
		updates := store.getUpdates(update)
		batch.Update(ref, updates)
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

// get all related refs for a contact
func (store UnifiedContactFirestore) getAllContactRefs(ctx context.Context, uID domain.UserID, contacts <-chan domain.Unified) (<-chan *firestore.DocumentRef, <-chan error) {
	refs := make(chan *firestore.DocumentRef, 100)
	errC := make(chan error, 1)

	go func() {
		defer close(refs)
		defer close(errC)
		for contact := range contacts {

			err := store.getAllTagRefs(ctx, uID, domain.UnifiedId(contact.ID), refs)
			if err != nil {
				errC <- err
				return
			}
			err = store.getAllNoteRefs(ctx, uID, domain.UnifiedId(contact.ID), refs)
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

func (store UnifiedContactFirestore) getAllTagRefs(ctx context.Context, uID domain.UserID, cID domain.UnifiedId, refs chan *firestore.DocumentRef) error {
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

func (store UnifiedContactFirestore) getAllNoteRefs(ctx context.Context, uID domain.UserID, cID domain.UnifiedId, refs chan *firestore.DocumentRef) error {
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

func (store UnifiedContactFirestore) deleteRefs(ctx context.Context, refs <-chan *firestore.DocumentRef) <-chan error {
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
