package services

import (
	"errors"

	"github.com/hilbertgreveling/dnd-character-api/models"
	"github.com/hilbertgreveling/dnd-character-api/repository"
)

type CharacterService interface {
	Create(character *models.Character) error
	GetAll() ([]*models.Character, error)
	GetByID(id int) (*models.Character, error)
	Update(character *models.Character) error
	Delete(id int) error
}

type characterServiceImpl struct {
	characterRepo repository.CharacterRepository
	userRepo      repository.UserRepository
}

func NewCharacterService(characterRepo repository.CharacterRepository, userRepo repository.UserRepository) CharacterService {
	return &characterServiceImpl{
		characterRepo: characterRepo,
		userRepo:      userRepo,
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

func (s *characterServiceImpl) Update(character *models.Character) error {
	return s.characterRepo.Update(character)
}

func (s *characterServiceImpl) Delete(id int) error {
	return s.characterRepo.Delete(id)
}
