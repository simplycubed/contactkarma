package firestore

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/simplycubed/contactkarma/contacts/domain"
	"google.golang.org/api/iterator"
)

const DefaultContactSourceId domain.ContactSourceID = "default"

type ContactSourceFirestore struct {
	client *firestore.Client
}

func NewContactSourceFirestore(c *firestore.Client) *ContactSourceFirestore {
	return &ContactSourceFirestore{
		client: c,
	}
}

func (u *ContactSourceFirestore) rootRef(id domain.UserID) *firestore.DocumentRef {
	return u.client.Collection(GetUserCollection()).Doc(id.String())
}

func (u *ContactSourceFirestore) collectionRef(id domain.UserID) *firestore.CollectionRef {
	return u.client.Collection(GetUserCollection()).Doc(id.String()).Collection(GetContactSourceCollection())
}

func (u *ContactSourceFirestore) GetByEmail(ctx context.Context, id domain.UserID, email string, source domain.Source) (cs *domain.ContactSource, err error) {
	iter := u.rootRef(id).Collection(GetContactSourceCollection()).
		Where("email", "==", email).
		Where("source", "==", source.String()).
		Limit(1).
		Documents(ctx)
	if err != nil {
		return
	}
	var doc *firestore.DocumentSnapshot
	for {
		doc, err = iter.Next()
		if err == iterator.Done {
			err = nil
			break
		}
		if err != nil {
			return
		}
		csDoc := domain.ContactSource{}
		err = doc.DataTo(&csDoc)
		if err != nil {
			return
		}
		cs = &csDoc
		iter.Stop()
	}
	return
}

func (u *ContactSourceFirestore) Save(ctx context.Context, id domain.UserID, sourceId domain.ContactSourceID, source domain.ContactSource) (created *domain.ContactSource, err error) {
	ref := u.rootRef(id).Collection(GetContactSourceCollection()).Doc(sourceId.String())
	source.ID = domain.ContactSourceID(ref.ID)
	_, err = ref.Set(ctx, source)
	if err != nil {
		return
	}
	created = &source
	return
}

func (u *ContactSourceFirestore) Create(ctx context.Context, id domain.UserID, source domain.ContactSource) (created *domain.ContactSource, err error) {
	ref := u.rootRef(id).Collection(GetContactSourceCollection()).NewDoc()
	source.ID = domain.ContactSourceID(ref.ID)
	_, err = ref.Set(ctx, source)
	if err != nil {
		return
	}
	created = &source
	return
}

func (u *ContactSourceFirestore) GetAll(ctx context.Context, id domain.UserID) (sources []domain.ContactSource, err error) {
	iter := u.collectionRef(id).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		source := domain.ContactSource{}
		err = doc.DataTo(&source)
		if err != nil {
			log.Printf("DataTo contact: %s user %s, %v", doc.Ref.ID, id, err)
			return nil, err
		}
		sources = append(sources, source)
	}
	return
}

func (u *ContactSourceFirestore) Get(ctx context.Context, id domain.UserID, limit int, lastDocumentId *domain.ContactSourceID) (sources []domain.ContactSource, err error) {
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
		source := domain.ContactSource{}
		err = doc.DataTo(&source)
		if err != nil {
			log.Printf("DataTo contact: %s user %s, %v", doc.Ref.ID, id, err)
			return nil, err
		}
		sources = append(sources, source)
	}
	return
}

func (u ContactSourceFirestore) UpdateByMap(ctx context.Context, uID domain.UserID, sourceId domain.ContactSourceID, updates map[string]interface{}) (err error) {
	_, err = u.collectionRef(uID).Doc(sourceId.String()).Set(ctx, updates, firestore.MergeAll)
	return
}

func (u ContactSourceFirestore) GetByNextUpdateAt(ctx context.Context, before time.Time) (sources []domain.ContactSource, err error) {
	query := u.client.CollectionGroup(GetContactSourceCollection()).
		WherePath(firestore.FieldPath{"next_sync_at"}, "<", before).
		OrderBy("next_sync_at", firestore.Asc)
	iter := query.Documents(ctx)
	sources = []domain.ContactSource{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		source := domain.ContactSource{}
		err = doc.DataTo(&source)
		if err != nil {
			return nil, err
		}
		sources = append(sources, source)
	}
	return
}

func (u *ContactSourceFirestore) Delete(ctx context.Context, id domain.UserID, sourceId domain.ContactSourceID) error {
	_, err := u.collectionRef(id).Doc(sourceId.String()).Delete(ctx)
	if err != nil {
		log.Printf("Unable to delete contact source: %s, %v", sourceId, err)
		return err
	}
	return nil
}

func (u *ContactSourceFirestore) GetById(ctx context.Context, id domain.UserID, sourceId domain.ContactSourceID) (source domain.ContactSource, err error) {
	doc, err := u.collectionRef(id).Doc(sourceId.String()).Get(ctx)
	if err != nil {
		return
	}

	if err = doc.DataTo(&source); err != nil {
		log.Printf("DataTo source: %s, %v", id, err)
		return
	}
	return
}
