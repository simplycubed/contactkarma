package search

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseEvent(t *testing.T) {
	rawJson := `{"addresses":{"nullValue":null},"birthdays":{"nullValue":null},"category":{"stringValue":"10"},"created_at":{"timestampValue":"2022-08-27T11:57:06.700502Z"},"email_addresses":{"nullValue":null},"genders":{"nullValue":null},"id":{"stringValue":"0FyMXFYxDNNnBD5fpuS4"},"last_contact":{"nullValue":null},"names":{"arrayValue":{"values":[{"mapValue":{"fields":{"display_name":{"stringValue":"John Doe"},"display_name_last_first":{"stringValue":"John Doe"},"given_name":{"stringValue":"John Doe"},"unstructured_name":{"stringValue":"John Doe"}}}}]}},"next_contact":{"timestampValue":"2022-08-27T11:57:06.700482Z"},"nicknames":{"nullValue":null},"occupations":{"nullValue":null},"organizations":{"nullValue":null},"origins":{"mapValue":{"fields":{"google:I5n3jzSbRg2SnrALJ2K1:c1819254288850202665":{"booleanValue":true}}}},"phone_numbers":{"arrayValue":{"values":[{"mapValue":{"fields":{"type":{"stringValue":"mobile"},"value":{"stringValue":"058 527 6679"}}}}]}},"photos":{"arrayValue":{"values":[{"mapValue":{"fields":{"default":{"booleanValue":true},"url":{"stringValue":"https://lh3.googleusercontent.com/cm/AATWAfufZClkaoX9Bf1XrL9nXCbabJbb5vg-DHA1h0Z0X5tKgjhJQKabQfVkXT6vYu1t=s100"}}}}]}},"relations":{"nullValue":null},"score":{"integerValue":"10"},"search_terms":{"arrayValue":{"values":[{"stringValue":"058 527 6679"},{"stringValue":"John Doe"}]}},"updated_at":{"timestampValue":"2022-08-27T11:57:06.700502Z"},"urls":{"nullValue":null}}`
	model := Unified{}
	err := json.Unmarshal([]byte(rawJson), &model)
	assert.Equal(t, true, err == nil)
	assert.Equal(t, 10, model.Score.IntegerValue)
	assert.Equal(t, model.Names.ArrayValue.Values[0].MapValues.Fields.GivenName.StringValue, "John Doe")
	assert.Equal(t, model.NextContact.TimestampValue.Year(), 2022)
}

func TestParseDate(t *testing.T) {
	input := "2022-09-15"
	date, _ := parseDate(input)
	assert.Equal(t, 2022, date.Year())
	assert.Equal(t, 9, int(date.Month()))
	assert.Equal(t, 15, date.Day())
}

func TestGetBirthdayNumber(t *testing.T) {
	input := "2022-09-15"
	date, _ := parseDate(input)
	birthdayNumber := GetBirthdayNumber(date)
	assert.Equal(t, 915, birthdayNumber)
	input = "2022-11-01"
	date, _ = parseDate(input)
	birthdayNumber = GetBirthdayNumber(date)
	assert.Equal(t, 1101, birthdayNumber)
}
