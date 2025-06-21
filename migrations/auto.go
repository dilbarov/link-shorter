package main

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"link-shorter/configs"
	"link-shorter/internal/click"
	"link-shorter/internal/link"
	"link-shorter/internal/user"
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

	err = dbManager.AutoMigrate(&link.Model{}, &user.Model{}, click.Model{})
	if err != nil {
		return
	}

	log.Info().Msg("Database migrated")
}
