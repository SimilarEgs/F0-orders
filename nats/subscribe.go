package nats

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/SimilarEgs/L0-orders/config"
	"github.com/nats-io/stan.go"
)

type Mock struct {
	Field  string
	Number int
}

func Subscriber(cfg *config.Config) {

	data := Mock{}

	con, err := NatsConnect(cfg, cfg.Nats.SubID)
	if err != nil {
		log.Fatalf("[Error] failed nats connection: %v\n", err)
	}

	log.Println("[Info] connection with nats streaming is established")

	_, err = con.Subscribe(cfg.Nats.Subject, func(msg *stan.Msg) {

		if err := msg.Ack(); err != nil {
			log.Printf("[Error] occurred while receiving message: %v\n", err)
		}

		if err := json.Unmarshal(msg.Data, &data); err != nil {
			log.Printf("[Error] occured while unmurshaling msg dta: %v\n", err.Error())
		}

		fmt.Println("Recived data:", data.Field, data.Number)

	}, stan.SetManualAckMode(), stan.AckWait(time.Duration(30)*time.Second), stan.DeliverAllAvailable(), stan.MaxInflight(10), stan.DurableName(cfg.Nats.NatsDurable))

	if err != nil {
		log.Printf("[Error] sub: %v\n", err)
	}
	log.Printf("[Info] client %s was subscribed to %s\n", cfg.Nats.SubID, cfg.Nats.Subject)

}
