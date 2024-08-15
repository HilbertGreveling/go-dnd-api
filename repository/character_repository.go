package repository

import "github.com/hilbertgreveling/dnd-character-api/models"

type CharacterRepository interface {
	GetAll() ([]*models.Character, error)
	GetByID(id int) (*models.Character, error)
	Create(character *models.Character) error
	Update(character *models.Character) error
	Delete(id int) error
}
