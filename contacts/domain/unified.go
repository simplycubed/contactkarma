package domain

import (
	"strings"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/simplycubed/contactkarma/contacts/gen/models"
)

type UnifiedId string

func (c UnifiedId) String() string {
	return string(c)
}

type UnifiedIds []UnifiedId

func (c *UnifiedIds) FromArray(ids []string) {
	unifiedIds := UnifiedIds{}
	for _, id := range ids {
		unifiedIds = append(unifiedIds, UnifiedId(id))
	}
	*c = unifiedIds
}

type ContactOrigin string

func (c ContactOrigin) String() string {
	return string(c)
}

func NewContactOrigin(source Source, contactSourceId ContactSourceID, contactId ContactID) ContactOrigin {
	return ContactOrigin(source.String() + ":" + contactSourceId.String() + ":" + contactId.String())
}

func (c ContactOrigin) Source() (source Source) {
	return Source(strings.Split(string(c), ":")[0])
}

func (c ContactOrigin) SourceId() (sourceId ContactSourceID) {
	return ContactSourceID(strings.Split(string(c), ":")[1])
}

func (c ContactOrigin) ContactId() (contactID ContactID) {
	return ContactID(strings.Split(string(c), ":")[2])
}

type Unified struct {
	ID          UnifiedId `firestore:"id" json:"id"`
	DisplayName string    `firestore:"display_name" json:"display_name"`

	Names          []*UserNames    `firestore:"names" json:"names"`
	Nicknames      []Nickname      `firestore:"nicknames" json:"nicknames" fakesize:"1"`
	Addresses      []Address       `firestore:"addresses" json:"addresses" fakesize:"1"`
	Birthdays      []Birthday      `firestore:"birthdays" json:"birthdays" fakesize:"1"`
	EmailAddresses []EmailAddress  `firestore:"email_addresses" json:"email_addresses" fakesize:"1"`
	Genders        []Gender        `firestore:"genders" json:"genders" fakesize:"1"`
	Occupations    []Occupation    `firestore:"occupations" json:"occupations" fakesize:"1"`
	PhoneNumbers   []PhoneNumber   `firestore:"phone_numbers" json:"phone_numbers" fakesize:"1"`
	Photos         []Photo         `firestore:"photos" json:"photos" fakesize:"1"`
	Relations      []Relation      `firestore:"relations" json:"relations" fakesize:"1"`
	Urls           []Url           `firestore:"urls" json:"urls" fakesize:"1"`
	Organizations  []Organization  `firestore:"organizations" json:"organizations" fakesize:"1"`
	CreatedAt      time.Time       `firestore:"created_at" json:"created_at"`
	UpdatedAt      time.Time       `firestore:"updated_at" json:"updated_at"`
	NextContact    *time.Time      `firestore:"next_contact" json:"next_contact"`
	LastContact    *time.Time      `firestore:"last_contact" json:"last_contact"`
	Score          int             `firestore:"score" json:"score"`
	Category       ContactCatgeory `firestore:"category" json:"category"`

	// Fields unique to unified
	// contact sources from which this unified contact is created/linked
	Origins     map[string]bool `firestore:"origins" json:"origins"`
	SearchTerms []string        `firestore:"search_terms" json:"search_terms"`
}

func (name *UserNames) Merge(duplicate *UserNames) {
	if name.DisplayName == "" {
		name.DisplayName = duplicate.DisplayName
	}
	if name.DisplayNameLastFirst == "" {
		name.DisplayNameLastFirst = duplicate.DisplayNameLastFirst
	}

	if name.FamilyName == "" {
		name.FamilyName = duplicate.FamilyName
	}

	if name.GivenName == "" {
		name.GivenName = duplicate.GivenName
	}

	if name.HonorificPrefix == "" {
		name.HonorificPrefix = duplicate.HonorificPrefix
	}

	if name.HonorificSuffix == "" {
		name.HonorificSuffix = duplicate.HonorificSuffix
	}

	if name.MiddleName == "" {
		name.MiddleName = duplicate.MiddleName
	}

	if name.PhoneticFamilyName == "" {
		name.PhoneticFamilyName = duplicate.PhoneticFamilyName
	}

	if name.PhoneticFullName == "" {
		name.PhoneticFullName = duplicate.PhoneticFullName
	}

	if name.PhoneticGivenName == "" {
		name.PhoneticFullName = duplicate.PhoneticGivenName
	}
	if name.PhoneticHonorificPrefix == "" {
		name.PhoneticFullName = duplicate.PhoneticHonorificPrefix
	}

	if name.PhoneticHonorificSuffix == "" {
		name.PhoneticHonorificSuffix = duplicate.PhoneticHonorificSuffix
	}

	if name.PhoneticMiddleName == "" {
		name.PhoneticMiddleName = duplicate.PhoneticMiddleName
	}

	if name.UnstructuredName == "" {
		name.UnstructuredName = duplicate.UnstructuredName
	}
}

