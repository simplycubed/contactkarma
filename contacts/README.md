# contacts

Contact Karma's Contact API

## Authenticating

- [Getting started with API Gateway and Cloud Run](https://cloud.google.com/api-gateway/docs/get-started-cloud-run)
- [Using Firebase to authenticate users](https://cloud.google.com/api-gateway/docs/authenticating-users-firebase)
- [Auth locally withGOOGLE_APPLICATION_CREDENTIALS env var](https://cloud.google.com/docs/authentication/getting-started#setting_the_environment_variable)

## Emulators & Local Dev

- [The Local Firebase Emulator UI in 15 minutes](https://www.youtube.com/watch?v=pkgvFNPdiEs)

- Firebase Emulators are executed with docker-compose.
- Firestore emulator is using port `9090` instead of `8080` as `8080` is being used by the contacts service.

```bash
cd ./emulators
docker-compose build
docker-compose up -d
```

## Environment Variables

The following environment variables are required.
For local development a .env file can be created from .env.sample.

| Env Var                        | Example                            | Notes                                        |
| :----------------------------- | :--------------------------------- | :------------------------------------------- |
| ALLOWED_ORIGIN                 | "*"                                | Needs to be restricted                       |
| ENV                            | "local"                            | local, dev, prod                             |
| FIREBASE_URL                   | "contactkarma-dev.firebaseapp.com" |                                              |
| FRONTEND_URL                   | "https://contactkarma.dev"         |                                              |
| GOOGLE_CLOUD_PROJECT           | "contactkarma-dev"                 |                                              |
| GOOGLE_APPLICATION_CREDENTIALS | "/path/to/key/file/key.json"       | Required for local dev to auth with Firebase |
| PORT                           | "8080"                             | Cloud Run default                            |

### Used for local dev with emulators

| Env Var                     | Example                               | Notes |
| :-------------------------- | :------------------------------------ | :---- |
| FIREBASE_AUTH_EMULATOR_HOST | "localhost:9090"                      |       |
| FIREBASE_URL                | "http://localhost:4000?ns=emulatorui" |       |
| FIRESTORE_EMULATOR_HOST     | "localhost:9090"                      |       |

## Authentication

- Firebase Auth is validated by API Gateway
- The origional Firebase Auth JWT is copied from Authorization to X-Apigateway-Api-Userinfo

```yaml
securityDefinitions:
  Bearer:
    type: apiKey
    in: header
    name: X-Apigateway-Api-Userinfo
    description: JWT Token
security:
  - Bearer: []
```

## Running Tests locally

Run test using the command (replace $PATH_TO_ENV_FILE with absoulte file path to env file):

  `TEST_ENV_PATH=$PATH_TO_ENV_FILE go test -cover ./...`

Optionally set TEST_ENV_PATH in vscode extension settings and use test runner tool

```json
"go.testEnvVars": {
  "TEST_ENV_PATH": "~INSERT_ABSOLUTE_PATH_TO_ENV_FILE~"
}
```
