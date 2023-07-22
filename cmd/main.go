package main

import (
	"fmt"
	"log"

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

	address := fmt.Sprintf(":%s", config.Cfg.PORT)
	http := server.NewServer()
	http.RouteInit(address)
}