func (u *Unified) MergeNicknames(duplicate []Nickname) {
	valueMap := map[string]bool{}
	for _, name := range u.Nicknames {
		valueMap[name.Value] = true
	}
	for _, name := range duplicate {
		_, ok := valueMap[name.Value]
		if !ok {
			u.Nicknames = append(u.Nicknames, name)
		}
	}
}

func (u *Unified) MergeBirthDays(duplicate []Birthday) {
	valueMap := map[string]bool{}
	for _, birthday := range u.Birthdays {
		valueMap[birthday.Date] = true
	}
	for _, birthday := range duplicate {
		_, ok := valueMap[birthday.Date]
		if !ok {
			u.Birthdays = append(u.Birthdays, birthday)
		}
	}
}

func (u *Unified) MergeAddresses(duplicate []Address) {
	for _, address := range duplicate {
		// check if address already added
		isPresent := false
		for _, toCheck := range u.Addresses {
			if address == toCheck {
				isPresent = true
				break
			}
		}
		if !isPresent {
			u.Addresses = append(u.Addresses, address)
		}
	}
}

func (u *Unified) MergeOccupations(duplicate []Occupation) {
	for _, occupation := range duplicate {
		// check if address already added
		isPresent := false
		for _, toCheck := range u.Occupations {
			if occupation == toCheck {
				isPresent = true
				break
			}
		}
		if !isPresent {
			u.Occupations = append(u.Occupations, occupation)
		}
	}
}

func (u *Unified) MergeGenders(duplicate []Gender) {
	valueMap := map[string]bool{}
	for _, gender := range u.Genders {
		valueMap[gender.Value] = true
	}
	for _, gender := range duplicate {
		_, ok := valueMap[gender.Value]
		if !ok {
			u.Genders = append(u.Genders, gender)
		}
	}
}

func (u *Unified) MergeEmails(duplicate []EmailAddress) {
	valueMap := map[string]bool{}
	for _, email := range u.EmailAddresses {
		valueMap[email.Value] = true
	}

	for _, email := range duplicate {
		_, ok := valueMap[email.Value]
		if !ok {
			u.EmailAddresses = append(u.EmailAddresses, email)
		}
	}
}

func (u *Unified) MergePhones(duplicate []PhoneNumber) {
	valueMap := map[string]bool{}
	for _, phone := range u.PhoneNumbers {
		valueMap[phone.Value] = true
	}

	for _, phone := range duplicate {
		_, ok := valueMap[phone.Value]
		if !ok {
			u.PhoneNumbers = append(u.PhoneNumbers, phone)
		}
	}

}

func (u *Unified) MergePhotos(duplicate []Photo) {
	valueMap := map[string]bool{}
	for _, photo := range u.Photos {
		valueMap[photo.Url] = true
	}

	for _, photo := range duplicate {
		_, ok := valueMap[photo.Url]
		if !ok {
			u.Photos = append(u.Photos, photo)
		}
	}

}

func (u *Unified) MergeRelations(duplicate []Relation) {
	valueMap := map[string]bool{}
	for _, relation := range u.Relations {
		valueMap[relation.Person] = true
	}

	for _, relation := range duplicate {
		_, ok := valueMap[relation.Person]
		if !ok {
			u.Relations = append(u.Relations, relation)
		}
	}

}

func (u *Unified) MergeUrls(duplicate []Url) {
	valueMap := map[string]bool{}
	for _, url := range u.Urls {
		valueMap[url.Value] = true
	}

	for _, url := range duplicate {
		_, ok := valueMap[url.Value]
		if !ok {
			u.Urls = append(u.Urls, url)
		}
	}

}

func (u *Unified) MergeOrganizations(duplicate []Organization) {
	for _, organization := range duplicate {
		// check if address already added
		isPresent := false
		for _, toCheck := range u.Organizations {
			if organization == toCheck {
				isPresent = true
				break
			}
		}
		if !isPresent {
			u.Organizations = append(u.Organizations, organization)
		}
	}
}

// Merge() merges values from another unified contact.

