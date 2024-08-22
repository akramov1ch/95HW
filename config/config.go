package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Environment faylini yuklab bo'lmadi, default konfiguratsiya qo'llaniladi")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DATABASE_URL muhit o'zgaruvchisi aniqlanmagan")
	}

	return &Config{
		DatabaseURL: dbURL,
	}, nil
}
