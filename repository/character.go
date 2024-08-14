package repository

const (
	SelectCharacterByID = `
		SELECT id, name, class, level, race_id, description
		FROM characters
		WHERE id = ?`

	SelectAllCharacters = `
		SELECT id, name, class, level, race_id, description
		FROM characters`

	InsertCharacter = `
		INSERT INTO characters (name, class, level, race_id, description)
		VALUES (?, ?, ?, ?, ?)`

	UpdateCharacter = `
		UPDATE characters
		SET name = ?, class = ?, level = ?, race_id = ?, description = ?
		WHERE id = ?`
)
