package cmd

import (
	"log"

	"github.com/simplycubed/contactkarma/contacts/adapters"
	"github.com/simplycubed/contactkarma/contacts/adapters/api"
	"github.com/simplycubed/contactkarma/contacts/adapters/contactsource"
	"github.com/simplycubed/contactkarma/contacts/adapters/contactsource/defaultcontactsource"
	"github.com/simplycubed/contactkarma/contacts/adapters/contactsource/googlecontactsource"
	"github.com/simplycubed/contactkarma/contacts/adapters/firestore"
	"github.com/simplycubed/contactkarma/contacts/adapters/googlepeople/servicefactory"
	"github.com/simplycubed/contactkarma/contacts/adapters/pubsub"
	"github.com/simplycubed/contactkarma/contacts/adapters/routes"
	"github.com/simplycubed/contactkarma/contacts/adapters/typesense"
	"github.com/simplycubed/contactkarma/contacts/application"
	"github.com/simplycubed/contactkarma/contacts/conf"
	"github.com/simplycubed/contactkarma/contacts/logger"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts Contact Karma HTTP Server",
	Run:   StartServer,
}

func StartServer(cmd *cobra.Command, args []string) {
	l := logger.NewLogger(conf.Env())

	dbpool := adapters.NewFirestore()

	userFirestore := firestore.NewUserFirestore(dbpool)
	defaultContactFirestore := firestore.NewDefaultContactsFirestore(dbpool)
	noteFirestore := firestore.NewNoteFirestore(dbpool)
	tagFirestore := firestore.NewTagFirestore(dbpool)
	unifiedContactFirestore := firestore.NewUnifiedContactFirestore(dbpool)

	googleOAuthService := adapters.NewGoogleOAuthService()
	contactSourceFirestore := firestore.NewContactSourceFirestore(dbpool)
	pubsubClient := pubsub.New(conf.ProjectID())
	pullContactsTopic := pubsubClient.Topic(conf.PullContactsSourceTopic())
	pullContactSourcePublisher := pubsub.NewPullContactSourcePublisher(pullContactsTopic)

	contactSourceDeletedPublisher := pubsub.NewContactSourceDeletedPublisher(pubsubClient.Topic(conf.ContactSourceDeletedTopic()))

	googleContactsFirestore := firestore.NewGoogleContactsFirestore(dbpool)
	peopleServiceFactory := servicefactory.NewPeopleServiceFactory(googleOAuthService)
	unifiedRepo := firestore.NewUnifiedContactFirestore(dbpool)
	contactLogRepo := firestore.NewContactLogFirestore(dbpool)

	linkSuggestionRepo := firestore.NewLinkSuggestionFirestore(dbpool)
	linkSuggestionService := application.NewLinkSuggestionService(unifiedRepo, linkSuggestionRepo)

	unifiedContactService := application.NewUnifiedContactService(unifiedRepo, linkSuggestionService, contactLogRepo)

	contactsApi := api.Create()

	csvImporter := application.NewCsvImporter(defaultContactFirestore, unifiedContactService)
	// configure routes

	defaultContactSource := defaultcontactsource.NewDefaultContactSource(defaultContactFirestore)
	googleContactSource := googlecontactsource.NewGoogleContactSource(googleContactsFirestore, contactSourceFirestore, peopleServiceFactory, googleContactsFirestore)
	contactSourceProvider := contactsource.NewContactSourceProvider(defaultContactSource, googleContactSource)

	contactSourceService := application.NewContactSourceService(googleOAuthService, contactSourceFirestore, pullContactSourcePublisher, unifiedContactService, userFirestore, contactSourceProvider, unifiedRepo, contactSourceDeletedPublisher)

	userService := application.NewUserService(userFirestore)
	contactService := application.NewContactService(userFirestore, defaultContactFirestore, unifiedContactFirestore, unifiedContactService, contactSourceProvider)

	tagService := application.NewTagService(tagFirestore)
	noteService := application.NewNoteService(noteFirestore)
	contactSearch := typesense.NewContactSearch(typesense.New(conf.TypesenseHost(), conf.TypesenseApiKey()))
	contactSearchService := application.NewContactSearchService(contactSearch, unifiedContactFirestore)
	routes.Users(contactsApi, userService)
	routes.Contacts(contactsApi, contactService, csvImporter)
	routes.Unified(contactsApi, unifiedContactService, linkSuggestionService, contactSearchService)
	routes.ContactSource(contactsApi, contactSourceService)
	routes.Tags(contactsApi, tagService)
	routes.Notes(contactsApi, noteService)

	srv := api.CreateServer(contactsApi)

	defer func() {
		l.Info("shutting down...")
		srv.Shutdown()
	}()

	l.Info("starting server...")

	if err := srv.Serve(); err != nil {
		log.Panicln("server err:", err)
	}
}
