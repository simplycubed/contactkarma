package application

import (
	"context"

	"github.com/simplycubed/contactkarma/contacts/domain/repository"
	"github.com/simplycubed/contactkarma/contacts/gen/models"

	"github.com/simplycubed/contactkarma/contacts/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userService struct {
	userRepo repository.IUser
}

func NewUserService(userRepo repository.IUser) *userService {
	return &userService{userRepo: userRepo}
}

func (a *userService) SaveUser(ctx context.Context, uID domain.UserID, user *models.User) (*models.User, error) {
	existing, err := a.userRepo.GetUserByID(ctx, uID)
	if err != nil {
		if !(status.Code(err) == codes.NotFound) {
			return nil, err
		}
		userToCreate := domain.User{}
		userToCreate.FromDto(user)
		created, err := a.userRepo.SaveUser(ctx, uID, userToCreate)
		if err != nil {
			return nil, err
		}
		return created.MapToDto(), nil
	}
	return existing.MapToDto(), nil
}

func (a *userService) UpdateUser(ctx context.Context, uID domain.UserID, req *models.User) (*models.User, error) {
	_, err := a.GetUser(ctx, uID)
	if err != nil {
		return nil, err
	}
	userToUpdate := domain.User{}
	userToUpdate.FromDto(req)
	err = a.userRepo.UpdateUser(ctx, uID, userToUpdate)
	if err != nil {
		return nil, err
	}

	user, err := a.GetUser(ctx, uID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *userService) GetUser(ctx context.Context, uID domain.UserID) (*models.User, error) {
	user, err := a.userRepo.GetUserByID(ctx, uID)
	if err != nil {
		return nil, err
	}
	return user.MapToDto(), nil
}

func (a *userService) DeleteUser(ctx context.Context, uID domain.UserID) error {
	user, err := a.userRepo.GetUserByID(ctx, uID)
	if err != nil {
		return err
	}

	// TODO: delete unified, contact-sources, link-suggestions, update-logs etc
	// delete all contacts of user
	// err = a.defaultSourceContactRepo.DeleteAllContacts(ctx, domain.UserID(user.ID))
	// if err != nil {
	// 	return err
	// }

	if err := a.userRepo.DeleteUser(ctx, domain.UserID(user.ID)); err != nil {
		return err
	}
	return nil
}
