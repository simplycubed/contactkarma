package search

import (
	"encoding/json"
	"strconv"
	"time"
)

type UnifiedSearchId string
type UnifiedId string
type UserID string

type ContactCatgeory string

const A ContactCatgeory = "A"
const B ContactCatgeory = "B"
const C ContactCatgeory = "C"
const D ContactCatgeory = "D"

type UnifiedSearch struct {
	ID             UnifiedSearchId `json:"id"`
	UnifiedId      UnifiedId       `json:"unified_id"`
	UserId         UserID          `json:"user_id"`
	DisplayName    string          `json:"display_name"`
	Names          []string        `json:"names"`
	Nicknames      []string        `json:"nicknames"`
	EmailAddresses []string        `json:"email_addresses" fakesize:"1"`
	Addresses      []string        `json:"addresses"`
	Genders        []string        `json:"genders"`
	Occupations    []string        `json:"occupations"`
	Organizations  []string        `json:"organizations"`
	PhoneNumbers   []string        `json:"phone_numbers"`
	Birthday       *int64          `json:"birthday"`
	NextContact    *int64          `json:"next_contact"` // unix timestamp
	LastContact    *int64          `json:"last_contact"` // unix timestamp
	Score          int             `json:"score"`
	Category       ContactCatgeory `json:"category"`

	// Derived
	// Eg: Number for "2020-01-01" is 101 (year ignored) and "1994-12-27" is 1227
	BirthdayNumber *int `json:"birthday_number"`
}

func GetBirthdayNumber(date time.Time) int {
	return int(date.Month())*100 + date.Day()
}

type StringValue struct {
	StringValue string `json:"stringValue"`
}

type IntegerValue struct {
	IntegerValue int `json:"integerValue"`
}

type TimestampValue struct {
	TimestampValue *time.Time `json:"timestampValue"`
}

func parseDate(date string) (t time.Time, err error) {
	layout := "2006-01-02"
	t, err = time.Parse(layout, date)
	return
}

func (t *TimestampValue) UnmarshalJSON(data []byte) (err error) {
	vMap := map[string]json.RawMessage{}
	err = json.Unmarshal(data, &vMap)
	if err != nil {
		return
	}
	_, ok := vMap["nullValue"]
	if ok {
		*t = TimestampValue{TimestampValue: nil}
		return nil
	}
	timeString := vMap["timestampValue"]
	if string(timeString) != "" {
		time := time.Time{}
		err = json.Unmarshal(timeString, &time)
		if err != nil {
			return err
		}
		*t = TimestampValue{TimestampValue: &time}
		return
	}

	*t = TimestampValue{TimestampValue: nil}
	return nil
}

func (t *IntegerValue) UnmarshalJSON(data []byte) (err error) {
	vMap := map[string]string{}
	err = json.Unmarshal(data, &vMap)
	if err != nil {
		return
	}
	valueString, ok := vMap["integerValue"]
	if ok {
		val, err := strconv.Atoi(valueString)
		if err != nil {
			return err
		}
		*t = IntegerValue{IntegerValue: val}
		return nil
	}
	return nil
}

// TODO: Update the current model once https://github.com/googleapis/google-cloud-go/issues/1438 is resolved.
type Unified struct {
	ID             StringValue       `firestore:"id" json:"id"`
	DisplayName    StringValue       `firestore:"display_name" json:"display_name"`
	Names          UsernamesArray    `firestore:"names" json:"names"`
	EmailAddresses EmailAddressArray `firestore:"email_addresses" json:"email_addresses" fakesize:"1"`
	Nicknames      NicknameArray     `firestore:"nicknames" json:"nicknames" fakesize:"1"`
	Addresses      AddressArray      `firestore:"addresses" json:"addresses" fakesize:"1"`
	Genders        GenderArray       `firestore:"genders" json:"genders" fakesize:"1"`
	Occupations    OccupationArray   `firestore:"occupations" json:"occupations" fakesize:"1"`
	Organizations  OrganizationArray `firestore:"organizations" json:"organizations" fakesize:"1"`
	PhoneNumbers   PhoneNumbersArray `firestore:"phone_numbers" json:"phone_numbers" fakesize:"1"`
	Birthdays      BirthdaysArray    `firestore:"birthdays" json:"birthdays" fakesize:"1"`
	NextContact    TimestampValue    `firestore:"next_contact" json:"next_contact"`
	LastContact    TimestampValue    `firestore:"last_contact" json:"last_contact"`
	Score          IntegerValue      `firestore:"score" json:"score"`
	Category       StringValue       `firestore:"category" json:"category"`
}

type UsernamesArray struct {
	ArrayValue struct {
		Values []struct {
			MapValues struct {
				Fields UserNames `json:"fields"`
			} `json:"mapValue"`
		} `json:"values"`
	} `json:"arrayValue"`
}

type UserNames struct {
	FamilyName StringValue `firestore:"family_name,omitempty" json:"family_name,omitempty" fake:"{lastname}"`
	GivenName  StringValue `firestore:"given_name,omitempty" json:"given_name,omitempty" fake:"{lastname}"`
	MiddleName StringValue `firestore:"middle_name,omitempty" json:"middle_name,omitempty" fake:"{lastname}"`
}

