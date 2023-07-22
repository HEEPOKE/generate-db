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
	MYSQL_DATABASE         string
	MYSQL_USER             string
	MYSQL_ROOT_PASSWORD    string
	MYSQL_HOST             string
	MYSQL_HOST_PORT        string
	MYSQL_CONTAINER_PORT   string
	REDIS_URL              string
	REDIS_PORT             string
	REDIS_PASSWORD         string
	REDIS_REPLICATION_MODE string
	PORT                   string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	config := &Config{
		MYSQL_DATABASE:         os.Getenv("MYSQL_DATABASE"),
		MYSQL_USER:             os.Getenv("MYSQL_USER"),
		MYSQL_ROOT_PASSWORD:    os.Getenv("MYSQL_ROOT_PASSWORD"),
		MYSQL_HOST:             os.Getenv("MYSQL_HOST"),
		MYSQL_HOST_PORT:        os.Getenv("MYSQL_HOST_PORT"),
		MYSQL_CONTAINER_PORT:   os.Getenv("MYSQL_CONTAINER_PORT"),
		REDIS_URL:              os.Getenv("REDIS_URL"),
		REDIS_PORT:             os.Getenv("REDIS_PORT"),
		REDIS_PASSWORD:         os.Getenv("REDIS_PASSWORD"),
		REDIS_REPLICATION_MODE: os.Getenv("REDIS_REPLICATION_MODE"),
		PORT:                   os.Getenv("PORT"),
	}

	Cfg = config

	return config, nil
}
