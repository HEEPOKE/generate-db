package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/HEEPOKE/generate-db/internals/domains/repositories"
	server "github.com/HEEPOKE/generate-db/internals/http"
	"github.com/HEEPOKE/generate-db/pkg/config"
	"github.com/HEEPOKE/generate-db/pkg/databases"
	"github.com/HEEPOKE/generate-db/pkg/enums"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
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
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := databases.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	databases.CheckRedis()

	if cfg.DB_TYPE == "" || cfg.DB_DSN == "" || cfg.DB_NAME == "" {
		log.Fatal("DB_TYPE, DB_DSN, and DB_NAME must be provided")
	}

	var sqlDB *gorm.DB
	var mongoDB *mongo.Database

	if cfg.DB_TYPE != string(enums.MONGO) {
		sqlDB, err = databases.ConnectDB()
		if err != nil {
			log.Fatalf("Failed to connect to SQL database: %v", err)
		}
	} else {
		if strings.HasPrefix(cfg.DB_DSN, "mongodb://") || strings.HasPrefix(cfg.DB_DSN, "mongodb+srv://") {
			mongoDB, err = databases.ConnectMongoDB(cfg.DB_DSN)
			if err != nil {
				log.Fatalf("Failed to connect to MongoDB: %v", err)
			}
		} else {
			log.Fatal("Invalid MongoDB connection string. Scheme must be 'mongodb' or 'mongodb+srv'")
		}
	}

	generateRepository := repositories.NewGenerateRepository(db, databases.Rdb)
	insertRepository := repositories.NewInsertRepository(sqlDB, mongoDB)
	utilitiesRepository := repositories.NewUtilitiesRepository(db, databases.Rdb)

	address := fmt.Sprintf(":%s", config.Cfg.PORT)
	http := server.NewServer(generateRepository, insertRepository, utilitiesRepository)
	http.RouteInit(address)
}
