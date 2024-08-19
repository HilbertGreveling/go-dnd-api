package sqlite

import (
	"database/sql"

	"github.com/hilbertgreveling/dnd-character-api/models"
)

type AuthRepositorySQLite struct {
	db *sql.DB
}

func NewAuthRepositorySQLite(db *sql.DB) *AuthRepositorySQLite {
	return &AuthRepositorySQLite{db: db}
}

func (r *AuthRepositorySQLite) Create(auth models.Auth) error {
	query := `INSERT INTO auth (password, user_id) VALUES (?, ?)`
	_, err := r.db.Exec(query, auth.Password, auth.UserID)
	return err
}

func (r *AuthRepositorySQLite) GetByUsername(username string) (*models.Auth, error) {
	var userID int
	userQuery := `SELECT id FROM users WHERE username = ?`
	err := r.db.QueryRow(userQuery, username).Scan(&userID)
	if err != nil {
		return nil, err
	}

	var auth models.Auth
	authQuery := `SELECT id, password, user_id FROM auth WHERE user_id = ?`
	err = r.db.QueryRow(authQuery, userID).Scan(&auth.ID, &auth.Password, &auth.UserID)
	if err != nil {
		return nil, err
	}

	return &auth, nil
}

func (r *AuthRepositorySQLite) GetByID(id int) (*models.Auth, error) {
	var auth models.Auth
	query := `SELECT id, password, user_id FROM auth WHERE id = ?`
	err := r.db.QueryRow(query, id).Scan(&auth.ID, &auth.Password, &auth.UserID)
	if err != nil {
		return nil, err
	}

	return &auth, nil
}
