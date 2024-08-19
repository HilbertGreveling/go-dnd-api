package repository

import "github.com/hilbertgreveling/dnd-character-api/models"

type AuthRepository interface {
	Create(authUser models.Auth) error
	GetByUsername(id string) (*models.Auth, error)
}
