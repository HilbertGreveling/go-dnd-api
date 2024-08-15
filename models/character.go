package models

type Character struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Level       int    `json:"level"`
	Description string `json:"description"`
	UserId      int    `json:"user_id"`
}
