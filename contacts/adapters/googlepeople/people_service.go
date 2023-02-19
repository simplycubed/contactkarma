package googlepeople

import (
	"context"

	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/api/people/v1"
)

type PeopleService struct {
	peopleService *people.Service
}

func NewPeopleService(ctx context.Context, config *oauth2.Config, token *oauth2.Token) (service *PeopleService) {
	peopleService, err := people.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
	if err != nil {
		return
	}
	service = &PeopleService{peopleService: peopleService}
	return
}

func (s *PeopleService) personFields() string {
	return "addresses,birthdays,emailAddresses,genders,interests,locations,names,nicknames,occupations,organizations,phoneNumbers,photos,relations,skills,urls,metadata"
}

func (s *PeopleService) updateFields() string {
	return "addresses,birthdays,emailAddresses,genders,interests,locations,names,nicknames,occupations,organizations,phoneNumbers,relations,skills,urls"
}

func (service *PeopleService) List(pageToken *string) (list *people.ListConnectionsResponse, err error) {
	listCall := service.peopleService.People.Connections.List("people/me").PersonFields(service.personFields()).PageSize(500)
	if pageToken != nil {
		listCall.PageToken(*pageToken)
	}
	list, err = listCall.Do()
	return
}

func (service *PeopleService) Update(personId string, person *people.Person) (updated *people.Person, err error) {
	updateCall := service.peopleService.People.UpdateContact(personId, person).PersonFields(service.personFields()).UpdatePersonFields(service.updateFields())
	updated, err = updateCall.Do()
	return
}

func (service *PeopleService) Get(personId string) (person *people.Person, err error) {
	personCall := service.peopleService.People.Get(personId).PersonFields(service.personFields())
	person, err = personCall.Do()
	return
}

func (service *PeopleService) BatchGet(personIds []string) (response *people.GetPeopleResponse, err error) {
	batchGetCall := service.peopleService.People.GetBatchGet().ResourceNames(personIds...).PersonFields(service.personFields())
	response, err = batchGetCall.Do()
	return
}

func (service *PeopleService) BatchUpdate(updates map[string]people.Person) (response *people.BatchUpdateContactsResponse, err error) {
	request := people.BatchUpdateContactsRequest{
		Contacts:   updates,
		ReadMask:   service.personFields(),
		UpdateMask: service.updateFields(),
	}
	updateCall := service.peopleService.People.BatchUpdateContacts(&request)
	response, err = updateCall.Do()
	return
}

func (service *PeopleService) BatchDelete(personIds []string) (err error) {
	request := people.BatchDeleteContactsRequest{
		ResourceNames: personIds,
	}
	updateCall := service.peopleService.People.BatchDeleteContacts(&request)
	_, err = updateCall.Do()
	return
}
