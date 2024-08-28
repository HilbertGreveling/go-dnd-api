package services

import (
	"context"
	"errors"

	"github.com/hilbertgreveling/dnd-character-api/models"
	"github.com/hilbertgreveling/dnd-character-api/policies"
	"github.com/hilbertgreveling/dnd-character-api/repository"
	"github.com/hilbertgreveling/dnd-character-api/utils"
)

type CharacterService interface {
	Create(character *models.Character) error
	GetAll() ([]*models.Character, error)
	GetByID(id int) (*models.Character, error)
	Update(character *models.Character, ctx context.Context) error
	Delete(character *models.Character, ctx context.Context) error
}

type characterServiceImpl struct {
	characterRepo   repository.CharacterRepository
	userRepo        repository.UserRepository
	characterPolicy policies.CharacterPolicy
}

func NewCharacterService(characterRepo repository.CharacterRepository, userRepo repository.UserRepository, characterPolicy policies.CharacterPolicy) CharacterService {
	return &characterServiceImpl{
		characterRepo:   characterRepo,
		userRepo:        userRepo,
		characterPolicy: characterPolicy,
	}
}

func (s *characterServiceImpl) Create(character *models.Character) error {
	if _, err := s.userRepo.GetByID(character.UserID); err != nil {
		return errors.New("user does not exist")
	}

	return s.characterRepo.Create(character)
}

func (s *characterServiceImpl) GetAll() ([]*models.Character, error) {
	return s.characterRepo.GetAll()
}

func (s *characterServiceImpl) GetByID(id int) (*models.Character, error) {
	return s.characterRepo.GetByID(id)
}

func (s *characterServiceImpl) Update(character *models.Character, ctx context.Context) error {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return err
	}

	if !s.characterPolicy.CanEdit(userID, character) {
		return errors.New("permission denied: cannot update this character")
	}

	return s.characterRepo.Update(character)
}

func (s *characterServiceImpl) Delete(character *models.Character, ctx context.Context) error {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return err
	}

	if !s.characterPolicy.CanDelete(userID, character) {
		return errors.New("permission denied: cannot delete this character")
	}

	return s.characterRepo.Delete(character.ID)
}
