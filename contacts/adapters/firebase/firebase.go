package firebase

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"

	f "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/simplycubed/contactkarma/contacts/domain"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
)

type Firebase interface {
	ParseEncodedUser(encodedUser string) (*UserInfo, error)
}

type firebase struct {
	a              *auth.Client
	ctx            context.Context
	Identities     Identities `json:"identities"`
	SignInProvider string     `json:"sign_in_provider"`
}

type UserID string

type UserInfo struct {
	Name          string       `json:"name"`
	Picture       string       `json:"picture"`
	Iss           string       `json:"iss"`
	Aud           string       `json:"aud"`
	AuthTime      int          `json:"auth_time"`
	UserID        UserID       `json:"user_id"`
	Sub           string       `json:"sub"`
	Iat           int          `json:"iat"`
	Exp           int          `json:"exp"`
	Email         string       `json:"email"`
	EmailVerified bool         `json:"email_verified"`
	Firebase      firebase     `json:"firebase"`
	StripeRole    *domain.Role `json:"stripeRole"`
}
type Identities struct {
	GoogleCom []string `json:"google.com"`
	Email     []string `json:"email"`
}

func NewFirebase(databaseURL string, opt ...option.ClientOption) (Firebase, error) {
	ctx := context.Background()
	firebaseApp, err := f.NewApp(ctx, &f.Config{
		DatabaseURL: databaseURL,
	}, opt...)
	if err != nil {
		return nil, errors.Wrap(err, "firebase.NewApp")
	}

	a, err := firebaseApp.Auth(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Auth")
	}

	return &firebase{
		a:   a,
		ctx: ctx,
	}, nil
}

// ParseEncodedUser is a piece of middleware that will parse the gatewayUserInfoHeader
func (f *firebase) ParseEncodedUser(encodedUser string) (*UserInfo, error) {
	var userToken UserInfo
	if encodedUser == "" {
		return &userToken, nil
	}

	decodedBytes, err := base64.RawURLEncoding.DecodeString(encodedUser)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(bytes.NewReader(decodedBytes))
	err = decoder.Decode(&userToken)
	if err != nil {
		return &userToken, err
	}

	return &userToken, nil
}

const GatewayUserInfoHeader = "X-Apigateway-Api-Userinfo"
const gatewayUserContext = "GATEWAY_USER"

func HandleAuth(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		encodedUser := r.Header.Get(GatewayUserInfoHeader)
		if encodedUser == "" {
			http.Error(w, "User Not Available", http.StatusForbidden)
			return
		}

		// Accept both full Firebase Auth JWT with header, payload, and claims
		// as well as just the payload which is sent from API Gateway
		var token string
		tokens := strings.Split(encodedUser, ".")
		if len(tokens) > 1 {
			token = tokens[1]
		} else {
			token = tokens[0]
		}

		decodedBytes, err := base64.RawURLEncoding.DecodeString(token)
		if err != nil {
			http.Error(w, "Invalid UserInfo", http.StatusForbidden)
			return
		}

		decoder := json.NewDecoder(bytes.NewReader(decodedBytes))
		var userToken UserInfo
		err = decoder.Decode(&userToken)
		if err != nil {
			http.Error(w, "Invalid UserInfo", http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), gatewayUserContext, userToken)
		h.ServeHTTP(w, r.WithContext(ctx))

	}
}
