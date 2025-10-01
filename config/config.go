package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BotToken string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, reading from environment")
	}

	return Config{
		BotToken: os.Getenv("TGBTOKEN"),
	}
}