package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Cfg *Config
)

type Config struct {
	POSTGRES_DATABASE      string
	POSTGRES_USER          string
	POSTGRES_PASSWORD      string
	POSTGRES_HOST          string
	POSTGRES_PORT          string
	POSTGRES_SSL           string
	POSTGRES_TIMEZONE      string
	REDIS_URL              string
	REDIS_PORT             string
	REDIS_PASSWORD         string
	REDIS_REPLICATION_MODE string
	PORT                   string
	DB_TYPE                string
	DB_DSN                 string
	DB_NAME                string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	config := &Config{
		POSTGRES_DATABASE:      os.Getenv("POSTGRES_DATABASE"),
		POSTGRES_USER:          os.Getenv("POSTGRES_USER"),
		POSTGRES_PASSWORD:      os.Getenv("POSTGRES_PASSWORD"),
		POSTGRES_HOST:          os.Getenv("POSTGRES_HOST"),
		POSTGRES_PORT:          os.Getenv("POSTGRES_PORT"),
		POSTGRES_SSL:           os.Getenv("POSTGRES_SSL"),
		POSTGRES_TIMEZONE:      os.Getenv("POSTGRES_TIMEZONE"),
		REDIS_URL:              os.Getenv("REDIS_URL"),
		REDIS_PORT:             os.Getenv("REDIS_PORT"),
		REDIS_PASSWORD:         os.Getenv("REDIS_PASSWORD"),
		REDIS_REPLICATION_MODE: os.Getenv("REDIS_REPLICATION_MODE"),
		PORT:                   os.Getenv("PORT"),
		DB_TYPE:                os.Getenv("DB_TYPE"),
		DB_DSN:                 os.Getenv("DB_DSN"),
		DB_NAME:                os.Getenv("DB_NAME"),
	}

	Cfg = config

	return config, nil
}
