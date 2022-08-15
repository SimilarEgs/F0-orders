package postgresql

import (
	"fmt"
	"log"

	"github.com/SimilarEgs/L0-orders/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	ordersTable   = "orders"
	deliveryTable = "delivery"
	paymentTable  = "payment"
	itemsTable    = "items"
)

type DB struct {
	Con *sqlx.DB
}

func (db *DB) Init(cfg *config.Config) {
	dbSource := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.PostgresSQL.PostgresqlHost,
		cfg.PostgresSQL.PostgresqlPort,
		cfg.PostgresSQL.PostgresqlUser,
		cfg.PostgresSQL.PostgresqlDBName,
		cfg.PostgresSQL.PostgresqlPassword,
		cfg.PostgresSQL.PostgresqlSslmode)

	var err error

	db.Con, err = sqlx.Connect("postgres", dbSource)
	if err != nil {
		log.Fatalf("[Error] occured while connecting to the db: %v", err)
	}

}
