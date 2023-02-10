package search

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api/pointer"
)

// FirestoreEvent is the payload of a Firestore event.
type FirestoreEvent struct {
	OldValue   FirestoreValue `json:"oldValue"`
	Value      FirestoreValue `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

// FirestoreValue holds Firestore fields.
type FirestoreValue struct {
	CreateTime time.Time `json:"createTime"`
	// Fields is the data for this value. The type depends on the format of your
	// database. Log an interface{} value and inspect the result to see a JSON
	// representation of your database fields.
	Fields     *Unified  `json:"fields"`
	Name       string    `json:"name"`
	UpdateTime time.Time `json:"updateTime"`
}

type SearchService struct {
	client *typesense.Client
}

func NewSearchService(client *typesense.Client) *SearchService {
	return &SearchService{client: client}
}

func parsePathIds(documentPath string) (userId string, unifiedId string) {
	fullPath := strings.Split(documentPath, "/documents/")[1]
	log.Println("document path:", fullPath)
	pathParts := strings.Split(fullPath, "/")
	userId = pathParts[1]
	unifiedId = pathParts[3]
	return
}

// OnWriteUnified
func (service *SearchService) OnWriteUnified(ctx context.Context, e FirestoreEvent) (err error) {
	log.Printf("Function triggered: %+v", e)
	log.Printf("value: %+v", e.Value.Fields)

	if e.Value.Fields == nil {
		userId, unifiedId := parsePathIds(e.OldValue.Name)
		log.Println("ids", userId, unifiedId)
		searchDocumentId := userId + "_" + unifiedId
		// delete
		_, err = service.client.Collection("contacts").Document(searchDocumentId).Delete()
	} else {
		userId, unifiedId := parsePathIds(e.Value.Name)
		log.Println("ids", userId, unifiedId)
		log.Printf("raw: %+v", *e.Value.Fields)
		unified := e.Value.Fields
		searchDocument := UnifiedSearch{}
		searchDocument.ID = UnifiedSearchId(userId + "_" + unifiedId)
		searchDocument.UserId = UserID(userId)
		searchDocument.UnifiedId = UnifiedId(unifiedId)
		searchDocument.DisplayName = unified.DisplayName.StringValue
		for _, email := range unified.EmailAddresses.ArrayValue.Values {
			if email.MapValues.Fields.Value.StringValue != "" {
				searchDocument.EmailAddresses = append(searchDocument.EmailAddresses, email.MapValues.Fields.Value.StringValue)
			}
		}
		for _, phone := range unified.PhoneNumbers.ArrayValue.Values {
			if phone.MapValues.Fields.Value.StringValue != "" {
				searchDocument.PhoneNumbers = append(searchDocument.PhoneNumbers, phone.MapValues.Fields.Value.StringValue)
			}
		}
		if len(unified.Birthdays.ArrayValue.Values) > 0 {
			birthday := unified.Birthdays.ArrayValue.Values[0].MapValues.Fields.Date
			if birthday.StringValue != "" {
				date, err := parseDate(birthday.StringValue)
				if err == nil {
					searchDocument.Birthday = pointer.Int64(date.Unix()) // TODO: replace with birthdayNumber if no real use for full birthday
					searchDocument.BirthdayNumber = pointer.Int(GetBirthdayNumber(date))
				}
			}
		}
		for _, names := range unified.Names.ArrayValue.Values {
			if names.MapValues.Fields.FamilyName.StringValue != "" {
				searchDocument.Names = append(searchDocument.Names, names.MapValues.Fields.FamilyName.StringValue)
			}
			if names.MapValues.Fields.GivenName.StringValue != "" {
				searchDocument.Names = append(searchDocument.Names, names.MapValues.Fields.GivenName.StringValue)
			}
			if names.MapValues.Fields.MiddleName.StringValue != "" {
				searchDocument.Names = append(searchDocument.Names, names.MapValues.Fields.MiddleName.StringValue)
			}
		}
		for _, names := range unified.Nicknames.ArrayValue.Values {
			if names.MapValues.Fields.Value.StringValue != "" {
				searchDocument.Nicknames = append(searchDocument.Nicknames, names.MapValues.Fields.Value.StringValue)
			}
		}
		for _, address := range unified.Addresses.ArrayValue.Values {
			addressString := address.MapValues.Fields.City.StringValue + " " + address.MapValues.Fields.Region.StringValue + " " + address.MapValues.Fields.Country.StringValue
			if strings.TrimSpace(addressString) != "" {
				searchDocument.Addresses = append(searchDocument.Addresses, addressString)
			}
		}
		for _, gender := range unified.Genders.ArrayValue.Values {
			if gender.MapValues.Fields.Value.StringValue != "" {
				searchDocument.Genders = append(searchDocument.Genders, gender.MapValues.Fields.Value.StringValue)
			}
		}
		for _, occupation := range unified.Occupations.ArrayValue.Values {
			if occupation.MapValues.Fields.Value.StringValue != "" {
				searchDocument.Occupations = append(searchDocument.Occupations, occupation.MapValues.Fields.Value.StringValue)
			}
		}
		for _, organization := range unified.Organizations.ArrayValue.Values {
			if organization.MapValues.Fields.Name.StringValue != "" {
				searchDocument.Organizations = append(searchDocument.Organizations, organization.MapValues.Fields.Name.StringValue)
			}
		}
		if unified.NextContact.TimestampValue != nil {
			searchDocument.NextContact = pointer.Int64(unified.NextContact.TimestampValue.Unix())
		}
		if unified.LastContact.TimestampValue != nil {
			searchDocument.LastContact = pointer.Int64(unified.LastContact.TimestampValue.Unix())
		}
		if unified.Category.StringValue != "" {
			searchDocument.Category = ContactCatgeory(unified.Category.StringValue)
		}
		if unified.Score.IntegerValue > 0 {
			searchDocument.Score = unified.Score.IntegerValue
		}

		_, err = service.client.Collection("contacts").Documents().Upsert(searchDocument)
	}

	if err != nil {
		log.Println(err)
	}

	return
}
