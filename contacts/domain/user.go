package domain

import (
	"time"

	"github.com/simplycubed/contactkarma/contacts/gen/models"
)

type UserID string

type User struct {
	ID             UserID         `firestore:"id" json:"id"`
	Names          []UserNames    `firestore:"names" json:"names"`
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

	Quota *Quota `firestore:"quota" json:"quota"`
}

func (u *User) InitQuota() {
	u.Quota = &Quota{TotalContacts: 0, TotalCategoryAssigned: 0}
}

func (u *User) AddToTotalContacts(count int64) {
	if u.Quota == nil {
		u.InitQuota()
	}
	u.Quota.TotalContacts += count
}

func (u *User) AddToTotalCategoryAssigned(count int64) {
	if u.Quota == nil {
		u.InitQuota()
	}
	u.Quota.TotalContacts += count
}

func (u *User) AddToTotalContactSources(count int64) {
	if u.Quota == nil {
		u.InitQuota()
	}
	u.Quota.TotalContactSources += count
}

type UserNames struct {
	DisplayName             string `firestore:"display_name,omitempty" json:"display_name,omitempty" fake:"{firstname}"`
	DisplayNameLastFirst    string `firestore:"display_name_last_first,omitempty" json:"display_name_last_first,omitempty" fake:"{firstname} {lastname}"`
	FamilyName              string `firestore:"family_name,omitempty" json:"family_name,omitempty" fake:"{lastname}"`
	GivenName               string `firestore:"given_name,omitempty" json:"given_name,omitempty" fake:"{lastname}"`
	HonorificPrefix         string `firestore:"honorific_prefix,omitempty" json:"honorific_prefix,omitempty"`
	HonorificSuffix         string `firestore:"honorific_suffix,omitempty" json:"honorific_suffix,omitempty"`
	MiddleName              string `firestore:"middle_name,omitempty" json:"middle_name,omitempty" fake:"{lastname}"`
	PhoneticFamilyName      string `firestore:"phonetic_family_name,omitempty" json:"phonetic_family_name,omitempty" fake:"{lastname}"`
	PhoneticFullName        string `firestore:"phonetic_full_name,omitempty" json:"phonetic_full_name,omitempty" fake:"{firstname} {lastname}"`
	PhoneticGivenName       string `firestore:"phonetic_given_name,omitempty" json:"phonetic_given_name,omitempty" fake:"{firstname}"`
	PhoneticHonorificPrefix string `firestore:"phonetic_honorific_prefix,omitempty" json:"phonetic_honorific_prefix,omitempty"`
	PhoneticHonorificSuffix string `firestore:"phonetic_honorific_suffix,omitempty" json:"phonetic_honorific_suffix,omitempty"`
	PhoneticMiddleName      string `firestore:"phonetic_middle_name,omitempty" json:"phonetic_middle_name,omitempty" fake:"{lastname}"`
	UnstructuredName        string `firestore:"unstructured_name,omitempty" json:"unstructured_name,omitempty" fake:"{firstname} {lastname}"`
}

type Nickname struct {
	Value string `firestore:"value,omitempty" json:"value,omitempty" fake:"{firstname}"`
}

type Address struct {
	City            string `firestore:"city,omitempty" json:"city,omitempty" fake:"{city}"`
	Country         string `firestore:"country,omitempty" json:"country,omitempty" fake:"{country}"`
	CountryCode     string `firestore:"country_code,omitempty" json:"country_code,omitempty" fake:"{countryabr}"`
	ExtendedAddress string `firestore:"extended_address,omitempty" json:"extended_address,omitempty" fake:"{street}, {streetnumber}, {city}"`
	PoBox           string `firestore:"po_box,omitempty" json:"po_box,omitempty"`
	PostalCode      string `firestore:"postal_code,omitempty" json:"postal_code,omitempty" fake:"{zip}"`
	Region          string `firestore:"region,omitempty" json:"region,omitempty" fake:"{state}"`
	StreetAddress   string `firestore:"street_address,omitempty" json:"street_address,omitempty" fake:"{street}"`
	Type            string `firestore:"type,omitempty" json:"type,omitempty"`
}

type Birthday struct {
	Date string `firestore:"date,omitempty" json:"date,omitempty" fake:"{year}-{month}-{day}"`
	Text string `firestore:"text,omitempty" json:"text,omitempty" fake:"{year}-{month}-{day}"`
}

