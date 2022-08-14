package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/SimilarEgs/L0-orders/config"
	"github.com/SimilarEgs/L0-orders/nats"
)

type Mock struct {
	Field  string
	Number int
}

func main() {
	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	con, err := nats.NatsConnect(cfg, cfg.Nats.NatsPubID)
	defer con.Close()

	if err != nil {
		log.Printf("[Error] occurred while connecting to the nats: %v", err)
	}

	for i := 0; ; i++ {

		data := Mock{"i - ", i}
		fmt.Println(data)
		d, err := json.Marshal(&data)

		err = con.Publish(cfg.Nats.Subject, d)
		if err != nil {
			log.Println("[Error] occurred while publishing the message ")
		}

		log.Println("[Info] message was successfully sent")
		time.Sleep(time.Second / 2)

	}
}
