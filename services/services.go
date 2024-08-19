package services

import (
	"github.com/hilbertgreveling/dnd-character-api/db"
	"github.com/hilbertgreveling/dnd-character-api/repository/sqlite"
)

type Services struct {
	CharacterService CharacterService
	UserService      UserService
	AuthService      AuthService
}

func SetupServices() *Services {
	db := db.GetDB()

	characterRepo := sqlite.NewCharacterRepositorySQLite(db)
	userRepo := sqlite.NewUserRepositorySQLite(db)
	authRepo := sqlite.NewAuthRepositorySQLite(db)

	// TODO: Add policies here
	characterService := NewCharacterService(characterRepo, userRepo)
	userService := NewUserService(userRepo)
	authService := NewAuthService(authRepo, userRepo)

	return &Services{
		CharacterService: characterService,
		UserService:      userService,
		AuthService:      authService,
	}
}
