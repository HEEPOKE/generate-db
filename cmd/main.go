package main

import (
	"fmt"
	"log"

	"github.com/HEEPOKE/generate-db/internals/domains/repositories"
	server "github.com/HEEPOKE/generate-db/internals/http"
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

	databases.CheckRedis()

	generateRepository := repositories.NewGenerateRepository(db, databases.Rdb)

	address := fmt.Sprintf(":%s", config.Cfg.PORT)
	http := server.NewServer(generateRepository)
	http.RouteInit(address)
}
