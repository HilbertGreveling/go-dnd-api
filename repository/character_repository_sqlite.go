package repository

import (
	"github.com/hilbertgreveling/dnd-character-api/db"
	"github.com/hilbertgreveling/dnd-character-api/models"
)

type CharacterRepositorySQLite struct{}

func NewCharacterRepositorySQLite() *CharacterRepositorySQLite {
	return &CharacterRepositorySQLite{}
}

func (r *CharacterRepositorySQLite) GetAll() ([]*models.Character, error) {
	db := db.GetDB()

	query := "SELECT id, name, level, description FROM characters"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var characters []*models.Character
	for rows.Next() {
		var character models.Character
		if err := rows.Scan(&character.ID, &character.Name, &character.Level, &character.Description); err != nil {
			return nil, err
		}
		characters = append(characters, &character)
	}
	return characters, nil
}
