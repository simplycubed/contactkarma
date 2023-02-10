/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"

	"github.com/simplycubed/contactkarma/contacts/adapters"
	"github.com/simplycubed/contactkarma/contacts/adapters/contactsource"
	"github.com/simplycubed/contactkarma/contacts/adapters/contactsource/defaultcontactsource"
	"github.com/simplycubed/contactkarma/contacts/adapters/contactsource/googlecontactsource"
	"github.com/simplycubed/contactkarma/contacts/adapters/firestore"
	"github.com/simplycubed/contactkarma/contacts/adapters/googlepeople/servicefactory"
	"github.com/simplycubed/contactkarma/contacts/adapters/jobs"
	"github.com/simplycubed/contactkarma/contacts/adapters/pubsub"
	"github.com/simplycubed/contactkarma/contacts/application"
	"github.com/simplycubed/contactkarma/contacts/conf"
	"github.com/simplycubed/contactkarma/contacts/logger"
	"github.com/spf13/cobra"
)

// jobsCmd represents the jobs command
var jobsCmd = &cobra.Command{
	Use:   "jobs",
	Short: "Starts Contact Karma Jobs Server",
	Long:  `Starts Contact Karma Jobs Server`,
	Run:   StartJobServer,
}

func init() {
	rootCmd.AddCommand(jobsCmd)
}

func StartJobServer(cmd *cobra.Command, args []string) {
	l := logger.NewLogger(conf.Env())

	dbpool := adapters.NewFirestore()

	googleOAuthService := adapters.NewGoogleOAuthService()
	contactSourceFirestore := firestore.NewContactSourceFirestore(dbpool)
	googleContactsFirestore := firestore.NewGoogleContactsFirestore(dbpool)
	peopleServiceFactory := servicefactory.NewPeopleServiceFactory(googleOAuthService)
	pubsubClient := pubsub.New(conf.ProjectID())
	pullContactsTopic := pubsubClient.Topic(conf.PullContactsSourceTopic())
	pullContactSourcePublisher := pubsub.NewPullContactSourcePublisher(pullContactsTopic)

	contactSourceDeletedTopic := pubsubClient.Topic(conf.ContactSourceDeletedTopic())
	contactSourceDeletedPublisher := pubsub.NewContactSourceDeletedPublisher(contactSourceDeletedTopic)

	unifiedRepo := firestore.NewUnifiedContactFirestore(dbpool)
	contactLogRepo := firestore.NewContactLogFirestore(dbpool)

	linkSuggestionRepo := firestore.NewLinkSuggestionFirestore(dbpool)
	linkSuggestionService := application.NewLinkSuggestionService(unifiedRepo, linkSuggestionRepo)
	unifiedContactService := application.NewUnifiedContactService(unifiedRepo, linkSuggestionService, contactLogRepo)
	userStore := firestore.NewUserFirestore(dbpool)
	defaultContactFirestore := firestore.NewDefaultContactsFirestore(dbpool)

	defaultContactSource := defaultcontactsource.NewDefaultContactSource(defaultContactFirestore)
	googleContactSource := googlecontactsource.NewGoogleContactSource(googleContactsFirestore, contactSourceFirestore, peopleServiceFactory, googleContactsFirestore)
	contactSourceProvider := contactsource.NewContactSourceProvider(defaultContactSource, googleContactSource)

	contactSourceService := application.NewContactSourceService(googleOAuthService, contactSourceFirestore, pullContactSourcePublisher, unifiedContactService, userStore, contactSourceProvider, unifiedRepo, contactSourceDeletedPublisher)

	api := jobs.CreateApi()
	// configure routes
	jobs.Routes(api, contactSourceService)

	srv := jobs.CreateServer(api)

	defer func() {
		l.Info("shutting down...")
		srv.Shutdown()
	}()

	l.Info("starting job server...")

	if err := srv.Serve(); err != nil {
		log.Panicln("server err:", err)
	}
}
