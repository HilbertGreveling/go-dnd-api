package repository

import "github.com/hilbertgreveling/dnd-character-api/models"

type UserRepository interface {
	Create(user *models.User) (int, error)
	GetByID(id int) (*models.User, error)
	GetByUsername(id string) (*models.User, error)
}
