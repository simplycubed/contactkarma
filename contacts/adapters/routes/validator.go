package routes

import (
	"fmt"

	"github.com/simplycubed/contactkarma/contacts/gen/models"
)

const defaultResultCount = 50
const maxAllowedLimit = 100

func ValidateLimit(limit *int64) (validatedLimit int, err *models.ErrorResponse) {
	validatedLimit = defaultResultCount
	if limit != nil {
		if *limit > maxAllowedLimit {
			err = &models.ErrorResponse{
				Description: fmt.Sprintf("limit can't be more than %d", maxAllowedLimit),
				Error:       "limit exceeded max limit",
			}
			return
		}
		validatedLimit = int(*limit)
	}
	return
}