func (u *Unified) Merge(duplicate Unified) {
	if len(u.Names) == 0 {
		u.Names = duplicate.Names
	} else if len(duplicate.Names) > 0 {
		name := u.Names[0]
		name.Merge(duplicate.Names[0])
	}
	u.MergeNicknames(duplicate.Nicknames)
	u.MergeAddresses(duplicate.Addresses)
	u.MergeBirthDays(duplicate.Birthdays)
	u.MergeEmails(duplicate.EmailAddresses)
	u.MergeGenders(duplicate.Genders)
	u.MergeOccupations(duplicate.Occupations)
	u.MergePhones(duplicate.PhoneNumbers)

	u.MergePhotos(duplicate.Photos)
	u.MergeRelations(duplicate.Relations)
	u.MergeUrls(duplicate.Urls)
	u.MergeOrganizations(duplicate.Organizations)

	if u.NextContact == nil {
		u.NextContact = duplicate.NextContact
	}
	if u.LastContact == nil {
		u.LastContact = duplicate.LastContact
	}
	if u.Score < duplicate.Score {
		u.Score = duplicate.Score
	}
	if u.Category == "" {
		u.Category = duplicate.Category
	}

	for origin := range duplicate.Origins {
		u.Origins[origin] = true
	}
	// reset search terms
	u.SearchTerms = u.ComputeSearchTerms()
}

func (u Unified) GetDisplayName() string {
	email := ""
	name := ""
	phone := ""
	if len(u.EmailAddresses) > 0 {
		email = u.EmailAddresses[0].Value
	}
	if len(u.Names) > 0 {
		name = u.Names[0].GivenName + " " + u.Names[0].MiddleName + " " + u.Names[0].FamilyName
	}
	if len(u.PhoneNumbers) > 0 {
		phone = u.PhoneNumbers[0].Value
	}
	displayName := name
	if phone != "" {
		displayName += " (" + phone + ")"
	}
	if email != "" {
		displayName += " (" + email + ")"
	}
	return displayName
}

func (user Unified) MapToDto() (dto *models.Unified) {
	dto = &models.Unified{}
	dto.ID = string(user.ID)
	dto.DisplayName = user.DisplayName
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

	if user.NextContact != nil {
		dto.NextContact = strfmt.DateTime(*user.NextContact)
	}

	if user.LastContact != nil {
		dto.LastContact = strfmt.DateTime(*user.LastContact)
	}

	dto.Category = string(user.Category)
	dto.Score = int64(user.Score)
	return
}

// TODO: type of param dto will be renamed to *models.CreateUserParams
func (contact *Unified) FromDto(dto *models.Unified) {

	contact.ID = UnifiedId(dto.ID)
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

	if !time.Time(dto.NextContact).IsZero() {
		contact.NextContact = (*time.Time)(&dto.NextContact)
	}

	if !time.Time(dto.LastContact).IsZero() {
		contact.LastContact = (*time.Time)(&dto.LastContact)
	}

	contact.Category = ContactCatgeory(dto.Category)
}

