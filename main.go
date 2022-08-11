package main

import (
	"log"

	"github.com/SimilarEgs/F0-orders/config"
)

func main() {

	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

}
