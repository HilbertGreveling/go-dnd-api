package services

import (
	"github.com/hilbertgreveling/dnd-character-api/models"
	"github.com/hilbertgreveling/dnd-character-api/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(character *models.User) error
	GetAll() ([]*models.UserResponse, error)
	GetByID(id int) (*models.UserResponse, error)
}

type UserServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &UserServiceImpl{userRepo: userRepo}
}

func (s *UserServiceImpl) Create(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return s.userRepo.Create(user)
}

func (s *UserServiceImpl) GetAll() ([]*models.UserResponse, error) {
	return s.userRepo.GetAll()
}

func (s *UserServiceImpl) GetByID(id int) (*models.UserResponse, error) {
	return s.userRepo.GetByID(id)
}