func (contact *Unified) FromUpdateDto(dto *models.UpdateUnifiedDto) {
	if len(dto.Addresses) > 0 {
		contact.Addresses = []Address{}
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
	}

	if len(dto.Birthdays) > 0 {
		contact.Birthdays = []Birthday{}
		for _, birthday := range dto.Birthdays {
			contact.Birthdays = append(contact.Birthdays, Birthday{
				Date: birthday.Date,
				Text: birthday.Text,
			})
		}
	}
	if len(dto.EmailAddresses) > 0 {
		contact.EmailAddresses = []EmailAddress{}
		for _, emailAddress := range dto.EmailAddresses {
			contact.EmailAddresses = append(contact.EmailAddresses, EmailAddress{
				DisplayName: emailAddress.DisplayName,
				Type:        emailAddress.Type,
				Value:       emailAddress.Value,
			})
		}
	}
	if len(dto.Genders) > 0 {
		contact.Genders = []Gender{}
		for _, gender := range dto.Genders {
			contact.Genders = append(contact.Genders, Gender{
				AddressMeAs: gender.AddressMeAs,
				Value:       gender.Value,
			})
		}
	}
	if len(dto.Names) > 0 {
		contact.Names = []*UserNames{}
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
	}
	if len(dto.Nicknames) > 0 {
		contact.Nicknames = []Nickname{}
		for _, nickname := range dto.Nicknames {
			contact.Nicknames = append(contact.Nicknames, Nickname{
				Value: nickname.Value,
			})
		}
	}

	if len(dto.Occupations) > 0 {
		contact.Occupations = []Occupation{}
		for _, occupation := range dto.Occupations {
			contact.Occupations = append(contact.Occupations, Occupation{
				Value: occupation.Value,
			})
		}
	}

	if len(dto.Organizations) > 0 {
		contact.Organizations = []Organization{}
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
	}

	if len(dto.PhoneNumbers) > 0 {
		contact.PhoneNumbers = []PhoneNumber{}
		for _, phone := range dto.PhoneNumbers {
			contact.PhoneNumbers = append(contact.PhoneNumbers, PhoneNumber{
				Type:  phone.Type,
				Value: phone.Value,
			})
		}
	}

	if len(dto.Photos) > 0 {
		contact.Photos = []Photo{}
		for _, photo := range dto.Photos {
			contact.Photos = append(contact.Photos, Photo{
				Default: photo.Default,
				Url:     photo.URL,
			})
		}
	}

	if len(dto.Relations) > 0 {
		contact.Relations = []Relation{}
		for _, relation := range dto.Relations {
			contact.Relations = append(contact.Relations, Relation{
				Person: relation.Person,
				Type:   relation.Type,
			})
		}
	}

	if len(contact.Urls) > 0 {
		contact.Urls = []Url{}
		for _, url := range dto.Urls {
			contact.Urls = append(contact.Urls, Url{
				Type:  url.Type,
				Value: url.Value,
			})
		}
	}

	if !time.Time(dto.NextContact).IsZero() {
		contact.NextContact = (*time.Time)(&dto.NextContact)
	}

	if !time.Time(dto.LastContact).IsZero() {
		contact.LastContact = (*time.Time)(&dto.LastContact)
	}

	if dto.DisplayName != "" {
		contact.DisplayName = dto.DisplayName
	}
	contact.SearchTerms = contact.ComputeSearchTerms()

}

// https://simplycubed.slack.com/archives/C03404BM7CK/p1661501428468729
func (contact *Unified) ComputeDisplayName() string {
	if len(contact.Names) > 0 {
		userName := contact.Names[0]
		if userName != nil {
			if userName.GivenName != "" && userName.FamilyName != "" {
				return userName.GivenName + " " + userName.FamilyName
			} else if userName.GivenName != "" {
				return userName.GivenName
			} else if userName.FamilyName != "" {
				return userName.FamilyName
			}
		}
	}
	if len(contact.Organizations) > 0 {
		organization := contact.Organizations[0]
		if organization.Name != "" {
			return organization.Name
		}
	}
	if len(contact.EmailAddresses) > 0 {
		email := contact.EmailAddresses[0]
		if email.Value != "" {
			return email.Value
		}
	}
	// return any non empty value
	if len(contact.PhoneNumbers) > 0 {
		phone := contact.PhoneNumbers[0]
		if phone.Value != "" {
			return phone.Value
		}
	}
	// TODO: add check for all fields
	return ""
}

func (contact Unified) ComputeSearchTerms() (terms []string) {
	searchTerms := map[string]bool{}
	for _, emailAddress := range contact.EmailAddresses {
		searchTerms[strings.ToLower(emailAddress.Value)] = true
	}
	for _, phone := range contact.PhoneNumbers {
		searchTerms[strings.ToLower(phone.Value)] = true
	}
	for _, name := range contact.Names {
		searchTerms[strings.ToLower(name.GivenName+" "+name.FamilyName)] = true
	}
	terms = []string{}
	for term := range searchTerms {
		if term != "" {
			terms = append(terms, term)
		}
	}
	return
}

func (contact *Unified) FromContact(dto Contact, sourceId ContactOrigin) {

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

	contact.Origins = map[string]bool{}
	contact.Origins[sourceId.String()] = true

	contact.SearchTerms = contact.ComputeSearchTerms()
	if contact.DisplayName == "" {
		contact.DisplayName = contact.ComputeDisplayName()
	}

}

func (contact *Unified) SetUpdates(dto Contact) {
	contact.Addresses = dto.Addresses
	contact.Birthdays = dto.Birthdays
	contact.EmailAddresses = dto.EmailAddresses
	contact.Genders = dto.Genders
	contact.Names = dto.Names
	contact.Nicknames = dto.Nicknames
	contact.Occupations = dto.Occupations
	contact.Organizations = dto.Organizations
	contact.PhoneNumbers = dto.PhoneNumbers
	contact.Photos = dto.Photos
	contact.Relations = dto.Relations
	contact.Urls = dto.Urls
	contact.SearchTerms = contact.ComputeSearchTerms()
	if contact.DisplayName == "" {
		contact.DisplayName = contact.ComputeDisplayName()
	}
}