type EmailAddress struct {
	DisplayName string `firestore:"display_name,omitempty" json:"display_name,omitempty" fake:"{firstname}"`
	Type        string `firestore:"type,omitempty" json:"type,omitempty"`
	Value       string `firestore:"value,omitempty" json:"value,omitempty" fake:"{email}"`
}

type Gender struct {
	AddressMeAs string `firestore:"address_me_as,omitempty" json:"address_me_as,omitempty" fake:"{pronounpersonal}"`
	Value       string `firestore:"value,omitempty" json:"value,omitempty" fake:"{gender}"`
}

type Occupation struct {
	Value string `firestore:"value,omitempty" json:"value,omitempty" fake:"{jobtitle}"`
}

type Organization struct {
	Department     string `firestore:"department,omitempty" json:"department,omitempty"`
	Domain         string `firestore:"domain,omitempty" json:"domain,omitempty"`
	EndDate        string `firestore:"end_date,omitempty" json:"end_date,omitempty" fake:"{year}-{month}-{day}"`
	JobDescription string `firestore:"job_description,omitempty" json:"job_description,omitempty" fake:"{sentence:3}"`
	Location       string `firestore:"location,omitempty" json:"location,omitempty"`
	Name           string `firestore:"name,omitempty" json:"name,omitempty" fake:"{company}"`
	PhoneticName   string `firestore:"phonetic_name,omitempty" json:"phonetic_name,omitempty" fake:"{company}"`
	StartDate      string `firestore:"start_date,omitempty" json:"start_date,omitempty" fake:"{year}-{month}-{day}"`
	Symbol         string `firestore:"symbol,omitempty" json:"symbol,omitempty"`
	Title          string `firestore:"title,omitempty" json:"title,omitempty" fake:"{jobtitle}"`
	Type           string `firestore:"type,omitempty" json:"type,omitempty"`
	IsCurrent      bool   `firestore:"is_current,omitempty" json:"is_current,omitempty"`
}

type PhoneNumber struct {
	Type  string `firestore:"type,omitempty" json:"type,omitempty"`
	Value string `firestore:"value,omitempty" json:"value,omitempty" fake:"{phone}"`
}

type Photo struct {
	Default bool   `firestore:"default,omitempty" json:"default,omitempty"`
	Url     string `firestore:"url,omitempty" json:"url,omitempty" fake:"{url}"`
}

type Relation struct {
	Person string `firestore:"person,omitempty" json:"person,omitempty"`
	Type   string `firestore:"type,omitempty" json:"type,omitempty"`
}

type Url struct {
	Type  string `firestore:"type,omitempty" json:"type,omitempty"`
	Value string `firestore:"value,omitempty" json:"value,omitempty" fake:"{url}"`
}

func (c UserID) String() string {
	return string(c)
}

func (address Address) MapToDto() (dto *models.Address) {
	dto = &models.Address{
		City:            address.City,
		Country:         address.Country,
		CountryCode:     address.CountryCode,
		ExtendedAddress: address.ExtendedAddress,
		PoBox:           address.PoBox,
		PostalCode:      address.PostalCode,
		Region:          address.Region,
		StreetAddress:   address.StreetAddress,
		Type:            address.Type,
	}
	return
}

func (birthday Birthday) MapToDto() (dto *models.Birthday) {
	dto = &models.Birthday{
		Date: birthday.Date,
		Text: "", // ?
	}
	return
}

func (email EmailAddress) MapToDto() (dto *models.EmailAddress) {
	dto = &models.EmailAddress{
		DisplayName: email.DisplayName,
		Type:        email.Type,
		Value:       email.Value,
	}
	return
}

func (gender Gender) MapToDto() (dto *models.Gender) {
	dto = &models.Gender{
		AddressMeAs: gender.AddressMeAs,
		Value:       gender.Value,
	}
	return
}

func (name UserNames) MapToDto() (dto *models.UserNames) {
	dto = &models.UserNames{
		DisplayName:             name.DisplayName,
		DisplayNameLastFirst:    name.DisplayNameLastFirst,
		FamilyName:              name.FamilyName,
		GivenName:               name.GivenName,
		HonorificPrefix:         name.HonorificPrefix,
		HonorificSuffix:         name.HonorificPrefix,
		MiddleName:              name.MiddleName,
		PhoneticFamilyName:      name.PhoneticFamilyName,
		PhoneticFullName:        name.PhoneticFullName,
		PhoneticGivenName:       name.PhoneticGivenName,
		PhoneticHonorificPrefix: name.PhoneticHonorificPrefix,
		PhoneticHonorificSuffix: name.PhoneticHonorificSuffix,
		PhoneticMiddleName:      name.PhoneticMiddleName,
		UnstructuredName:        name.UnstructuredName,
	}
	return
}

