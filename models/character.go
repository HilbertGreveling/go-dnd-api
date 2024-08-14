package models

import (
	"github.com/hilbertgreveling/dnd-character-api/db"
	"github.com/hilbertgreveling/dnd-character-api/repository"
)

type Character struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Level       int    `json:"level"`
	Description string `json:"description"`
	Race        *Race  `json:"race,omitempty"`
	// Class       *Class `json:"class,omitempty"`
}

func GetCharacter(id int) (*Character, error) {
	db := db.GetDB()
	character := &Character{}
	row := db.QueryRow(repository.SelectCharacterByID, id)
	err := row.Scan(&character.ID, &character.Name, &character.Level, &character.Description)
	if err != nil {
		return nil, err
	}

	return character, nil
}

func GetAllCharacters() ([]Character, error) {
	db := db.GetDB()
	rows, err := db.Query(repository.SelectAllCharacters)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var characters []Character

	for rows.Next() {
		var character Character
		err := rows.Scan(&character.ID, &character.Name, &character.Level, &character.Description)
		if err != nil {
			return nil, err
		}

		characters = append(characters, character)
	}

	return characters, nil
}
