// services/auth_service.go
package services

import (
	"errors"

	"github.com/hilbertgreveling/dnd-character-api/models"
	"github.com/hilbertgreveling/dnd-character-api/repository"
	"github.com/hilbertgreveling/dnd-character-api/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	GetAuthByUsername(username string) (*models.Auth, error)
	RegisterUser(user models.User, password string) error
	LoginUser(username, password string) (string, error)
}

type AuthServiceImpl struct {
	authRepo repository.AuthRepository
	userRepo repository.UserRepository
}

func NewAuthService(authRepo repository.AuthRepository, userRepo repository.UserRepository) AuthService {
	return &AuthServiceImpl{
		authRepo: authRepo,
		userRepo: userRepo,
	}
}

func (s *AuthServiceImpl) GetAuthByUsername(username string) (*models.Auth, error) {
	return s.authRepo.GetByUsername(username)
}

func (s *AuthServiceImpl) RegisterUser(user models.User, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	id, err := s.userRepo.Create(&user)
	if err != nil {
		return err
	}

	authUser := models.Auth{
		ID:       id,
		Password: string(hashedPassword),
	}

	if err := s.authRepo.Create(authUser); err != nil {
		return err
	}

	return nil
}

func (s *AuthServiceImpl) LoginUser(username, password string) (string, error) {
	auth, err := s.authRepo.GetByUsername(username)
	if err != nil {
		return "", errors.New("authentication failed")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(password)); err != nil {
		return "", errors.New("invalid password")
	}

	token, err := utils.GenerateJWT(auth.UserID)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
