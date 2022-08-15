package main

import (
	"log"

	"github.com/SimilarEgs/L0-orders/config"
	"github.com/SimilarEgs/L0-orders/internal/server"
	"github.com/SimilarEgs/L0-orders/nats"
)

func main() {

	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	sub, err := nats.Subscriber(cfg)
	if err != nil {
		log.Println(err)
	}
	defer sub.Unsubscribe()
	defer sub.Close()

	srv := new(server.Server)

	if err := srv.RunServer(server.ServeRoutes(), cfg); err != nil {
		log.Fatalf("[Error] failed to start server: %s", err.Error())
	}

}
