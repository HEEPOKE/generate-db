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

	sqlDB, err := databases.ConnectDB()
	if err != nil {
		panic("Failed to connect to SQL database: " + err.Error())
	}

	mongoDB, err := databases.ConnectMongoDB(config.Cfg.DB_DSN)
	if err != nil {
		panic("Failed to connect to MongoDB: " + err.Error())
	}

	generateRepository := repositories.NewGenerateRepository(db, databases.Rdb)
	insertRepository := repositories.NewInsertRepository(sqlDB, mongoDB)

	address := fmt.Sprintf(":%s", config.Cfg.PORT)
	http := server.NewServer(generateRepository, insertRepository)
	http.RouteInit(address)
}