type EmailAddressArray struct {
	ArrayValue struct {
		Values []struct {
			MapValues struct {
				Fields EmailAddress `json:"fields"`
			} `json:"mapValue"`
		} `json:"values"`
	} `json:"arrayValue"`
}
type EmailAddress struct {
	Type  StringValue `firestore:"type,omitempty" json:"type,omitempty"`
	Value StringValue `firestore:"value,omitempty" json:"value,omitempty" fake:"{email}"`
}

type PhoneNumbersArray struct {
	ArrayValue struct {
		Values []struct {
			MapValues struct {
				Fields PhoneNumber `json:"fields"`
			} `json:"mapValue"`
		} `json:"values"`
	} `json:"arrayValue"`
}
type PhoneNumber struct {
	Type  StringValue `firestore:"type,omitempty" json:"type,omitempty"`
	Value StringValue `firestore:"value,omitempty" json:"value,omitempty" fake:"{email}"`
}

type BirthdaysArray struct {
	ArrayValue struct {
		Values []struct {
			MapValues struct {
				Fields Birthday `json:"fields"`
			} `json:"mapValue"`
		} `json:"values"`
	} `json:"arrayValue"`
}

type Birthday struct {
	Date StringValue `firestore:"date,omitempty" json:"date,omitempty" fake:"{year}-{month}-{day}"`
	Text StringValue `firestore:"text,omitempty" json:"text,omitempty" fake:"{year}-{month}-{day}"`
}

type AddressArray struct {
	ArrayValue struct {
		Values []struct {
			MapValues struct {
				Fields Address `json:"fields"`
			} `json:"mapValue"`
		} `json:"values"`
	} `json:"arrayValue"`
}

type Address struct {
	City            StringValue `firestore:"city,omitempty" json:"city,omitempty" fake:"{city}"`
	Country         StringValue `firestore:"country,omitempty" json:"country,omitempty" fake:"{country}"`
	CountryCode     StringValue `firestore:"country_code,omitempty" json:"country_code,omitempty" fake:"{countryabr}"`
	ExtendedAddress StringValue `firestore:"extended_address,omitempty" json:"extended_address,omitempty" fake:"{street}, {streetnumber}, {city}"`
	PoBox           StringValue `firestore:"po_box,omitempty" json:"po_box,omitempty"`
	PostalCode      StringValue `firestore:"postal_code,omitempty" json:"postal_code,omitempty" fake:"{zip}"`
	Region          StringValue `firestore:"region,omitempty" json:"region,omitempty" fake:"{state}"`
	StreetAddress   StringValue `firestore:"street_address,omitempty" json:"street_address,omitempty" fake:"{street}"`
	Type            StringValue `firestore:"type,omitempty" json:"type,omitempty"`
}

type GenderArray struct {
	ArrayValue struct {
		Values []struct {
			MapValues struct {
				Fields Gender `json:"fields"`
			} `json:"mapValue"`
		} `json:"values"`
	} `json:"arrayValue"`
}
type Gender struct {
	AddressMeAs StringValue `firestore:"address_me_as,omitempty" json:"address_me_as,omitempty" fake:"{pronounpersonal}"`
	Value       StringValue `firestore:"value,omitempty" json:"value,omitempty" fake:"{gender}"`
}

type OccupationArray struct {
	ArrayValue struct {
		Values []struct {
			MapValues struct {
				Fields Occupation `json:"fields"`
			} `json:"mapValue"`
		} `json:"values"`
	} `json:"arrayValue"`
}
type Occupation struct {
	Value StringValue `firestore:"value,omitempty" json:"value,omitempty" fake:"{jobtitle}"`
}

type OrganizationArray struct {
	ArrayValue struct {
		Values []struct {
			MapValues struct {
				Fields Organization `json:"fields"`
			} `json:"mapValue"`
		} `json:"values"`
	} `json:"arrayValue"`
}
type Organization struct {
	Department     StringValue `firestore:"department,omitempty" json:"department,omitempty"`
	Domain         StringValue `firestore:"domain,omitempty" json:"domain,omitempty"`
	EndDate        StringValue `firestore:"end_date,omitempty" json:"end_date,omitempty" fake:"{year}-{month}-{day}"`
	JobDescription StringValue `firestore:"job_description,omitempty" json:"job_description,omitempty" fake:"{sentence:3}"`
	Location       StringValue `firestore:"location,omitempty" json:"location,omitempty"`
	Name           StringValue `firestore:"name,omitempty" json:"name,omitempty" fake:"{company}"`
	PhoneticName   StringValue `firestore:"phonetic_name,omitempty" json:"phonetic_name,omitempty" fake:"{company}"`
	StartDate      StringValue `firestore:"start_date,omitempty" json:"start_date,omitempty" fake:"{year}-{month}-{day}"`
	Symbol         StringValue `firestore:"symbol,omitempty" json:"symbol,omitempty"`
	Title          StringValue `firestore:"title,omitempty" json:"title,omitempty" fake:"{jobtitle}"`
	Type           StringValue `firestore:"type,omitempty" json:"type,omitempty"`
}

type NicknameArray struct {
	ArrayValue struct {
		Values []struct {
			MapValues struct {
				Fields Nickname `json:"fields"`
			} `json:"mapValue"`
		} `json:"values"`
	} `json:"arrayValue"`
}
type Nickname struct {
	Value StringValue `firestore:"value,omitempty" json:"value,omitempty" fake:"{firstname}"`
}
