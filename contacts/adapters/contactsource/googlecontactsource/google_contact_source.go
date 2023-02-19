package googlecontactsource

import (
	"context"
	"strconv"
	"strings"

	"github.com/simplycubed/contactkarma/contacts/adapters/firestore"
	"github.com/simplycubed/contactkarma/contacts/application"
	"github.com/simplycubed/contactkarma/contacts/domain"
	"github.com/simplycubed/contactkarma/contacts/domain/repository"
	"google.golang.org/api/people/v1"
)

type GoogleContactSource struct {
	repo                 *firestore.GoogleContactsFirestore
	contactSourceRepo    *firestore.ContactSourceFirestore
	peopleServiceFactory application.PeopleServiceFactory
	googleContactsRepo   repository.IGoogleContacts
}

func NewGoogleContactSource(
	repo *firestore.GoogleContactsFirestore,
	contactSourceRepo *firestore.ContactSourceFirestore,
	peopleServiceFactory application.PeopleServiceFactory,
	googleContactsRepo repository.IGoogleContacts,
) *GoogleContactSource {
	return &GoogleContactSource{
		repo:                 repo,
		contactSourceRepo:    contactSourceRepo,
		peopleServiceFactory: peopleServiceFactory,
		googleContactsRepo:   googleContactsRepo,
	}
}

// Google person update: https://developers.google.com/people/api/rest/v1/people/updateContact?hl=en
func (source *GoogleContactSource) Update(ctx context.Context, userId domain.UserID, sourceId domain.ContactSourceID, updates []domain.ContactSourceUpdate) (err error) {

	// apply update to source, then firestore
	contactSource, err := source.contactSourceRepo.GetById(ctx, userId, sourceId)
	if err != nil {
		return
	}

	peopleService := source.peopleServiceFactory.New(ctx, contactSource.AccessToken, contactSource.RefreshToken, contactSource.TokenExpiry)

	personIds := []string{}
	for _, update := range updates {
		personIds = append(personIds, "people/"+string(update.ContactId))
	}
	batchGetResponse, err := peopleService.BatchGet(personIds)
	if err != nil {
		return
	}
	personMap := map[string]*people.Person{}
	for _, person := range batchGetResponse.Responses {
		personId := strings.ReplaceAll(person.Person.ResourceName, "people/", "")
		personMap[personId] = person.Person
	}

	updatesMap := map[string]people.Person{}
	for _, update := range updates {
		personBeforeUpdate := personMap[string(update.ContactId)]
		person := MapPerson(update.Unified)
		person.Etag = personBeforeUpdate.Etag
		updatesMap["people/"+string(update.ContactId)] = *person
	}
	response, err := peopleService.BatchUpdate(updatesMap)
	if err != nil {
		return
	}
	// TODO: bulk update repo
	for _, update := range response.UpdateResult {
		contact := domain.GoogleContact{}
		contact.FromGooglePerson(update.Person)
		err = source.repo.Update(ctx, userId, sourceId, contact.ID.String(), contact)
		if err != nil {
			return
		}
	}
	return
}

func (source *GoogleContactSource) Delete(ctx context.Context, userId domain.UserID, sourceId domain.ContactSourceID, contactIds []domain.ContactID) (err error) {
	// apply update to source, then firestore
	contactSource, err := source.contactSourceRepo.GetById(ctx, userId, sourceId)
	if err != nil {
		return
	}

	peopleService := source.peopleServiceFactory.New(ctx, contactSource.AccessToken, contactSource.RefreshToken, contactSource.TokenExpiry)
	personIds := []string{}
	for _, contactId := range contactIds {
		personIds = append(personIds, "people/"+contactId.String())
	}
	err = peopleService.BatchDelete(personIds)
	return
}

func (source *GoogleContactSource) Puller(ctx context.Context, userId domain.UserID, contactSource domain.ContactSource) (puller application.IContactSourcePuller) {
	peopleService := source.peopleServiceFactory.New(ctx, contactSource.AccessToken, contactSource.RefreshToken, contactSource.TokenExpiry)
	return NewGoogleContactSourcePuller(userId, contactSource.ID, source.googleContactsRepo, peopleService)
}

func (source *GoogleContactSource) Reader(ctx context.Context, userId domain.UserID, contactSourceId domain.ContactSourceID) (puller application.IContactSourceReader) {
	batchSize := 100 // TODO: move batch size to contract and set in env
	return NewGoogleContactsReader(source.googleContactsRepo, userId, contactSourceId, batchSize)
}

func (source *GoogleContactSource) Remove(ctx context.Context, userId domain.UserID, sourceId domain.ContactSourceID, contactIds []domain.ContactID) (err error) {
	return source.repo.BulkDeleteContacts(ctx, userId, sourceId, contactIds)
}

