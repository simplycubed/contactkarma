package servicefactory

import (
	"context"
	"time"

	"github.com/simplycubed/contactkarma/contacts/adapters/googlepeople"
	"github.com/simplycubed/contactkarma/contacts/application"
	"golang.org/x/oauth2"
)

type PeopleServiceFactory struct {
	googleOAuthService application.GoogleOAuthService
}

func NewPeopleServiceFactory(googleOAuthService application.GoogleOAuthService) *PeopleServiceFactory {
	return &PeopleServiceFactory{googleOAuthService: googleOAuthService}
}

func (p *PeopleServiceFactory) New(ctx context.Context, accessToken string, refreshToken string, expiry time.Time) (service application.PeopleService) {
	token := &oauth2.Token{AccessToken: accessToken, RefreshToken: refreshToken, Expiry: expiry}
	config := p.googleOAuthService.GetConfig()
	service = googlepeople.NewPeopleService(ctx, config, token)
	return
}
