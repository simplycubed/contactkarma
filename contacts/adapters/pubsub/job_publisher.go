package pubsub

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/pubsub"
	"github.com/simplycubed/contactkarma/contacts/gen-jobs/models"
)

type PullContactSourcePublisher struct {
	topic *pubsub.Topic
}

func NewPullContactSourcePublisher(topic *pubsub.Topic) *PullContactSourcePublisher {
	return &PullContactSourcePublisher{topic: topic}
}

func (publisher *PullContactSourcePublisher) Publish(ctx context.Context, request models.PullContactsRequest) (err error) {
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
