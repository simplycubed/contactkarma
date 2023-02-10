package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type firestoreClient struct {
	db        *firestore.Client
	projectID string
}

// InitClient sets client
func (c *firestoreClient) InitClient(opt ...option.ClientOption) error {
	ctx := context.Background()
	dbClient, err := firestore.NewClient(ctx, c.projectID, opt...)
	if err != nil {
		return err
	}
	c.db = dbClient

	return nil
}

// NewFirestoreClient returns ref to a new firestoreClient object
func NewFirestoreClient(projectID string, opt ...option.ClientOption) (c *firestore.Client, err error) {
	client := &firestoreClient{
		projectID: projectID,
	}
	if err := client.InitClient(opt...); err != nil {
		return nil, err
	}

	return client.db, nil
}
