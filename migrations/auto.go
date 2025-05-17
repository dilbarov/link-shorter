package main

import (
	"github.com/joho/godotenv"
	"link-shorter/configs"
	"link-shorter/internal/link"
	"link-shorter/pkg/db"
	"log"
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

	err = dbManager.AutoMigrate(&link.Link{})
	if err != nil {
		return
	}

	log.Printf("Database migrated")
}
