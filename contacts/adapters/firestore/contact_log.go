package firestore

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/simplycubed/contactkarma/contacts/domain"
	"google.golang.org/api/iterator"
)

type ContactLogFirestore struct {
	client *firestore.Client
}

func NewContactLogFirestore(c *firestore.Client) *ContactLogFirestore {
	return &ContactLogFirestore{
		client: c,
	}
}

func (u *ContactLogFirestore) collectionRef(id domain.UserID) *firestore.CollectionRef {
	return u.client.Collection(GetUserCollection()).Doc(id.String()).Collection(GetContactLogCollection())
}

func (u *ContactLogFirestore) Create(ctx context.Context, userId domain.UserID, log domain.ContactLog) (created *domain.ContactLog, err error) {
	ref := u.collectionRef(userId).NewDoc()
	log.ID = domain.ContactLogId(ref.ID)
	_, err = ref.Set(ctx, log)
	if err != nil {
		return
	}
	created = &log
	return
}

func (u *ContactLogFirestore) Save(ctx context.Context, userId domain.UserID, id domain.ContactLogId, log domain.ContactLog) (err error) {
	ref := u.collectionRef(userId).Doc(id.String())
	_, err = ref.Set(ctx, log)
	if err != nil {
		return
	}
	return
}

func (u *ContactLogFirestore) GetById(ctx context.Context, userId domain.UserID, id domain.ContactLogId) (link domain.ContactLog, err error) {
	doc, err := u.collectionRef(userId).Doc(id.String()).Get(ctx)
	if err != nil {
		return
	}

	if err = doc.DataTo(&link); err != nil {
		log.Printf("DataTo source: %s, %v", id, err)
		return
	}
	return
}

func (u *ContactLogFirestore) GetByUnifiedId(ctx context.Context, userId domain.UserID, unifiedId domain.UnifiedId) (logs []domain.ContactLog, err error) {
	query := u.collectionRef(userId).WherePath(firestore.FieldPath{"unified_id"}, "==", unifiedId).OrderBy("id", firestore.Asc)
	iter := query.Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		contactLog := domain.ContactLog{}
		err = doc.DataTo(&contactLog)
		if err != nil {
			log.Printf("DataTo contact: %s user %s, %v", doc.Ref.ID, userId, err)
			return nil, err
		}
		logs = append(logs, contactLog)
	}
	return
}

func (u *ContactLogFirestore) Get(ctx context.Context, id domain.UserID, limit int, lastDocumentId *domain.ContactLogId) (logs []domain.ContactLog, err error) {
	query := u.collectionRef(id).Limit(limit).OrderBy("id", firestore.Asc)
	if lastDocumentId != nil {
		query = query.StartAfter(lastDocumentId)
	}
	iter := query.Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		contactLog := domain.ContactLog{}
		err = doc.DataTo(&contactLog)
		if err != nil {
			log.Printf("DataTo contact: %s user %s, %v", doc.Ref.ID, id, err)
			return nil, err
		}
		logs = append(logs, contactLog)
	}
	return
}

func (u *ContactLogFirestore) Delete(ctx context.Context, userId domain.UserID, id domain.ContactLogId) error {
	_, err := u.collectionRef(userId).Doc(id.String()).Delete(ctx)
	if err != nil {
		log.Printf("Unable to delete log: %s, %v", id, err)
		return err
	}
	return nil
}
