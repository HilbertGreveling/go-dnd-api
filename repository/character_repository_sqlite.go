package repository

import (
	"database/sql"
	"errors"

	"github.com/hilbertgreveling/dnd-character-api/models"
)

type CharacterRepositorySQLite struct {
	db *sql.DB
}

func NewCharacterRepositorySQLite(db *sql.DB) *CharacterRepositorySQLite {
	return &CharacterRepositorySQLite{db: db}
}

func (r *CharacterRepositorySQLite) Create(character *models.Character) error {
	query := "INSERT INTO characters (name, level, description) VALUES (?, ?, ?)"
	result, err := r.db.Exec(query, character.Name, character.Level, character.Description)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	character.ID = int(id)
	return nil
}

func (r *CharacterRepositorySQLite) GetAll() ([]*models.Character, error) {
	query := "SELECT id, name, level, description FROM characters"
	rows, err := r.db.Query(query)
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

func (r *CharacterRepositorySQLite) GetByID(id int) (*models.Character, error) {
	var character models.Character
	query := "SELECT id, name, level, description FROM characters WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(&character.ID, &character.Name, &character.Level, &character.Description)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &character, nil
}

func (r *CharacterRepositorySQLite) Update(character *models.Character) error {
	query := "UPDATE characters SET name = ?, level = ?, description = ? WHERE id = ?"
	_, err := r.db.Exec(query, character.Name, character.Level, character.Description, character.ID)
	return err
}

func (r *CharacterRepositorySQLite) Delete(id int) error {
	query := "DELETE FROM characters WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}
