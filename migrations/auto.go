package main

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"link-shorter/configs"
	clickModels "link-shorter/internal/click/models"
	linkModels "link-shorter/internal/link/models"
	userModels "link-shorter/internal/user/models"
	"link-shorter/pkg/db"
	"link-shorter/pkg/logger"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	log.Logger = logger.SetupLogger(os.Getenv("Env"))

	dbManager := db.NewDb(&configs.DbConfig{
		Dsn: os.Getenv("DSN"),
	})

	err = dbManager.AutoMigrate(&linkModels.Model{}, &userModels.Model{}, clickModels.Model{})
	if err != nil {
		return
	}

	log.Info().Msg("Database migrated")
}
