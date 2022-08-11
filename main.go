package main

import (
	"log"

	"github.com/SimilarEgs/F0-orders/config"
	"github.com/SimilarEgs/F0-orders/internal/server"
)

func main() {

	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	srv := new(server.Server)

	if err := srv.RunServer(server.ServeRoutes(), cfg); err != nil {
		log.Fatalf("[Error] failed to start server: %s", err.Error())
	}

}
