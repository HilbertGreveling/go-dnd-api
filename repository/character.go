package repository

const (
	SelectCharacterByID = `
		SELECT id, name, level, description
		FROM characters
		WHERE id = ?`

	SelectAllCharacters = `
		SELECT id, name, level, description
		FROM characters`

	InsertCharacter = `
		INSERT INTO characters (name, level, description)
		VALUES (?, ?, ?, ?, ?)`

	UpdateCharacter = `
		UPDATE characters
		SET name = ?, level = ?, description = ?
		WHERE id = ?`
)
