package typesense

import (
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/simplycubed/contactkarma/contacts/conf"
	"github.com/simplycubed/contactkarma/contacts/domain"
	"github.com/simplycubed/contactkarma/contacts/test/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api/pointer"
)

var client *typesense.Client

func TestMain(m *testing.M) {
	testutils.LoadEnvFile()
	host, key := conf.TypesenseConfig()
	client = New(host, key)
	exitVal := m.Run()
	os.Exit(exitVal)
}

func addMockUnifiedSearch(searchService *ContactSearch, userId string, modifier func(v *domain.UnifiedSearch)) (modified domain.UnifiedSearch, err error) {
	unifiedId := "123" + strconv.Itoa(rand.Intn(1000))
	unified := domain.UnifiedSearch{
		ID:             domain.UnifiedSearchId(userId + "_" + unifiedId),
		UserId:         domain.UserID(userId),
		UnifiedId:      domain.UnifiedId(unifiedId),
		DisplayName:    "John Doe",
		Names:          []string{"john doe"},
		EmailAddresses: []string{"john@gmail.com"},
	}
	if modifier != nil {
		modifier(&unified)
	}
	_, err = searchService.Upsert(unified.UserId, unified)
	if err != nil {
		return
	}
	modified = unified
	return
}

func TestSearch(t *testing.T) {
	searchService := NewContactSearch(client)
	err := searchService.CreateCollection()
	assert.Equal(t, true, err == nil)
	defer func() {
		client.Collection(GetContactCollection()).Delete()
	}()
	userId := "123"
	mockContacts := []domain.UnifiedSearch{
		{
			DisplayName:    "Dwight Shrute",
			Names:          []string{"Dwight", "Shrute"},
			EmailAddresses: []string{"dwight@dundermifflin.com"},
			PhoneNumbers:   []string{"50900100"},
			Birthday:       pointer.Int64(time.Date(1990, 6, 1, 0, 0, 0, 0, time.UTC).Unix()),
			BirthdayNumber: pointer.Int(domain.GetBirthdayNumber(time.Date(1990, 6, 1, 0, 0, 0, 0, time.UTC))),
			LastContact:    pointer.Int64(time.Date(2010, 1, 4, 0, 0, 0, 0, time.UTC).Unix()),
			NextContact:    pointer.Int64(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC).Unix()),
			Score:          99,
			Category:       domain.A,
			Genders:        []string{"male"},
		},
		{
			DisplayName:    "Michael Scott",
			Names:          []string{"Michael", "Scott"},
			EmailAddresses: []string{"micheal@dundermifflin.com"},
			PhoneNumbers:   []string{"50900101"},
			Birthday:       pointer.Int64(time.Date(1994, 1, 1, 0, 0, 0, 0, time.UTC).Unix()),
			BirthdayNumber: pointer.Int(domain.GetBirthdayNumber(time.Date(1994, 1, 1, 0, 0, 0, 0, time.UTC))),
			LastContact:    pointer.Int64(time.Date(2010, 1, 3, 0, 0, 0, 0, time.UTC).Unix()),
			NextContact:    pointer.Int64(time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC).Unix()),
			Score:          90,
			Category:       domain.A,
			Genders:        []string{"male"},
		},
		{
			DisplayName:    "Jim Halpert",
			Names:          []string{"Jim", "Halpert"},
			EmailAddresses: []string{"halpert@dundermifflin.com"},
			PhoneNumbers:   []string{"50900102"},
			Birthday:       pointer.Int64(time.Date(2000, 12, 12, 0, 0, 0, 0, time.UTC).Unix()),
			BirthdayNumber: pointer.Int(domain.GetBirthdayNumber(time.Date(2000, 12, 12, 0, 0, 0, 0, time.UTC))),
			LastContact:    pointer.Int64(time.Date(2010, 1, 2, 0, 0, 0, 0, time.UTC).Unix()),
			NextContact:    pointer.Int64(time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC).Unix()),
			Score:          70,
			Category:       domain.B,
			Genders:        []string{"male"},
		},
		{
			DisplayName:    "Pam Beesly Halpert",
			Names:          []string{"Pam", "Beesly", "Halpert"},
			EmailAddresses: []string{"pam@dundermifflin.com"},
			PhoneNumbers:   []string{"5090010349"},
			Birthday:       pointer.Int64(time.Date(1999, 8, 19, 0, 0, 0, 0, time.UTC).Unix()),
			BirthdayNumber: pointer.Int(domain.GetBirthdayNumber(time.Date(1999, 8, 19, 0, 0, 0, 0, time.UTC))),
			LastContact:    pointer.Int64(time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC).Unix()),
			NextContact:    pointer.Int64(time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC).Unix()),
			Score:          75,
			Category:       domain.B,
			Genders:        []string{"female"},
		},
	}
	docs := []domain.UnifiedSearch{}
	for _, contact := range mockContacts {
		search, err := addMockUnifiedSearch(searchService, userId, func(v *domain.UnifiedSearch) {
			v.EmailAddresses = contact.EmailAddresses
			v.Names = contact.Names
			v.DisplayName = contact.DisplayName
			v.PhoneNumbers = contact.PhoneNumbers
			v.Birthday = contact.Birthday
			v.LastContact = contact.LastContact
			v.NextContact = contact.NextContact
			v.Score = contact.Score
			v.Category = contact.Category
			v.BirthdayNumber = contact.BirthdayNumber
			v.Genders = contact.Genders
		})
		assert.Equal(t, nil, err)
		docs = append(docs, search)
	}
	dwight, michael, jim, pam := docs[0], docs[1], docs[2], docs[3]
	defaultResults, err := searchService.Search(domain.UserID(userId), domain.SearchParams{})
	assert.Equal(t, true, err == nil)
	assert.Equal(t, 4, len(defaultResults))

	// sort by score
	sortedResults, err := searchService.Search(domain.UserID(userId), domain.SearchParams{Sort: []domain.Sort{{Field: "score", Order: domain.ASC}}})
	assert.Equal(t, true, err == nil)
	assert.Equal(t, 4, len(sortedResults))
	assert.Equal(t, mapToIds(sortedResults), mapToIds([]domain.UnifiedSearch{jim, pam, michael, dwight}))

	sortedResults, err = searchService.Search(domain.UserID(userId), domain.SearchParams{Sort: []domain.Sort{{Field: "score", Order: domain.Desc}}})
	assert.Equal(t, true, err == nil)
	assert.Equal(t, 4, len(sortedResults))
	assert.Equal(t, mapToIds(sortedResults), mapToIds([]domain.UnifiedSearch{dwight, michael, pam, jim}))

	// sort by next contact
	sortedResults, err = searchService.Search(domain.UserID(userId), domain.SearchParams{Sort: []domain.Sort{{Field: "next_contact", Order: domain.ASC}}})
	assert.Equal(t, true, err == nil)
	assert.Equal(t, 4, len(sortedResults))
	assert.Equal(t, mapToIds(sortedResults), mapToIds([]domain.UnifiedSearch{dwight, pam, michael, jim}))

	// sort by last contact
	sortedResults, err = searchService.Search(domain.UserID(userId), domain.SearchParams{Sort: []domain.Sort{{Field: "last_contact", Order: domain.Desc}}})
	assert.Equal(t, true, err == nil)
	assert.Equal(t, 4, len(sortedResults))
	assert.Equal(t, mapToIds(sortedResults), mapToIds([]domain.UnifiedSearch{dwight, michael, jim, pam}))

	// sort by birthday
	sortedResults, err = searchService.Search(domain.UserID(userId), domain.SearchParams{
		Sort: []domain.Sort{{Field: "birthday", Order: domain.ASC}},
		Filters: []domain.ContactSearchFilter{
			{
				Field:    "birthday",
				Operator: ">=",
				Value:    "01-01",
			},
		},
	})
	assert.Equal(t, true, err == nil)
	assert.Equal(t, 4, len(sortedResults))
	assert.Equal(t, mapToIds(sortedResults), mapToIds([]domain.UnifiedSearch{michael, dwight, pam, jim}))

	// Query by name
	queryResults, err := searchService.Search(domain.UserID(userId), domain.SearchParams{Query: "halpert"})
	assert.Equal(t, true, err == nil)
	assert.Equal(t, 2, len(queryResults))
	assert.ElementsMatch(t, mapToIds(queryResults), mapToIds([]domain.UnifiedSearch{pam, jim}))

	// Query by email
	queryResults, err = searchService.Search(domain.UserID(userId), domain.SearchParams{Query: "halpert@dundermifflin.com"})
	assert.Equal(t, true, err == nil)
	assert.Equal(t, 1, len(queryResults))
	assert.ElementsMatch(t, mapToIds(queryResults), mapToIds([]domain.UnifiedSearch{jim}))

	//Query by phone
	queryResults, err = searchService.Search(domain.UserID(userId), domain.SearchParams{Query: "5090010349"})
	assert.Equal(t, true, err == nil)
	assert.Equal(t, 1, len(queryResults))
	assert.ElementsMatch(t, mapToIds(queryResults), mapToIds([]domain.UnifiedSearch{pam}))

	//Filter by category
	queryResults, err = searchService.Search(domain.UserID(userId), domain.SearchParams{Filters: []domain.ContactSearchFilter{
		{
			Field:    domain.CategoryField,
			Operator: "=",
			Value:    domain.B.String(),
		},
	}})
	assert.Equal(t, true, err == nil)
	assert.Equal(t, 2, len(queryResults))
	assert.ElementsMatch(t, mapToIds(queryResults), mapToIds([]domain.UnifiedSearch{pam, jim}))

	//Filter by gender
	queryResults, err = searchService.Search(domain.UserID(userId), domain.SearchParams{Filters: []domain.ContactSearchFilter{
		{
			Field:    domain.GenderField,
			Operator: "=",
			Value:    "male",
		},
	}})
	assert.Equal(t, true, err == nil)
	assert.Equal(t, 3, len(queryResults))
	assert.ElementsMatch(t, mapToIds(queryResults), mapToIds([]domain.UnifiedSearch{dwight, michael, jim}))
}

func mapToIds(values []domain.UnifiedSearch) (ids []string) {
	for _, v := range values {
		ids = append(ids, string(v.UnifiedId))
	}
	return
}
