package main

import (
	"fmt"
	"log"

	"github.com/HEEPOKE/generate-db/internals/domains/repositories"
	server "github.com/HEEPOKE/generate-db/internals/http"
	"github.com/HEEPOKE/generate-db/pkg/config"
	"github.com/HEEPOKE/generate-db/pkg/databases"
)

// @title Generate Mock-up API
// @version 1.0
// @description Auto Generate Mock-up Data
// @contact.name HEEPOKE Support
// @contact.url https://github.com/HEEPOKE
// @contact.email Damon1FX@gmail.com
// @host localhost:6476
// @BasePath /apis
// @schemes http
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
