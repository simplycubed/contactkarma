.PHONY: setup gen clean

lint:
	golangci-lint run --issues-exit-code 0

gen: clean # swagger codegen
	swagger generate server -t gen -f ./openapi/api.yml --exclude-main -A contacts -P github.com/simplycubed/contactkarma/contacts/domain.User
	swagger generate server -t gen-jobs -f ./openapi/jobs.yml --exclude-main -A contacts-jobs

gen-client: clean # swagger codegen
	swagger generate client -t gen -f ./openapi/api.yml -A contacts -P github.com/simplycubed/contactkarma/contacts/domain.User
	swagger generate client -t gen-jobs -f ./openapi/jobs.yml -A contacts-jobs

gen-mocks: clean
	mockgen -destination gen/mocks/mock_application/google_oauth_service_mock.go -source application/contact_source.go GoogleOAuthService