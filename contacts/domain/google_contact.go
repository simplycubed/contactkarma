package domain

import (
	"strings"
	"time"

	"google.golang.org/api/people/v1"
)

type GoogleContact struct {
	ID             ContactID      `firestore:"id" json:"id"` // person id from google person
	Names          []*UserNames   `firestore:"names" json:"names"`
	Nicknames      []Nickname     `firestore:"nicknames" json:"nicknames" fakesize:"1"`
	Addresses      []Address      `firestore:"addresses" json:"addresses" fakesize:"1"`
	Birthdays      []Birthday     `firestore:"birthdays" json:"birthdays" fakesize:"1"`
	EmailAddresses []EmailAddress `firestore:"email_addresses" json:"email_addresses" fakesize:"1"`
	Genders        []Gender       `firestore:"genders" json:"genders" fakesize:"1"`
	Occupations    []Occupation   `firestore:"occupations" json:"occupations" fakesize:"1"`
	PhoneNumbers   []PhoneNumber  `firestore:"phone_numbers" json:"phone_numbers" fakesize:"1"`
	Photos         []Photo        `firestore:"photos" json:"photos" fakesize:"1"`
	Relations      []Relation     `firestore:"relations" json:"relations" fakesize:"1"`
	Urls           []Url          `firestore:"urls" json:"urls" fakesize:"1"`
	Organizations  []Organization `firestore:"organizations" json:"organizations" fakesize:"1"`
	CreatedAt      time.Time      `firestore:"created_at" json:"created_at"`
	UpdatedAt      time.Time      `firestore:"updated_at" json:"updated_at"`

	ResourceName string                 `firestore:"resource_name" json:"resource_name,omitempty"`
	Metadata     *people.PersonMetadata `firestore:"metadata" json:"metadata,omitempty"`
}

func (contact *GoogleContact) FromGooglePerson(person *people.Person) {

	for _, address := range person.Addresses {
		contact.Addresses = append(contact.Addresses, Address{
			City:            address.City,
			Country:         address.Country,
			CountryCode:     address.CountryCode,
			ExtendedAddress: address.ExtendedAddress,
			PoBox:           address.PoBox,
			PostalCode:      address.PostalCode,
			Region:          address.Region,
			StreetAddress:   address.StreetAddress,
			Type:            address.Type,
		})
	}

	for _, birthday := range person.Birthdays {
		contact.Birthdays = append(contact.Birthdays, Birthday{
			Date: getDate(birthday.Date),
			Text: birthday.Text,
		})
	}

	for _, emailAddress := range person.EmailAddresses {
		contact.EmailAddresses = append(contact.EmailAddresses, EmailAddress{
			DisplayName: emailAddress.DisplayName,
			Type:        emailAddress.Type,
			Value:       emailAddress.Value,
		})
	}

	for _, gender := range person.Genders {
		contact.Genders = append(contact.Genders, Gender{
			AddressMeAs: gender.AddressMeAs,
			Value:       gender.Value,
		})
	}
	for _, name := range person.Names {
		contact.Names = append(contact.Names, &UserNames{
			DisplayName:             name.DisplayName,
			DisplayNameLastFirst:    name.DisplayNameLastFirst,
			FamilyName:              name.FamilyName,
			GivenName:               name.GivenName,
			HonorificPrefix:         name.HonorificPrefix,
			HonorificSuffix:         name.HonorificSuffix,
			MiddleName:              name.MiddleName,
			PhoneticFamilyName:      name.PhoneticFamilyName,
			PhoneticFullName:        name.PhoneticFullName,
			PhoneticGivenName:       name.PhoneticGivenName,
			PhoneticHonorificPrefix: name.PhoneticHonorificSuffix,
			PhoneticHonorificSuffix: name.PhoneticHonorificSuffix,
			PhoneticMiddleName:      name.PhoneticMiddleName,
			UnstructuredName:        name.UnstructuredName,
		})
	}

	for _, nickname := range person.Nicknames {
		contact.Nicknames = append(contact.Nicknames, Nickname{
			Value: nickname.Value,
		})
	}

	for _, occupation := range person.Occupations {
		contact.Occupations = append(contact.Occupations, Occupation{
			Value: occupation.Value,
		})
	}

	for _, organization := range person.Organizations {
		contact.Organizations = append(contact.Organizations, Organization{
			Department:     organization.Department,
			Domain:         organization.Domain,
			EndDate:        getDate(organization.EndDate),
			JobDescription: organization.JobDescription,
			Location:       organization.Location,
			Name:           organization.Name,
			PhoneticName:   organization.PhoneticName,
			StartDate:      getDate(organization.StartDate),
			Symbol:         organization.Symbol,
			Title:          organization.Title,
			Type:           organization.Type,
			IsCurrent:      organization.Current,
		})
	}

	for _, phone := range person.PhoneNumbers {
		contact.PhoneNumbers = append(contact.PhoneNumbers, PhoneNumber{
			Type:  phone.Type,
			Value: phone.Value,
		})
	}

	for _, photo := range person.Photos {
		contact.Photos = append(contact.Photos, Photo{
			Default: photo.Default,
			Url:     photo.Url,
		})
	}

	for _, relation := range person.Relations {
		contact.Relations = append(contact.Relations, Relation{
			Person: relation.Person,
			Type:   relation.Type,
		})
	}

	for _, url := range person.Urls {
		contact.Urls = append(contact.Urls, Url{
			Type:  url.Type,
			Value: url.Value,
		})
	}
	contact.Metadata = person.Metadata // TODO: remove if not needed
	contact.ResourceName = person.ResourceName
	personId := strings.ReplaceAll(person.ResourceName, "people/", "")
	contact.ID = ContactID(personId)
}

