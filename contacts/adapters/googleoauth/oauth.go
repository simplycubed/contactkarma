package googleoauth

import (
	"context"

	"github.com/simplycubed/contactkarma/contacts/conf"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	googleOAuth2 "google.golang.org/api/oauth2/v1"
	"google.golang.org/api/option"
	"google.golang.org/api/people/v1"
)

type GoogleOAuthService struct {
	oauth2Service *googleOAuth2.Service
	clientId      string
	secret        string
	redirectUrl   string
}

func NewGoogleOAuthService(opt ...option.ClientOption) (service *GoogleOAuthService, err error) {
	ctx := context.Background()
	oauth2Service, err := googleOAuth2.NewService(ctx, opt...)
	if err != nil {
		return
	}
	clientId, secret, redirectUrl := conf.GoogleOAuthConfig()
	service = &GoogleOAuthService{oauth2Service: oauth2Service, clientId: clientId, secret: secret, redirectUrl: redirectUrl}
	return
}

// GetOAuthRedirectUrl returns the url to redirect user to google oauth consent screen
func (service *GoogleOAuthService) GetRedirectUrl(ctx context.Context) (url string, err error) {
	config := service.GetConfig()
	url = config.AuthCodeURL("", oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	return
}

func (service *GoogleOAuthService) Scopes() []string {
	return []string{
		"email",
		"profile",
		people.ContactsScope,
	}
}

func (service *GoogleOAuthService) GetToken(ctx context.Context, code string) (token *oauth2.Token, err error) {
	config := &oauth2.Config{
		ClientID:     service.clientId,
		ClientSecret: service.secret,
		RedirectURL:  service.redirectUrl,
		Scopes: []string{
			"email",
			"profile",
			people.ContactsScope,
		},
		Endpoint: google.Endpoint,
	}
	token, err = config.Exchange(ctx, code)
	return
}

func (service *GoogleOAuthService) GetUserInfo(ctx context.Context, accessToken string) (info *googleOAuth2.Tokeninfo, err error) {
	info, err = service.oauth2Service.Tokeninfo().AccessToken(accessToken).Do()
	return
}

func (service *GoogleOAuthService) GetConfig() (config *oauth2.Config) {
	config = &oauth2.Config{
		ClientID:     service.clientId,
		ClientSecret: service.secret,
		RedirectURL:  service.redirectUrl,
		Scopes:       service.Scopes(),
		Endpoint:     google.Endpoint,
	}
	return
}
