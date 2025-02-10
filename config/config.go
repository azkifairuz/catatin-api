package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	DBSource string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("no env found....")
	}

	return Config{
		DBSource: os.Getenv("DATABASE_URL"),
	}
}