package nats

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/SimilarEgs/L0-orders/config"
	"github.com/SimilarEgs/L0-orders/internal/models"
	"github.com/nats-io/stan.go"
)

func Subscriber(cfg *config.Config) (stan.Subscription, error) {

	order := models.Order{}

	con, err := NatsConnect(cfg, cfg.Nats.SubID)
	if err != nil {
		log.Fatalf("[Error] failed nats connection: %v\n", err)
	}

	log.Println("[Info] connection with nats streaming is established")

	sub, err := con.Subscribe(cfg.Nats.Subject, func(msg *stan.Msg) {

		if err := msg.Ack(); err != nil {
			log.Printf("[Error] occurred while receiving message: %v\n", err)
		}

		if err := json.Unmarshal(msg.Data, &order); err != nil {
			msg := fmt.Sprintf("[Error] occured while unmurshaling msg dta: %v\n", err.Error())
			log.Println(msg)
			return

		}

		fmt.Println("Recived data:", string(msg.Data))

	}, stan.SetManualAckMode(), stan.AckWait(time.Duration(30)*time.Second), stan.DeliverAllAvailable(), stan.MaxInflight(10), stan.DurableName(cfg.Nats.NatsDurable))

	if err != nil {
		log.Printf("[Error] sub: %v\n", err)
		return nil, err
	}

	log.Printf("[Info] client %s was subscribed to %s\n", cfg.Nats.SubID, cfg.Nats.Subject)

	return sub, err
}
