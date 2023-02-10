package typesense

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/simplycubed/contactkarma/contacts/domain"
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
	"github.com/typesense/typesense-go/typesense/api/pointer"
)

type ContactSearch struct {
	client *typesense.Client
}

func NewContactSearch(client *typesense.Client) *ContactSearch {
	return &ContactSearch{client: client}
}

func (c *ContactSearch) CreateCollection() (err error) {
	_, err = c.client.Collections().Create(&api.CollectionSchema{
		DefaultSortingField: new(string),
		Fields: []api.Field{
			{
				Facet:    pointer.False(),
				Index:    pointer.True(),
				Name:     "user_id",
				Optional: pointer.False(),
				Type:     "string",
			},
			{
				Facet:    pointer.False(),
				Index:    pointer.True(),
				Name:     "unified_id",
				Optional: pointer.False(),
				Type:     "string",
			},
			{
				Facet:    pointer.False(),
				Index:    pointer.True(),
				Name:     "display_name",
				Infix:    pointer.True(),
				Optional: pointer.True(),
				Type:     "string",
			},
			{
				Facet:    pointer.False(),
				Index:    pointer.True(),
				Name:     "names",
				Infix:    pointer.True(),
				Optional: pointer.True(),
				Type:     "string[]",
			},
			{
				Facet:    pointer.False(),
				Index:    pointer.True(),
				Name:     "email_addresses",
				Optional: pointer.True(),
				Type:     "string[]",
				Infix:    pointer.True(),
			},
			{
				Facet:    pointer.False(),
				Index:    pointer.True(),
				Name:     "addresses",
				Optional: pointer.True(),
				Type:     "string[]",
				Infix:    pointer.True(),
			},
			{
				Facet:    pointer.False(),
				Index:    pointer.True(),
				Name:     "genders",
				Optional: pointer.True(),
				Type:     "string[]",
				Infix:    pointer.True(),
			},
			{
				Facet:    pointer.False(),
				Index:    pointer.True(),
				Name:     "occupations",
				Optional: pointer.True(),
				Type:     "string[]",
				Infix:    pointer.True(),
			},
			{
				Facet:    pointer.False(),
				Index:    pointer.True(),
				Name:     "organizations",
				Optional: pointer.True(),
				Type:     "string[]",
				Infix:    pointer.True(),
			},
			{
				Facet:    pointer.False(),
				Index:    pointer.True(),
				Name:     "phone_numbers",
				Optional: pointer.True(),
				Type:     "string[]",
				Infix:    pointer.True(),
			},
			{
				Facet:    pointer.False(),
				Index:    pointer.True(),
				Name:     "birthday",
				Optional: pointer.True(),
				Type:     "int64",
				Sort:     pointer.True(),
			},
			{
				Facet:    pointer.False(),
				Index:    pointer.True(),
				Name:     "birthday_number",
				Optional: pointer.True(),
				Type:     "int64",
				Sort:     pointer.True(),
			},
			{
				Facet:    pointer.False(),
				Index:    pointer.True(),
				Name:     "next_contact",
				Optional: pointer.True(),
				Type:     "int64",
				Sort:     pointer.True(),
			},
			{
				Facet:    pointer.False(),
				Index:    pointer.True(),
				Name:     "last_contact",
				Optional: pointer.True(),
				Type:     "int64",
				Sort:     pointer.True(),
			},
			{
				Facet:    pointer.False(),
				Index:    pointer.True(),
				Name:     "score",
				Optional: pointer.True(),
				Type:     "int64",
				Sort:     pointer.True(),
			},
			{
				Facet:    pointer.False(),
				Index:    pointer.True(),
				Name:     "category",
				Optional: pointer.True(),
				Type:     "string",
			},
		},
		Name:            GetContactCollection(),
		SymbolsToIndex:  &[]string{},
		TokenSeparators: &[]string{},
	})
	return
}

func (c *ContactSearch) Upsert(userId domain.UserID, unified domain.UnifiedSearch) (unifiedIds []domain.UnifiedId, err error) {
	_, err = c.client.Collection(GetContactCollection()).Documents().Upsert(unified)
	return
}

func (c *ContactSearch) Search(userId domain.UserID, searchParams domain.SearchParams) (documents []domain.UnifiedSearch, err error) {
	filterBy := "user_id:=" + userId.String()
	for _, filter := range searchParams.Filters {
		value := filter.Value
		field := filter.Field.String()
		if field == "birthday" {
			birthdayNumber, err := domain.GetBirthdayNumberFromString(filter.Value)
			if err != nil {
				return nil, err
			}
			field = "birthday_number"
			value = strconv.Itoa(birthdayNumber)
		}
		filterBy += " && " + field + ":" + filter.Operator.String() + value
	}
	params := &api.SearchCollectionParams{
		FilterBy: &filterBy,
		QueryBy:  "display_name,names,email_addresses,phone_numbers,addresses,organizations,occupations",
		Q:        searchParams.Query,
		// TODO: infix param is missing in the client.
	}
	sortBy := []string{}
	for _, sort := range searchParams.Sort {
		field := sort.Field
		order := string(sort.Order)
		// TODO: what if user wants to sort by birthdates, not numbers?
		if field == "birthday" {
			field = "birthday_number"
		}
		sortBy = append(sortBy, field+":"+order)

	}
	if len(sortBy) > 0 {
		params.SortBy = pointer.String(strings.Join(sortBy, ","))
	}
	if searchParams.Page > 0 {
		params.Page = &searchParams.Page
	}
	if searchParams.PerPage > 0 {
		params.PerPage = &searchParams.PerPage
	}
	results, err := c.client.Collection(GetContactCollection()).Documents().Search(params)
	if err != nil {
		return
	}
	v := *results.Hits
	for _, hit := range v {
		if hit.Document != nil {
			doc := *hit.Document
			jsonRaw, err := json.Marshal(doc)
			if err != nil {
				return nil, err
			}
			searchDoc := domain.UnifiedSearch{}
			err = json.Unmarshal(jsonRaw, &searchDoc)
			if err != nil {
				return nil, err
			}
			documents = append(documents, searchDoc)
		}
	}
	return
}
