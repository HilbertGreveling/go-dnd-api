package policies

import "github.com/hilbertgreveling/dnd-character-api/models"

type CharacterPolicy interface {
	CanView(userID int, character *models.Character) bool
	CanEdit(userID int, character *models.Character) bool
	CanDelete(userID int, character *models.Character) bool
}

type CharacterPolicyImpl struct{}

func NewCharacterPolicy() CharacterPolicy {
	return &CharacterPolicyImpl{}
}

func (p *CharacterPolicyImpl) CanView(userID int, character *models.Character) bool {
	return true
}

func (p *CharacterPolicyImpl) CanEdit(userID int, character *models.Character) bool {
	return character.UserID == userID
}

func (p *CharacterPolicyImpl) CanDelete(userID int, character *models.Character) bool {
	return character.UserID == userID
}