func (c GoogleContact) IsContactDataUpdated(update GoogleContact) bool {

	if IsUpdated(c.Nicknames, update.Nicknames) {
		return true
	}

	if IsUpdated(c.Addresses, update.Addresses) {
		return true
	}

	if IsUpdated(c.Birthdays, update.Birthdays) {
		return true
	}

	if IsUpdated(c.EmailAddresses, update.EmailAddresses) {
		return true
	}

	if IsUpdated(c.Genders, update.Genders) {
		return true
	}

	if IsUpdated(c.Occupations, update.Occupations) {
		return true
	}

	if IsUpdated(c.PhoneNumbers, update.PhoneNumbers) {
		return true
	}

	if IsUpdated(c.Photos, update.Photos) {
		return true
	}

	if IsUpdated(c.Relations, update.Relations) {
		return true
	}

	if IsUpdated(c.Urls, update.Urls) {
		return true
	}

	if IsUpdated(c.Organizations, update.Organizations) {
		return true
	}

	return false
}

func (user GoogleContact) MapToDomain() (dto Contact) {
	dto = Contact{}
	dto.ID = ContactID(user.ID)
	dto.Addresses = append(dto.Addresses, user.Addresses...)
	dto.Birthdays = append(dto.Birthdays, user.Birthdays...)
	dto.EmailAddresses = append(dto.EmailAddresses, user.EmailAddresses...)
	dto.Genders = append(dto.Genders, user.Genders...)
	dto.Names = append(dto.Names, user.Names...)
	dto.Nicknames = append(dto.Nicknames, user.Nicknames...)
	dto.Occupations = append(dto.Occupations, user.Occupations...)
	dto.Organizations = append(dto.Organizations, user.Organizations...)
	dto.PhoneNumbers = append(dto.PhoneNumbers, user.PhoneNumbers...)
	dto.Photos = append(dto.Photos, user.Photos...)
	dto.Relations = append(dto.Relations, user.Relations...)
	dto.Urls = append(dto.Urls, user.Urls...)
	return
}