type UnifiedContacts []Unified

func (contacts UnifiedContacts) MapToDto() (dto []*models.Unified) {
	for _, contact := range contacts {
		dto = append(dto, contact.MapToDto())
	}
	return
}

func (contacts UnifiedContacts) FilterByEmail(emailFilter string) (filtered []Unified) {
	for _, contact := range contacts {
		for _, email := range contact.EmailAddresses {
			if strings.EqualFold(email.Value, emailFilter) {
				filtered = append(filtered, contact)
				break
			}
		}
	}
	return
}

func (contacts UnifiedContacts) FilterByPhone(phoneFilter string) (filtered []Unified) {
	for _, contact := range contacts {
		for _, phone := range contact.PhoneNumbers {
			if strings.EqualFold(phone.Value, phoneFilter) {
				filtered = append(filtered, contact)
				break
			}
		}
	}
	return
}

func (contacts UnifiedContacts) FilterByGivenName(nameFilter string) (filtered []Unified) {
	for _, contact := range contacts {
		for _, name := range contact.Names {
			if strings.EqualFold(name.GivenName, nameFilter) {
				filtered = append(filtered, contact)
				break
			}
		}
	}
	return
}

func (contacts UnifiedContacts) FilterByMiddleName(nameFilter string) (filtered []Unified) {
	for _, contact := range contacts {
		for _, name := range contact.Names {
			if strings.EqualFold(name.MiddleName, nameFilter) {
				filtered = append(filtered, contact)
				break
			}
		}
	}
	return
}

func (contacts UnifiedContacts) FilterByFamilyName(nameFilter string) (filtered []Unified) {
	for _, contact := range contacts {
		for _, name := range contact.Names {
			if strings.EqualFold(name.FamilyName, nameFilter) {
				filtered = append(filtered, contact)
				break
			}
		}
	}
	return
}

type UnifiedSearchId string

type SortOrder string

var ASC SortOrder = "asc"
var Desc SortOrder = "desc"

type Sort struct {
	Field string
	Order SortOrder
}

type SearchFilterField string

func (c SearchFilterField) String() string {
	return string(c)
}

var CategoryField SearchFilterField = "category"
var ContactField SearchFilterField = "next_contact"
var LastContactField SearchFilterField = "last_contact"
var ScoreField SearchFilterField = "score"
var BirthdayField SearchFilterField = "birthday"
var GenderField SearchFilterField = "genders"

type SearchOperator string

func (c SearchOperator) String() string {
	return string(c)
}

var Equal SearchOperator = "="
var NotEqual SearchOperator = "!="
var GreaterThan SearchOperator = ">"
var LessThan SearchOperator = "<"
var GreaterThanOrEqual = ">="
var LessThanOrEqual = "<="

type ContactSearchFilter struct {
	Field    SearchFilterField
	Operator SearchOperator
	Value    string
}

func MapSearchFilters(filters []*models.SearchFilter) (mapped []ContactSearchFilter) {
	for _, filter := range filters {
		if filter.Field == "gender" {
			mapped = append(mapped, ContactSearchFilter{
				Operator: SearchOperator(filter.Operator),
				Field:    GenderField,
				Value:    filter.Value,
			})
		} else {
			mapped = append(mapped, ContactSearchFilter{
				Operator: SearchOperator(filter.Operator),
				Field:    SearchFilterField(filter.Field),
				Value:    filter.Value,
			})
		}

	}
	return
}

func MapSearchSorts(sorts []*models.SearchSort) (mapped []Sort) {
	for _, sort := range sorts {
		mapped = append(mapped, Sort{
			Field: sort.Field,
			Order: SortOrder(sort.Order),
		})
	}
	return
}

type SearchParams struct {
	Filters []ContactSearchFilter
	Query   string
	Sort    []Sort
	Page    int
	PerPage int
}

func (params *SearchParams) FromModel(model *models.SearchContactDto) {
	params.Query = model.Query
	params.Sort = MapSearchSorts(model.Sort)
	params.Filters = MapSearchFilters(model.Filters)
	if model.Page > 0 {
		params.Page = int(model.Page)
	}
	if model.PerPage > 0 {
		params.PerPage = int(model.PerPage)
	}
	return
}

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
	BirthdayNumber *int            `json:"birthday_number"`
}

type UnifiedSearchResults []UnifiedSearch

func (results UnifiedSearchResults) MapToUnifiedIds() (ids []UnifiedId) {
	for _, result := range results {
		ids = append(ids, result.UnifiedId)
	}
	return
}
