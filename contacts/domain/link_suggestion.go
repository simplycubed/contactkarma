package domain

import (
	"github.com/simplycubed/contactkarma/contacts/gen/models"
)

type LinkSuggestionID string

func (c LinkSuggestionID) String() string {
	return string(c)
}

type LinkSuggestionKey string

func (c LinkSuggestionKey) String() string {
	return string(c)
}

const KeyEmail LinkSuggestionKey = "email"
const KeyName LinkSuggestionKey = "name"
const KeyPhone LinkSuggestionKey = "phone"

type LinkMatch struct {
	UnifiedId   UnifiedId `firestore:"unified_id" json:"unified_id"`
	DisplayName string    `firestore:"display_name" json:"display_name"`
}

type LinkSuggestion struct {
	ID      LinkSuggestionID  `firestore:"id" json:"id"`
	Key     LinkSuggestionKey `firestore:"key" json:"key"`
	Value   string            `firestore:"value" json:"value"`
	Matches []LinkMatch       `firestore:"matches" json:"matches"`
}

type LinkSuggestions []LinkSuggestion

func (match LinkMatch) MapToDto() (dto *models.LinkMatch) {
	dto = &models.LinkMatch{}
	dto.DisplayName = match.DisplayName
	dto.UnifiedID = string(match.UnifiedId)
	return
}
func (suggestion LinkSuggestion) MapToDto() (dto *models.LinkSuggestion) {
	dto = &models.LinkSuggestion{}
	dto.ID = string(suggestion.ID)
	dto.Key = string(suggestion.Key)
	dto.Value = suggestion.Value
	for _, match := range suggestion.Matches {
		dto.Matches = append(dto.Matches, match.MapToDto())
	}
	return
}

func (suggestions LinkSuggestions) MapToDto() (dto []*models.LinkSuggestion) {
	for _, suggestion := range suggestions {
		dto = append(dto, suggestion.MapToDto())
	}
	return
}
