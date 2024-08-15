package repository

import "github.com/hilbertgreveling/dnd-character-api/models"

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByUsername(username string) (*models.User, error)
}
