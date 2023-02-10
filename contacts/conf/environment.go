package conf

import (
	"fmt"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

const (
	_ALLOWED_ORIGIN               = `ALLOWED_ORIGIN`
	_ENV                          = `ENV`
	_FIREBASE_URL                 = `FIREBASE_URL`
	_FRONTEND_URL                 = `FRONTEND_URL`
	_GOOGLE_CLOUD_PROJECT         = `GOOGLE_CLOUD_PROJECT`
	_PORT                         = `PORT`
	_JOB_PORT                     = `JOB_PORT`
	_GOOGLE_AUTH_CLIENT_ID        = `GOOGLE_AUTH_CLIENT_ID`
	_GOOGLE_AUTH_CLIENT_SECRET    = `GOOGLE_AUTH_CLIENT_SECRET`
	_GOOGLE_AUTH_REDIRECT_URL     = `GOOGLE_AUTH_REDIRECT_URL`
	_PULL_CONTACT_SOURCE_TOPIC    = `PULL_CONTACT_SOURCE_TOPIC`
	_CONTACT_SOURCE_DELETED_TOPIC = `CONTACT_SOURCE_DELETED_TOPIC`

	_X_GOOGLE_ISSUER    = `X_GOOGLE_ISSUER`
	_X_GOOGLE_JWKS_URI  = `X_GOOGLE_JWKS_URI`
	_X_GOOGLE_AUDIENCES = `X_GOOGLE_AUDIENCES`
	_TYPESENSE_HOST     = `TYPESENSE_HOST`
	_TYPESENSE_API_KEY  = `TYPESENSE_API_KEY`
)

func AllowedOrigin() string {
	return mustStr(_ALLOWED_ORIGIN)
}

func Env() string {
	return mustStr(_ENV)
}

func FirebaseURL() string {
	return mustStr(_FIREBASE_URL)
}

func FrontendURL() string {
	return mustStr(_FRONTEND_URL)
}

func PORT() int {
	return intWithDefault(_PORT)
}

func JobPort() int {
	return intWithDefault(_JOB_PORT)
}

func PullContactsSourceTopic() string {
	return mustStr(_PULL_CONTACT_SOURCE_TOPIC)
}

func ContactSourceDeletedTopic() string {
	return mustStr(_CONTACT_SOURCE_DELETED_TOPIC)
}

func ProjectID() string {
	return mustStr(_GOOGLE_CLOUD_PROJECT)
}

func GoogleOAuthConfig() (clientId, clientSecret, redirectUri string) {
	return mustStr(_GOOGLE_AUTH_CLIENT_ID), mustStr(_GOOGLE_AUTH_CLIENT_SECRET), mustStr(_GOOGLE_AUTH_REDIRECT_URL)
}

func GatewaySecurityConfig() (issuer, jwksUri, audiences string) {
	return mustStr(_X_GOOGLE_ISSUER), mustStr(_X_GOOGLE_JWKS_URI), mustStr(_X_GOOGLE_AUDIENCES)
}

func TypesenseConfig() (host, apiKey string) {
	return mustStr(_TYPESENSE_HOST), mustStr(_TYPESENSE_API_KEY)
}

func TypesenseHost() (host string) {
	return mustStr(_TYPESENSE_HOST)
}

func TypesenseApiKey() (apiKey string) {
	return mustStr(_TYPESENSE_API_KEY)
}

func mustStr(envName string) string {
	sv := str(envName)
	if sv == `` {
		panic(fmt.Sprintf("%s is missing", envName))
	}

	return sv
}

func str(paramKey string) string {
	return os.Getenv(paramKey) // load from ENV
}

func intWithDefault(envName string) int {
	sv := str(envName)
	if sv == `` {
		panic(fmt.Sprintf("%s is missing", envName))
	}
	iv, err := strconv.Atoi(sv)
	if err != nil {
		panic(fmt.Sprintf("%s is missing", envName))
	}

	return iv
}
