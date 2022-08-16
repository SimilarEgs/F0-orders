package nats

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/SimilarEgs/L0-orders/config"
	"github.com/SimilarEgs/L0-orders/internal/models"
	"github.com/SimilarEgs/L0-orders/pkg/cache"
	"github.com/SimilarEgs/L0-orders/pkg/postgresql"
	"github.com/nats-io/stan.go"
)

const (
	cacheDuration = 3 * time.Hour
	cacheCleanUp  = 6 * time.Hour
)

var AppCache = cache.New(cacheDuration, cacheCleanUp)

func Subscriber(cfg *config.Config) (stan.Subscription, error) {

	order := models.Order{}
	var db postgresql.DB

	con, err := NatsConnect(cfg, cfg.Nats.SubID)
	if err != nil {
		log.Fatalf("[Error] failed nats connection: %v\n", err)
	}

	log.Println("[Info] connection with nats streaming is established")

	db.Init(cfg)
	log.Println("[Info] successfully connected to the db")

	sub, err := con.Subscribe(cfg.Nats.Subject, func(msg *stan.Msg) {

		if err := msg.Ack(); err != nil {
			log.Printf("[Error] occurred while receiving message: %v\n", err)
		}

		if err := json.Unmarshal(msg.Data, &order); err != nil {
			msg := fmt.Sprintf("[Error] order ID - «%s»\n", order.OrderUID)
			log.Print(msg)
			log.Printf("[Error] occured while unmurshaling payload data: %v\n\n", err.Error())
			return
		}

		if err := order.ValidateOrder(); err != nil {
			log.Printf("[Error] occurred while validating message data: %v\n", err)
			return

		}

		db.Insert(&order)
		AppCache.Set(order.OrderUID, order, cacheDuration)

		log.Printf("[Info] order - «%s» was successfully inserted into the DB\n", order.OrderUID)

	}, stan.SetManualAckMode(), stan.AckWait(time.Duration(30)*time.Second), stan.DeliverAllAvailable(), stan.MaxInflight(10), stan.DurableName(cfg.Nats.NatsDurable))

	if err != nil {
		log.Printf("[Error] sub: %v\n", err)
		return nil, err
	}

	log.Printf("[Info] client «%s» was subscribed to «%s» subject\n", cfg.Nats.SubID, cfg.Nats.Subject)

	return sub, err
}
