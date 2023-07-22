package main

import (
	"log"

	"github.com/HEEPOKE/generate-db/pkg/config"
	"github.com/HEEPOKE/generate-db/pkg/databases"
)

func main() {
	_, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := databases.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}
}
