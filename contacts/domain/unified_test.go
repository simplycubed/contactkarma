package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	unified := Unified{
		Names: []*UserNames{
			{
				DisplayName: "John",
				GivenName:   "John",
				MiddleName:  "",
				FamilyName:  "Doe",
			},
		},
		Nicknames: []Nickname{
			{
				Value: "jhonny",
			},
		},
		Addresses: []Address{
			{
				City: "Dubai",
			},
		},
		Birthdays: []Birthday{
			{
				Date: "2000-01-01",
			},
		},
		EmailAddresses: []EmailAddress{{
			Value: "john@gmail.com",
		}},
		Genders: []Gender{{
			Value: "male",
		}},
		Occupations: []Occupation{{
			Value: "Software Engineer",
		}},
		PhoneNumbers: []PhoneNumber{{
			Value: "971508547300",
		}},
		Photos: []Photo{{
			Url: "https://url-1",
		}},
		Relations: []Relation{{
			Person: "Father",
		}},
		Urls: []Url{{
			Value: "https://url-1",
		}},
		Organizations: []Organization{{
			Name: "Google",
		}},
		Origins: map[string]bool{
			"Origin-1": true,
		},
	}
	now := time.Now()
	duplicate := Unified{
		Names: []*UserNames{
			{
				DisplayName: "john",
				GivenName:   "john",
				MiddleName:  "Middle",
				FamilyName:  "",
			},
		},
		Nicknames: []Nickname{
			{
				Value: "jhonny",
			},
			{
				Value: "doey",
			},
		},
		Addresses: []Address{
			{
				City: "Delhi",
			},
			{
				City: "Dubai",
			},
		},
		Birthdays: []Birthday{
			{
				Date: "2000-01-01",
			},
			{
				Date: "2000-02-01",
			},
		},
		EmailAddresses: []EmailAddress{{
			Value: "john@gmail.com",
		}, {
			Value: "doe@gmail.com",
		}},
		Genders: []Gender{{
			Value: "male",
		}, {
			Value: "none",
		}},
		Occupations: []Occupation{{
			Value: "Software Engineer",
		}, {
			Value: "Devops Engineer",
		}},
		PhoneNumbers: []PhoneNumber{{
			Value: "971508547300",
		}, {
			Value: "971508547999",
		}},
		Photos: []Photo{{
			Url: "https://url-1",
		}, {
			Url: "https://url-2",
		}},
		Relations: []Relation{{
			Person: "Father",
		}, {
			Person: "Dad",
		}},
		Urls: []Url{{
			Value: "https://url-1",
		}, {
			Value: "https://url-2",
		}},
		Organizations: []Organization{{
			Name: "Google",
		}, {
			Name: "Amazon",
		}},
		LastContact: &now,
		NextContact: &now,
		Score:       100,
		Category:    A,
		Origins: map[string]bool{
			"Origin-2": true,
		},
	}

	unified.Merge(duplicate)
	assert.Equal(t, "John", unified.Names[0].DisplayName)
	assert.Equal(t, "Middle", unified.Names[0].MiddleName)
	assert.Equal(t, "Doe", unified.Names[0].FamilyName)
	assert.Equal(t, "jhonny", unified.Nicknames[0].Value)
	assert.Equal(t, "doey", unified.Nicknames[1].Value)
	assert.Equal(t, "Dubai", unified.Addresses[0].City)
	assert.Equal(t, "Delhi", unified.Addresses[1].City)
	assert.Equal(t, "2000-01-01", unified.Birthdays[0].Date)
	assert.Equal(t, "2000-02-01", unified.Birthdays[1].Date)
	assert.Equal(t, "john@gmail.com", unified.EmailAddresses[0].Value)
	assert.Equal(t, "doe@gmail.com", unified.EmailAddresses[1].Value)

	assert.Equal(t, "male", unified.Genders[0].Value)
	assert.Equal(t, "none", unified.Genders[1].Value)

	assert.Equal(t, "Software Engineer", unified.Occupations[0].Value)
	assert.Equal(t, "Devops Engineer", unified.Occupations[1].Value)

	assert.Equal(t, "971508547300", unified.PhoneNumbers[0].Value)
	assert.Equal(t, "971508547999", unified.PhoneNumbers[1].Value)

	assert.Equal(t, "https://url-1", unified.Photos[0].Url)
	assert.Equal(t, "https://url-2", unified.Photos[1].Url)

	assert.Equal(t, "Father", unified.Relations[0].Person)
	assert.Equal(t, "Dad", unified.Relations[1].Person)

	assert.Equal(t, "https://url-1", unified.Urls[0].Value)
	assert.Equal(t, "https://url-2", unified.Urls[1].Value)

	assert.Equal(t, "Google", unified.Organizations[0].Name)
	assert.Equal(t, "Amazon", unified.Organizations[1].Name)

	assert.Equal(t, now.Format(time.RFC3339), unified.LastContact.Format(time.RFC3339))
	assert.Equal(t, now.Format(time.RFC3339), unified.NextContact.Format(time.RFC3339))

	assert.Equal(t, 100, unified.Score)
	assert.Equal(t, A, unified.Category)

	assert.Equal(t, true, unified.Origins["Origin-1"])
	assert.Equal(t, true, unified.Origins["Origin-2"])

	assert.Equal(t, true, len(unified.SearchTerms) > 0)
}
