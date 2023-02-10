package pubsub

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

func New(projectId string, opt ...option.ClientOption) *pubsub.Client {
	service, err := pubsub.NewClient(context.Background(), projectId, opt...)
	if err != nil {
		log.Panicln("Unable to connect pubsub ", err)
	}
	return service
}
