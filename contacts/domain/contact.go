package domain

import (
	"time"

	"github.com/simplycubed/contactkarma/contacts/gen/models"
	people "google.golang.org/api/people/v1"
)

type ContactID string

func (c ContactID) String() string {
	return string(c)
}

type ContactCatgeory string

func (c ContactCatgeory) String() string {
	return string(c)
}

const A ContactCatgeory = "A"
const B ContactCatgeory = "B"
const C ContactCatgeory = "C"
const D ContactCatgeory = "D"

func (catgeory ContactCatgeory) Valid() bool {
	if catgeory == A || catgeory == B || catgeory == C || catgeory == D {
		return true
	}
	return true
}

type Contact struct {
	ID             ContactID      `firestore:"id" json:"id"`
	DisplayName    string         `firestore:"display_name" json:"display_name"`
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
}

func IsUpdated[T comparable](current []T, update []T) bool {
	if len(current) != len(update) {
		return true
	} else {
		// compare all indices
		for index, current := range current {
			if current != update[index] {
				return true
			}
		}
	}
	return false
}

func (c Contact) IsUpdated(update Contact) bool {

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

func (user Contact) MapToDto() (dto *models.Contact) {
	dto = &models.Contact{}
	dto.ID = string(user.ID)
	for _, address := range user.Addresses {
		dto.Addresses = append(dto.Addresses, address.MapToDto())
	}

	for _, birthday := range user.Birthdays {
		dto.Birthdays = append(dto.Birthdays, birthday.MapToDto())
	}

	for _, emailAddress := range user.EmailAddresses {
		dto.EmailAddresses = append(dto.EmailAddresses, emailAddress.MapToDto())
	}

	for _, gender := range user.Genders {
		dto.Genders = append(dto.Genders, gender.MapToDto())
	}

	for _, name := range user.Names {
		dto.Names = append(dto.Names, name.MapToDto())
	}

	for _, nickname := range user.Nicknames {
		dto.Nicknames = append(dto.Nicknames, nickname.MapToDto())
	}

	for _, occupation := range user.Occupations {
		dto.Occupations = append(dto.Occupations, occupation.MapToDto())
	}

	for _, organization := range user.Organizations {
		dto.Organizations = append(dto.Organizations, organization.MapToDto())
	}

	for _, phone := range user.PhoneNumbers {
		dto.PhoneNumbers = append(dto.PhoneNumbers, phone.MapToDto())
	}

	for _, photo := range user.Photos {
		dto.Photos = append(dto.Photos, photo.MapToDto())
	}

	for _, relation := range user.Relations {
		dto.Relations = append(dto.Relations, relation.MapToDto())
	}

	for _, url := range user.Urls {
		dto.Urls = append(dto.Urls, url.MapToDto())
	}

	return
}

func (contact *Contact) FromDto(dto *models.CreateContactDto) {

	for _, address := range dto.Addresses {
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

	for _, birthday := range dto.Birthdays {
		contact.Birthdays = append(contact.Birthdays, Birthday{
			Date: birthday.Date,
			Text: birthday.Text,
		})
	}

	for _, emailAddress := range dto.EmailAddresses {
		contact.EmailAddresses = append(contact.EmailAddresses, EmailAddress{
			DisplayName: emailAddress.DisplayName,
			Type:        emailAddress.Type,
			Value:       emailAddress.Value,
		})
	}

	for _, gender := range dto.Genders {
		contact.Genders = append(contact.Genders, Gender{
			AddressMeAs: gender.AddressMeAs,
			Value:       gender.Value,
		})
	}
	for _, name := range dto.Names {
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

	for _, nickname := range dto.Nicknames {
		contact.Nicknames = append(contact.Nicknames, Nickname{
			Value: nickname.Value,
		})
	}

	for _, occupation := range dto.Occupations {
		contact.Occupations = append(contact.Occupations, Occupation{
			Value: occupation.Value,
		})
	}

	for _, organization := range dto.Organizations {
		contact.Organizations = append(contact.Organizations, Organization{
			Department:     organization.Department,
			Domain:         organization.Domain,
			EndDate:        organization.EndDate,
			JobDescription: organization.JobDescription,
			Location:       organization.Location,
			Name:           organization.Name,
			PhoneticName:   organization.PhoneticName,
			StartDate:      organization.StartDate,
			Symbol:         organization.Symbol,
			Title:          organization.Title,
			Type:           organization.Type,
			IsCurrent:      organization.IsCurrent,
		})
	}

	for _, phone := range dto.PhoneNumbers {
		contact.PhoneNumbers = append(contact.PhoneNumbers, PhoneNumber{
			Type:  phone.Type,
			Value: phone.Value,
		})
	}

	for _, photo := range dto.Photos {
		contact.Photos = append(contact.Photos, Photo{
			Default: photo.Default,
			Url:     photo.URL,
		})
	}

	for _, relation := range dto.Relations {
		contact.Relations = append(contact.Relations, Relation{
			Person: relation.Person,
			Type:   relation.Type,
		})
	}

	for _, url := range dto.Urls {
		contact.Urls = append(contact.Urls, Url{
			Type:  url.Type,
			Value: url.Value,
		})
	}

	contact.DisplayName = dto.DisplayName

}

func (contact *Contact) FromUnified(dto Unified) {

	for _, address := range dto.Addresses {
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

	for _, birthday := range dto.Birthdays {
		contact.Birthdays = append(contact.Birthdays, Birthday{
			Date: birthday.Date,
			Text: birthday.Text,
		})
	}

	for _, emailAddress := range dto.EmailAddresses {
		contact.EmailAddresses = append(contact.EmailAddresses, EmailAddress{
			DisplayName: emailAddress.DisplayName,
			Type:        emailAddress.Type,
			Value:       emailAddress.Value,
		})
	}

	for _, gender := range dto.Genders {
		contact.Genders = append(contact.Genders, Gender{
			AddressMeAs: gender.AddressMeAs,
			Value:       gender.Value,
		})
	}
	for _, name := range dto.Names {
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

	for _, nickname := range dto.Nicknames {
		contact.Nicknames = append(contact.Nicknames, Nickname{
			Value: nickname.Value,
		})
	}

	for _, occupation := range dto.Occupations {
		contact.Occupations = append(contact.Occupations, Occupation{
			Value: occupation.Value,
		})
	}

	for _, organization := range dto.Organizations {
		contact.Organizations = append(contact.Organizations, Organization{
			Department:     organization.Department,
			Domain:         organization.Domain,
			EndDate:        organization.EndDate,
			JobDescription: organization.JobDescription,
			Location:       organization.Location,
			Name:           organization.Name,
			PhoneticName:   organization.PhoneticName,
			StartDate:      organization.StartDate,
			Symbol:         organization.Symbol,
			Title:          organization.Title,
			Type:           organization.Type,
			IsCurrent:      organization.IsCurrent,
		})
	}

	for _, phone := range dto.PhoneNumbers {
		contact.PhoneNumbers = append(contact.PhoneNumbers, PhoneNumber{
			Type:  phone.Type,
			Value: phone.Value,
		})
	}

	for _, photo := range dto.Photos {
		contact.Photos = append(contact.Photos, Photo{
			Default: photo.Default,
			Url:     photo.Url,
		})
	}

	for _, relation := range dto.Relations {
		contact.Relations = append(contact.Relations, Relation{
			Person: relation.Person,
			Type:   relation.Type,
		})
	}

	for _, url := range dto.Urls {
		contact.Urls = append(contact.Urls, Url{
			Type:  url.Type,
			Value: url.Value,
		})
	}

}

func getDate(date *people.Date) string {
	if date == nil {
		return ""
	}
	layout := "2006-01-02"
	return time.Date(int(date.Year), time.Month(date.Month), int(date.Day), 0, 0, 0, 0, time.UTC).Format(layout)
}

// dateStr must of the format 'MM-dd"
func PareseBirthdayDate(dateStr string) (date time.Time, err error) {
	layout := "01-02"
	return time.Parse(layout, dateStr)
}

func GetBirthdayNumber(date time.Time) int {
	return int(date.Month())*100 + date.Day()
}

// dateStr must of the format 'MM-dd"
func GetBirthdayNumberFromString(dateStr string) (number int, err error) {
	date, err := PareseBirthdayDate(dateStr)
	if err != nil {
		return
	}
	number = GetBirthdayNumber(date)
	return
}

type Contacts []Contact

func (contacts Contacts) MapToDto() (dto []*models.Contact) {
	for _, contact := range contacts {
		dto = append(dto, contact.MapToDto())
	}
	return
}

func (contacts Contacts) MapToIds() (ids []ContactID) {
	for _, contact := range contacts {
		ids = append(ids, contact.ID)
	}
	return
}

func (contact *Contact) FromCsv(dto CsvContact) {
	names := UserNames{
		DisplayName:             dto.DisplayName,
		DisplayNameLastFirst:    dto.DisplayNameLastFirst,
		FamilyName:              dto.FamilyName,
		GivenName:               dto.GivenName,
		HonorificPrefix:         dto.HonorificPrefix,
		HonorificSuffix:         dto.HonorificSuffix,
		MiddleName:              dto.MiddleName,
		PhoneticFamilyName:      dto.PhoneticFamilyName,
		PhoneticFullName:        dto.PhoneticFullName,
		PhoneticGivenName:       dto.PhoneticGivenName,
		PhoneticHonorificPrefix: dto.PhoneticHonorificPrefix,
		PhoneticHonorificSuffix: dto.PhoneticHonorificSuffix,
		PhoneticMiddleName:      dto.PhoneticMiddleName,
		UnstructuredName:        dto.UnstructuredName,
	}
	contact.Names = append(contact.Names, &names)

	address := Address{
		City:            dto.City,
		Country:         dto.Country,
		CountryCode:     dto.CountryCode,
		ExtendedAddress: dto.ExtendedAddress,
		PoBox:           dto.PoBox,
		PostalCode:      dto.PostalCode,
		Region:          dto.Region,
		StreetAddress:   dto.StreetAddress,
		Type:            dto.AddressType,
	}
	contact.Addresses = append(contact.Addresses, address)

	birthday := Birthday{
		Date: dto.BirthDate,
		Text: dto.BirthText,
	}

	contact.Birthdays = append(contact.Birthdays, birthday)

	emailAddress := EmailAddress{
		DisplayName: dto.EmailDisplayName,
		Type:        dto.EmailType,
		Value:       dto.Email,
	}
	contact.EmailAddresses = append(contact.EmailAddresses, emailAddress)

	gender := Gender{
		AddressMeAs: dto.AddressMeAs,
		Value:       dto.Gender,
	}
	contact.Genders = append(contact.Genders, gender)

	nickName := Nickname{
		Value: dto.Nickname,
	}
	contact.Nicknames = append(contact.Nicknames, nickName)

	occupation := Occupation{
		Value: dto.Occupation,
	}
	contact.Occupations = append(contact.Occupations, occupation)

	organization := Organization{
		Department:     dto.Department,
		Domain:         dto.Domain,
		EndDate:        dto.EndDate,
		JobDescription: dto.JobDescription,
		Location:       dto.Location,
		Name:           dto.Name,
		PhoneticName:   dto.PhoneticName,
		StartDate:      dto.StartDate,
		Symbol:         dto.Symbol,
		Title:          dto.Title,
		Type:           dto.OrganizationType,
		IsCurrent:      dto.IsCurrent == "true",
	}
	contact.Organizations = append(contact.Organizations, organization)

	phoneNumber := PhoneNumber{
		Type:  dto.PhoneType,
		Value: dto.Phone,
	}
	contact.PhoneNumbers = append(contact.PhoneNumbers, phoneNumber)

	photo := Photo{
		Default: dto.PhotoDefault == "true",
		Url:     dto.PhotoUrl,
	}
	contact.Photos = append(contact.Photos, photo)

	relation := Relation{
		Person: dto.Relation,
		Type:   dto.RelationType,
	}
	contact.Relations = append(contact.Relations, relation)

	url := Url{
		Type:  dto.UrlType,
		Value: dto.Url,
	}
	contact.Urls = append(contact.Urls, url)
}
