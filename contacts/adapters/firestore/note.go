package firestore

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/simplycubed/contactkarma/contacts/domain"
	"google.golang.org/api/iterator"
)

type NoteFirestore struct {
	client *firestore.Client
}

func NewNoteFirestore(c *firestore.Client) *NoteFirestore {
	return &NoteFirestore{
		client: c,
	}
}

func (u *NoteFirestore) collectionRef(id domain.UserID, contactId domain.UnifiedId) *firestore.CollectionRef {
	return u.client.Collection(GetUserCollection()).Doc(id.String()).
		Collection(GetUnifiedContactsCollection()).Doc(contactId.String()).
		Collection(GetNoteCollection())
}

func (u NoteFirestore) GetNoteByID(ctx context.Context, uID domain.UserID, cID domain.UnifiedId, nID domain.NoteID) (*domain.Note, error) {
	docSnap, err := u.collectionRef(uID, cID).Doc(nID.String()).Get(ctx)
	if err != nil {
		log.Printf("Unable to get note for contact %s, %v", cID, err)
		return nil, err
	}

	note := domain.Note{}
	if err := docSnap.DataTo(&note); err != nil {
		log.Printf("Datato: %v", err)
		return nil, err
	}

	return &note, nil
}

func (u NoteFirestore) GetAllNotes(ctx context.Context, uID domain.UserID, cID domain.UnifiedId) (notes []domain.Note, err error) {
	iter := u.collectionRef(uID, cID).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		note := domain.Note{}
		if err := doc.DataTo(&note); err != nil {
			log.Printf("DataTo: %v", err)
			return nil, err
		}

		note.ID = domain.NoteID(doc.Ref.ID)
		notes = append(notes, note)
	}

	return notes, nil
}

func (u NoteFirestore) GetNotes(ctx context.Context, uID domain.UserID, cID domain.UnifiedId, limit int, lastDocumentId *domain.NoteID) (notes []domain.Note, err error) {
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

		note := domain.Note{}
		if err := doc.DataTo(&note); err != nil {
			log.Printf("DataTo: %v", err)
			return nil, err
		}

		note.ID = domain.NoteID(doc.Ref.ID)
		notes = append(notes, note)
	}

	return notes, nil
}

func (u NoteFirestore) SaveNote(ctx context.Context, uID domain.UserID, cID domain.UnifiedId, note domain.Note) (created *domain.Note, err error) {
	ref := u.collectionRef(uID, cID).NewDoc()
	note.ID = domain.NoteID(ref.ID)
	note.IsUpdated = false

	_, err = ref.Set(ctx, note)
	if err != nil {
		return
	}
	created = &note
	return
}

func (u NoteFirestore) UpdateNote(ctx context.Context, uID domain.UserID, cID domain.UnifiedId, nID domain.NoteID, note domain.Note) error {
	_, err := u.collectionRef(uID, cID).Doc(nID.String()).
		Get(ctx)
	if err != nil {
		log.Printf("Note not found for contact %s, %v", cID, err)
		return err
	}

	note.IsUpdated = true

	_, err = u.collectionRef(uID, cID).Doc(nID.String()).Set(ctx, note)
	if err != nil {
		log.Printf("Unable to update note for contact %s, %v", cID, err)
		return err
	}

	return nil
}

func (u NoteFirestore) DeleteNote(ctx context.Context, uID domain.UserID, cID domain.UnifiedId, nID domain.NoteID) error {
	_, err := u.collectionRef(uID, cID).Doc(nID.String()).Delete(ctx)
	if err != nil {
		log.Printf("Unable to delete note for contact %s, %v", cID, err)
		return err
	}

	return nil
}

func (u NoteFirestore) DeleteAllNotes(ctx context.Context, uID domain.UserID, cID domain.UnifiedId) error {
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
		log.Printf("Deleted batch of %d notes\n", batchSize)
	}
}