func (name Nickname) MapToDto() (dto *models.Nickname) {
	dto = &models.Nickname{
		Value: name.Value,
	}
	return
}

func (occupation Occupation) MapToDto() (dto *models.Occupation) {
	dto = &models.Occupation{
		Value: occupation.Value,
	}
	return
}

func (phone PhoneNumber) MapToDto() (dto *models.PhoneNumber) {
	dto = &models.PhoneNumber{
		Type:  phone.Type,
		Value: phone.Value,
	}
	return
}

func (phone Photo) MapToDto() (dto *models.Photo) {
	dto = &models.Photo{
		Default: phone.Default,
		URL:     phone.Url,
	}
	return
}

func (relation Relation) MapToDto() (dto *models.Relation) {
	dto = &models.Relation{
		Person: relation.Person,
		Type:   relation.Type,
	}
	return
}

func (url Url) MapToDto() (dto *models.URL) {
	dto = &models.URL{
		Type:  url.Type,
		Value: url.Value,
	}
	return
}

func (organization Organization) MapToDto() (dto *models.Organization) {
	dto = &models.Organization{
		Department:     organization.Department,
		Domain:         organization.Domain,
		EndDate:        organization.EndDate,
		IsCurrent:      organization.IsCurrent,
		JobDescription: organization.JobDescription,
		Location:       organization.Location,
		Name:           organization.Name,
		PhoneticName:   organization.PhoneticName,
		StartDate:      organization.StartDate,
		Symbol:         organization.Symbol,
		Title:          organization.Title,
		Type:           organization.Type,
	}
	return
}

func (user User) MapToDto() (dto *models.User) {
	dto = &models.User{}
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
	if user.Quota != nil {
		dto.Quota = user.Quota.MapToDto()
	}
	return
}

// TODO: type of param dto will be renamed to *models.CreateUserParams
func (user *User) FromDto(dto *models.User) {

	user.ID = UserID(dto.ID)
	for _, address := range dto.Addresses {
		user.Addresses = append(user.Addresses, Address{
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
		user.Birthdays = append(user.Birthdays, Birthday{
			Date: birthday.Date,
			Text: birthday.Text,
		})
	}

	for _, emailAddress := range dto.EmailAddresses {
		user.EmailAddresses = append(user.EmailAddresses, EmailAddress{
			DisplayName: emailAddress.DisplayName,
			Type:        emailAddress.Type,
			Value:       emailAddress.Value,
		})
	}

	for _, gender := range dto.Genders {
		user.Genders = append(user.Genders, Gender{
			AddressMeAs: gender.AddressMeAs,
			Value:       gender.Value,
		})
	}

	for _, name := range dto.Names {
		user.Names = append(user.Names, UserNames{
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
		user.Nicknames = append(user.Nicknames, Nickname{
			Value: nickname.Value,
		})
	}

	for _, occupation := range dto.Occupations {
		user.Occupations = append(user.Occupations, Occupation{
			Value: occupation.Value,
		})
	}

	for _, organization := range dto.Organizations {
		user.Organizations = append(user.Organizations, Organization{
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
		user.PhoneNumbers = append(user.PhoneNumbers, PhoneNumber{
			Type:  phone.Type,
			Value: phone.Value,
		})
	}

	for _, photo := range dto.Photos {
		user.Photos = append(user.Photos, Photo{
			Default: photo.Default,
			Url:     photo.URL,
		})
	}

	for _, relation := range dto.Relations {
		user.Relations = append(user.Relations, Relation{
			Person: relation.Person,
			Type:   relation.Type,
		})
	}

	for _, url := range dto.Urls {
		user.Urls = append(user.Urls, Url{
			Type:  url.Type,
			Value: url.Value,
		})
	}

}

type Users []User

func (users Users) MapToDto() (dto []*models.User) {
	for _, contact := range users {
		dto = append(dto, contact.MapToDto())
	}
	return
}
