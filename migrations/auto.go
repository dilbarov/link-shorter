package main

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"link-shorter/configs"
	"link-shorter/internal/link/models"
	"link-shorter/pkg/db"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	dbManager := db.NewDb(&configs.DbConfig{
		Dsn: os.Getenv("DSN"),
	})

	err = dbManager.AutoMigrate(&models.LinkModel{})
	if err != nil {
		return
	}

	log.Info().Msg("Database migrated")
}
