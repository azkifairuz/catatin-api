package config

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type Config struct {
	DBSource string
}

var DB *pgxpool.Pool

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("no env found....")
	}

	return Config{
		DBSource: os.Getenv("DATABASE_URL"),
	}
}

func ConnectDB() {

	dbURL := LoadConfig()

	dbpool, err := pgxpool.New(context.Background(), dbURL.DBSource)
	if err != nil {
		log.Fatalf("Gagal konek ke database: %v", err)
	}

	DB = dbpool
	log.Println("Koneksi ke database berhasil!")
}
