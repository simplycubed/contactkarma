package firestore

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/simplycubed/contactkarma/contacts/domain"
	"github.com/simplycubed/contactkarma/contacts/domain/repository"
	"google.golang.org/api/iterator"
)

type LinkSuggestionFirestore struct {
	client *firestore.Client
}

func NewLinkSuggestionFirestore(c *firestore.Client) *LinkSuggestionFirestore {
	return &LinkSuggestionFirestore{
		client: c,
	}
}

func (u *LinkSuggestionFirestore) collectionRef(id domain.UserID) *firestore.CollectionRef {
	return u.client.Collection(GetUserCollection()).Doc(id.String()).Collection(GetLinkSuggestionsCollection())
}

func (u *LinkSuggestionFirestore) Create(ctx context.Context, userId domain.UserID, suggestion domain.LinkSuggestion) (created *domain.LinkSuggestion, err error) {
	ref := u.collectionRef(userId).NewDoc()
	suggestion.ID = domain.LinkSuggestionID(ref.ID)
	_, err = ref.Set(ctx, suggestion)
	if err != nil {
		return
	}
	created = &suggestion
	return
}

func (u *LinkSuggestionFirestore) Save(ctx context.Context, userId domain.UserID, id domain.LinkSuggestionID, suggestion domain.LinkSuggestion) (err error) {
	ref := u.collectionRef(userId).Doc(id.String())
	_, err = ref.Set(ctx, suggestion)
	if err != nil {
		return
	}
	return
}

func (u *LinkSuggestionFirestore) GetById(ctx context.Context, userId domain.UserID, id domain.LinkSuggestionID) (link domain.LinkSuggestion, err error) {
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

func (u *LinkSuggestionFirestore) GetByKeyValue(ctx context.Context, userId domain.UserID, key domain.LinkSuggestionKey, value string) (link domain.LinkSuggestion, err error) {
	docs, err := u.collectionRef(userId).
		WherePath(firestore.FieldPath{"key"}, "==", key.String()).
		WherePath(firestore.FieldPath{"value"}, "==", value).
		Limit(1).
		Documents(ctx).
		GetAll()
	if err != nil {
		return
	}
	if len(docs) == 0 {
		err = repository.ErrLinkSuggestionNotFound
		return
	}
	err = docs[0].DataTo(&link)
	return
}

func (u *LinkSuggestionFirestore) GetAll(ctx context.Context, id domain.UserID) (suggestions []domain.LinkSuggestion, err error) {
	iter := u.collectionRef(id).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		suggestion := domain.LinkSuggestion{}
		err = doc.DataTo(&suggestion)
		if err != nil {
			log.Printf("DataTo contact: %s user %s, %v", doc.Ref.ID, id, err)
			return nil, err
		}
		suggestions = append(suggestions, suggestion)
	}
	return
}

func (u *LinkSuggestionFirestore) Get(ctx context.Context, id domain.UserID, limit int, lastDocumentId *domain.LinkSuggestionID) (suggestions []domain.LinkSuggestion, err error) {
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
		suggestion := domain.LinkSuggestion{}
		err = doc.DataTo(&suggestion)
		if err != nil {
			log.Printf("DataTo contact: %s user %s, %v", doc.Ref.ID, id, err)
			return nil, err
		}
		suggestions = append(suggestions, suggestion)
	}
	return
}

func (u *LinkSuggestionFirestore) Delete(ctx context.Context, userId domain.UserID, id domain.LinkSuggestionID) error {
	_, err := u.collectionRef(userId).Doc(id.String()).Delete(ctx)
	if err != nil {
		log.Printf("Unable to delete suggestion: %s, %v", id, err)
		return err
	}
	return nil
}
