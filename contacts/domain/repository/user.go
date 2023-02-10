//go:generate mockgen --build_flags=--mod=mod -package=mocks -destination=../../mocks/user.go . IUser
package repository

import (
	"context"

	"github.com/simplycubed/contactkarma/contacts/domain"
)

type IUser interface {
	//Get
	GetUserByID(context.Context, domain.UserID) (*domain.User, error)

	//Patch
	UpdateUser(context.Context, domain.UserID, domain.User) error

	//Save
	SaveUser(context.Context, domain.UserID, domain.User) (created *domain.User, err error)

	//Delete
	DeleteUser(ctx context.Context, uID domain.UserID) error
}
