package main

import (
	"fmt"
	"log"
	"os"
	"realworld-go-fiber/adapter/handler/restful"
	"realworld-go-fiber/adapter/logger"
	"realworld-go-fiber/adapter/repository/sql"
	"realworld-go-fiber/adapter/repository/sql/db"
	"realworld-go-fiber/core/uc"
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

	repo := sql.NewSQLRepository(database.DB(), logger)
	uc, err := uc.NewUsecase(config, repo, logger)
	if err != nil {
		log.Fatal(err)
	}

	server := restful.NewServer(config, logger, uc)
	log.Fatal(server.Start())
}
