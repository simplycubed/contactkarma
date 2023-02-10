package firestore

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/simplycubed/contactkarma/contacts/domain"
	"google.golang.org/api/iterator"
)

type TagFirestore struct {
	client *firestore.Client
}

func NewTagFirestore(c *firestore.Client) *TagFirestore {
	return &TagFirestore{
		client: c,
	}
}

func (u *TagFirestore) collectionRef(id domain.UserID, contactId domain.UnifiedId) *firestore.CollectionRef {
	return u.client.Collection(GetUserCollection()).Doc(id.String()).
		Collection(GetUnifiedContactsCollection()).Doc(contactId.String()).
		Collection(GetTagCollection())
}

func (u TagFirestore) GetTagByID(ctx context.Context, uID domain.UserID, cID domain.UnifiedId, tID domain.TagID) (*domain.Tag, error) {
	docSnap, err := u.collectionRef(uID, cID).Doc(tID.String()).
		Get(ctx)
	if err != nil {
		log.Printf("Unable to get tag for contact %s, %v", cID, err)
		return nil, err
	}

	tag := domain.Tag{}
	if err := docSnap.DataTo(&tag); err != nil {
		log.Printf("DataTo: %v", err)
		return nil, err
	}

	return &tag, nil
}

func (u TagFirestore) GetAllTags(ctx context.Context, uID domain.UserID, cID domain.UnifiedId) (tags []domain.Tag, err error) {
	iter := u.collectionRef(uID, cID).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		tag := domain.Tag{}
		if err := doc.DataTo(&tag); err != nil {
			log.Printf("DataTo: %v", err)
			return nil, err
		}

		tag.ID = domain.TagID(doc.Ref.ID)
		tags = append(tags, tag)
	}

	return tags, nil
}

func (u TagFirestore) GetTags(ctx context.Context, uID domain.UserID, cID domain.UnifiedId, limit int, lastDocumentId *domain.TagID) (tags []domain.Tag, err error) {
	query := u.collectionRef(uID, cID).Limit(limit).OrderBy("id", firestore.Asc)
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

		tag := domain.Tag{}
		if err := doc.DataTo(&tag); err != nil {
			log.Printf("DataTo: %v", err)
			return nil, err
		}

		tag.ID = domain.TagID(doc.Ref.ID)
		tags = append(tags, tag)
	}

	return tags, nil
}

func (u TagFirestore) SaveTag(ctx context.Context, uID domain.UserID, cID domain.UnifiedId, tag domain.Tag) (created *domain.Tag, err error) {
	ref := u.collectionRef(uID, cID).NewDoc()
	tag.ID = domain.TagID(ref.ID)
	_, err = ref.Set(ctx, tag)
	if err != nil {
		return
	}
	created = &tag
	return
}

func (u TagFirestore) UpdateTag(ctx context.Context, uID domain.UserID, cID domain.UnifiedId, tID domain.TagID, tag domain.Tag) error {
	_, err := u.collectionRef(uID, cID).Doc(tID.String()).Get(ctx)
	if err != nil {
		log.Printf("Tag not found for contact %s, %v", cID, err)
		return err
	}

	_, err = u.collectionRef(uID, cID).Doc(tID.String()).Set(ctx, tag)
	if err != nil {
		log.Printf("Unable to update tag for contact %s, %v", cID, err)
		return err
	}

	return nil
}

func (u TagFirestore) DeleteTag(ctx context.Context, uID domain.UserID, cID domain.UnifiedId, tID domain.TagID) error {
	_, err := u.collectionRef(uID, cID).Doc(tID.String()).Delete(ctx)
	if err != nil {
		log.Printf("Unable to delete tag for contact %s, %v", cID, err)
		return err
	}
	return nil
}

func (u TagFirestore) DeleteAllTags(ctx context.Context, uID domain.UserID, cID domain.UnifiedId) error {
	ref := u.collectionRef(uID, cID)
	batchSize := 50
	for {
		// Get a batch of documents
		iter := ref.Limit(batchSize).Documents(ctx)
		numDeleted := 0

		batch := u.client.Batch()
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}

			batch.Delete(doc.Ref)
			numDeleted++
		}

		if numDeleted == 0 {
			return nil
		}

		_, err := batch.Commit(ctx)
		if err != nil {
			return err
		}
		log.Printf("Deleted batch of %d tags\n", batchSize)
	}
}
