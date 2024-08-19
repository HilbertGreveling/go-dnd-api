package sqlite

import (
	"database/sql"

	"github.com/hilbertgreveling/dnd-character-api/models"
)

type UserRepositorySQLite struct {
	db *sql.DB
}

func NewUserRepositorySQLite(db *sql.DB) *UserRepositorySQLite {
	return &UserRepositorySQLite{db: db}
}

func (r *UserRepositorySQLite) Create(user *models.User) (int, error) {
	result, err := r.db.Exec("INSERT INTO users (username) VALUES (?)", user.Username)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *UserRepositorySQLite) GetByUsername(username string) (*models.User, error) {
	user := &models.User{}
	err := r.db.QueryRow("SELECT id, username FROM users WHERE username = ?", username).
		Scan(&user.ID, &user.Username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepositorySQLite) GetByID(id int) (*models.User, error) {
	user := &models.User{}
	err := r.db.QueryRow("SELECT id, username FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.Username)
	if err != nil {
		return nil, err
	}

	return user, nil
}
