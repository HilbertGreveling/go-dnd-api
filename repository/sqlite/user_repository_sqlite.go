package sqlite

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

func (repo *UserRepositorySQLite) GetByID(id int) (*models.UserResponse, error) {
	user := &models.UserResponse{}
	query := "SELECT id, username FROM users WHERE id = ?"
	err := repo.db.QueryRow(query, id).Scan(&user.ID, &user.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}
