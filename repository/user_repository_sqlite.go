package repository

import (
	"database/sql"
	"errors"

	"github.com/hilbertgreveling/dnd-character-api/models"
)

type UserRepositorySQLite struct {
	db *sql.DB
}

func NewUserRepositorySQLite(db *sql.DB) *UserRepositorySQLite {
	return &UserRepositorySQLite{db: db}
}

func (repo *UserRepositorySQLite) Create(user *models.User) error {
	query := "INSERT INTO users (username, password) VALUES (?, ?)"
	_, err := repo.db.Exec(query, user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepositorySQLite) GetByUsername(username string) (*models.User, error) {
	query := "SELECT id, username, password FROM users WHERE username = ?"
	row := repo.db.QueryRow(query, username)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}
