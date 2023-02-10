package integrations

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/simplycubed/contactkarma/integrations/modules/auth"
	"github.com/simplycubed/contactkarma/integrations/modules/search"
	"github.com/typesense/typesense-go/typesense"
)

func OnCreateAuthUser(ctx context.Context, e auth.AuthEvent) (err error) {
	project := os.Getenv("GCP_PROJECT")
	firestoreClient, err := firestore.NewClient(ctx, project)
	if err != nil {
		return
	}
	authService := auth.NewAuthService(firestoreClient)
	return authService.OnCreateAuthUser(ctx, e)
}

func OnWriteUnified(ctx context.Context, e search.FirestoreEvent) (err error) {
	host := os.Getenv("TYPESENSE_HOST")
	apiKey := os.Getenv("TYPESENSE_API_KEY")
	log.Println("Typesense keys", host, apiKey)
	typesenseClient := typesense.NewClient(typesense.WithServer(host), typesense.WithAPIKey(apiKey))
	searchService := search.NewSearchService(typesenseClient)
	return searchService.OnWriteUnified(ctx, e)
}
