package services

import (
	"github.com/hilbertgreveling/dnd-character-api/models"
	"github.com/hilbertgreveling/dnd-character-api/repository"
)

type UserService interface {
	GetByID(id int) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
}

type UserServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}

func (s *UserServiceImpl) GetByID(id int) (*models.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *UserServiceImpl) GetByUsername(username string) (*models.User, error) {
	return s.userRepo.GetByUsername(username)
}
