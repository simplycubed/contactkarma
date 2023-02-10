package application

import (
	"context"

	"github.com/simplycubed/contactkarma/contacts/domain"
	"github.com/simplycubed/contactkarma/contacts/domain/repository"
	"github.com/simplycubed/contactkarma/contacts/gen/models"
)

type tagService struct {
	tagRepo repository.ITag
}

func NewTagService(tagRepo repository.ITag) *tagService {
	return &tagService{tagRepo: tagRepo}
}

func (a *tagService) SaveTag(ctx context.Context, uID domain.UserID, cID domain.UnifiedId, tag *models.Tag) (*models.Tag, error) {
	tagToCreate := domain.Tag{}
	tagToCreate.FromDto(tag)
	created, err := a.tagRepo.SaveTag(ctx, uID, cID, tagToCreate)
	if err != nil {
		return nil, err
	}
	return created.MapToDto(), nil
}
func (a *tagService) UpdateTag(ctx context.Context, uID domain.UserID, cID domain.UnifiedId, nID domain.TagID, tag *models.Tag) (*models.Tag, error) {
	_, err := a.tagRepo.GetTagByID(ctx, uID, cID, nID)
	if err != nil {
		return nil, err
	}

	tagToUpdate := domain.Tag{}
	tagToUpdate.FromDto(tag)
	if err := a.tagRepo.UpdateTag(ctx, uID, cID, nID, tagToUpdate); err != nil {
		return nil, err
	}
	updatedNote, err := a.tagRepo.GetTagByID(ctx, uID, cID, nID)
	if err != nil {
		return nil, err
	}
	return updatedNote.MapToDto(), nil
}
func (a *tagService) GetAllTags(ctx context.Context, uID domain.UserID, cID domain.UnifiedId) ([]*models.Tag, error) {
	tags, err := a.tagRepo.GetAllTags(ctx, uID, cID)
	if err != nil {
		return nil, err
	}
	return domain.Tags(tags).MapToDto(), nil
}

func (a *tagService) GetTags(ctx context.Context, uID domain.UserID, cID domain.UnifiedId, limit int, lastDocumentId *domain.TagID) ([]*models.Tag, error) {
	tags, err := a.tagRepo.GetTags(ctx, uID, cID, limit, lastDocumentId)
	if err != nil {
		return nil, err
	}
	return domain.Tags(tags).MapToDto(), nil
}

func (a *tagService) DeleteTag(ctx context.Context, uID domain.UserID, cID domain.UnifiedId, nID domain.TagID) error {
	return a.tagRepo.DeleteTag(ctx, uID, cID, nID)
}
