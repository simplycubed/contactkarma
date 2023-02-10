package typesense

import "github.com/typesense/typesense-go/typesense"

func New(host string, apiKey string) *typesense.Client {
	return typesense.NewClient(typesense.WithServer(host), typesense.WithAPIKey(apiKey))
}
