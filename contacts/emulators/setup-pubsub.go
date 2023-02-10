package main

import (
	"context"
	"fmt"
	"log"
	"os"

	googlePubsub "cloud.google.com/go/pubsub"
	"github.com/simplycubed/contactkarma/contacts/adapters/pubsub"
	"github.com/simplycubed/contactkarma/contacts/conf"
	"github.com/simplycubed/contactkarma/contacts/test/testutils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// setup pubsub for running/testing the project locally
func main() {
	testutils.LoadEnvFile()
	pubsubClient := pubsub.New(conf.ProjectID())

	_, err := pubsubClient.CreateTopic(context.Background(), conf.PullContactsSourceTopic())
	if err != nil && status.Code(err) != codes.AlreadyExists {
		log.Fatal("Could not create topic:", err)
		return
	}

	topic := pubsubClient.Topic(conf.PullContactsSourceTopic())
	jobsPort := conf.JobPort()
	jobServerHost := os.Getenv("JOB_SERVER_HOST")
	// create a push subscription to local jobs instance
	_, err = pubsubClient.CreateSubscription(context.Background(), "pull-contacts-subscription", googlePubsub.SubscriptionConfig{
		Topic: topic,
		PushConfig: googlePubsub.PushConfig{
			Endpoint:             fmt.Sprintf("http://%s:%d/pull-contacts-source", jobServerHost, jobsPort),
			Attributes:           map[string]string{},
			AuthenticationMethod: nil,
		},
	})
	if err != nil && status.Code(err) != codes.AlreadyExists {
		log.Fatal("Could not create subscription:", err)
		return
	}
}