func MapAddress(address domain.Address) (dto *people.Address) {
	dto = &people.Address{
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

func getDate(dateString string) (date *people.Date) {
	if dateString != "" {
		return
	}

	dates := strings.Split(dateString, "-")
	if len(dates) < 3 {
		return
	}
	year, _ := strconv.Atoi(dates[0])
	month, _ := strconv.Atoi(dates[1])
	day, _ := strconv.Atoi(dates[2])
	date = &people.Date{
		Day:   int64(day),
		Month: int64(month),
		Year:  int64(year),
	}
	return
}

func MapBirthday(birthday domain.Birthday) (dto *people.Birthday) {
	dto = &people.Birthday{
		Date: getDate(birthday.Date),
		Text: birthday.Text,
	}
	return
}

func MapEmailAddress(email domain.EmailAddress) (dto *people.EmailAddress) {
	dto = &people.EmailAddress{
		DisplayName: email.DisplayName,
		Type:        email.Type,
		Value:       email.Value,
	}
	return
}

func MapGender(gender domain.Gender) (dto *people.Gender) {
	dto = &people.Gender{
		AddressMeAs: gender.AddressMeAs,
		Value:       gender.Value,
	}
	return
}

func MapName(name *domain.UserNames) (dto *people.Name) {
	dto = &people.Name{
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

func MapNickname(name domain.Nickname) (dto *people.Nickname) {
	dto = &people.Nickname{
		Value: name.Value,
	}
	return
}

func MapOccupation(occupation domain.Occupation) (dto *people.Occupation) {
	dto = &people.Occupation{
		Value: occupation.Value,
	}
	return
}

func MapOrganization(organization domain.Organization) (dto *people.Organization) {
	dto = &people.Organization{
		Department:     organization.Department,
		Domain:         organization.Domain,
		EndDate:        getDate(organization.EndDate),
		Current:        organization.IsCurrent,
		JobDescription: organization.JobDescription,
		Location:       organization.Location,
		Name:           organization.Name,
		PhoneticName:   organization.PhoneticName,
		StartDate:      getDate(organization.StartDate),
		Symbol:         organization.Symbol,
		Title:          organization.Title,
		Type:           organization.Type,
	}
	return
}

func MapPhone(phone domain.PhoneNumber) (dto *people.PhoneNumber) {
	dto = &people.PhoneNumber{
		Type:  phone.Type,
		Value: phone.Value,
	}
	return
}

func MapPhoto(phone domain.Photo) (dto *people.Photo) {
	dto = &people.Photo{
		Default: phone.Default,
		Url:     phone.Url,
	}
	return
}

func MapRelation(relation domain.Relation) (dto *people.Relation) {
	dto = &people.Relation{
		Person: relation.Person,
		Type:   relation.Type,
	}
	return
}

func MapUrl(url domain.Url) (dto *people.Url) {
	dto = &people.Url{
		Type:  url.Type,
		Value: url.Value,
	}
	return
}

func MapPerson(user domain.Unified) (dto *people.Person) {
	dto = &people.Person{}
	for _, address := range user.Addresses {
		dto.Addresses = append(dto.Addresses, MapAddress(address))
	}

	for _, birthday := range user.Birthdays {
		dto.Birthdays = append(dto.Birthdays, MapBirthday(birthday))
	}

	for _, emailAddress := range user.EmailAddresses {
		dto.EmailAddresses = append(dto.EmailAddresses, MapEmailAddress(emailAddress))
	}

	for _, gender := range user.Genders {
		dto.Genders = append(dto.Genders, MapGender(gender))
	}

	for _, name := range user.Names {
		dto.Names = append(dto.Names, MapName(name))
	}

	for _, nickname := range user.Nicknames {
		dto.Nicknames = append(dto.Nicknames, MapNickname(nickname))
	}

	for _, occupation := range user.Occupations {
		dto.Occupations = append(dto.Occupations, MapOccupation(occupation))
	}

	for _, organization := range user.Organizations {
		dto.Organizations = append(dto.Organizations, MapOrganization(organization))
	}

	for _, phone := range user.PhoneNumbers {
		dto.PhoneNumbers = append(dto.PhoneNumbers, MapPhone(phone))
	}

	for _, photo := range user.Photos {
		dto.Photos = append(dto.Photos, MapPhoto(photo))
	}

	for _, relation := range user.Relations {
		dto.Relations = append(dto.Relations, MapRelation(relation))
	}

	for _, url := range user.Urls {
		dto.Urls = append(dto.Urls, MapUrl(url))
	}

	return
}
