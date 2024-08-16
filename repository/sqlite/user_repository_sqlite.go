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

func (r *UserRepositorySQLite) GetAll() ([]*models.UserResponse, error) {
	query := "SELECT id, username FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.UserResponse
	for rows.Next() {
		var user models.UserResponse
		if err := rows.Scan(&user.ID, &user.Username); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
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

func (r *UserRepositorySQLite) GetByUsername(username string) (*models.User, error) {
	query := `SELECT id, username, password FROM users WHERE username = ?`
	row := r.db.QueryRow(query, username)

	var user models.User
	if err := row.Scan(&user.ID, &user.Username, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.New("user not found")
	}

	return &user, nil
}
