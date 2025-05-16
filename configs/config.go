package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	App  AppConfig
	Db   DbConfig
	Auth AuthConfig
}

type AppConfig struct {
	Port int
}

type DbConfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default config")
	}

	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8080" // дефолтный порт, если в .env пусто
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Printf("Invalid PORT value %q, fallback to 8080\n", portStr)
		port = 8080
	}

	return &Config{
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("TOKEN"),
		},
		App: AppConfig{
			Port: port,
		},
	}
}
