package pubsub

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/pubsub"
	"github.com/simplycubed/contactkarma/contacts/gen-jobs/models"
)

type ContactSourceDeletedPublisher struct {
	topic *pubsub.Topic
}

func NewContactSourceDeletedPublisher(topic *pubsub.Topic) *ContactSourceDeletedPublisher {
	return &ContactSourceDeletedPublisher{topic: topic}
}

func (publisher *ContactSourceDeletedPublisher) Publish(ctx context.Context, request models.ContactSourceDeleted) (err error) {
	msg, err := json.Marshal(request)
	if err != nil {
		return
	}
	result := publisher.topic.Publish(ctx, &pubsub.Message{
		Data: msg,
	})
	_, err = result.Get(ctx) // block till publish
	return
}
