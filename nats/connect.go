package nats

import (
	"log"

	"github.com/SimilarEgs/L0-orders/config"
	"github.com/nats-io/stan.go"
)

func NatsConnect(cfg *config.Config, natsClientId string) (stan.Conn, error) {
	return stan.Connect(
		cfg.Nats.ClusterID,
		natsClientId,
		stan.NatsURL(stan.DefaultNatsURL),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("[Error] connection lost, reason: %v\n", reason)
		}),
	)
}
