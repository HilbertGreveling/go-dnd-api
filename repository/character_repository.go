package repository

import "github.com/hilbertgreveling/dnd-character-api/models"

type CharacterRepository interface {
	Create(character *models.Character) error
	GetAll() ([]*models.Character, error)
	GetByID(id int) (*models.Character, error)
	Update(character *models.Character) error
	Delete(id int) error
}
