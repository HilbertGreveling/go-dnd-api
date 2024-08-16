package repository

import "github.com/hilbertgreveling/dnd-character-api/models"

type UserRepository interface {
	Create(user *models.User) error
	GetAll() ([]*models.UserResponse, error)
	GetByID(id int) (*models.UserResponse, error)
}
