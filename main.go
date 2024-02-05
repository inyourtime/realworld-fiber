package main

import (
	"fmt"
	"log"
	"os"
	"realworld-go-fiber/adapter/handler/restful"
	"realworld-go-fiber/adapter/logger"
	"realworld-go-fiber/adapter/repository/sql/db"
	"realworld-go-fiber/core/util"
)

func main() {
	config, err := util.LoadConfig(".env")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load env config: %s", err)
		os.Exit(1)
	}

	logger := logger.NewLogger(config)

	database, err := db.New(config, logger)
	if err != nil {
		log.Fatal(err)
	}
	logger.Info().Msg("Connect Database successfully")

	server := restful.NewServer(config, logger, database.DB())
	log.Fatal(server.Start())
}
