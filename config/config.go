package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Environment faylini yuklab bo'lmadi, default konfiguratsiya qo'llaniladi")
	}

	return &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
}
