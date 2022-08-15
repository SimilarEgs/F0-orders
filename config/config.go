package config

import (
	"log"
	"os"
	"time"

	"github.com/SimilarEgs/L0-orders/pkg/constants"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	HTTP        HTTP
	Nats        Nats
	PostgresSQL PostgresSQL
}

type HTTP struct {
	Port         string
	Timeout      time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Nats struct {
	URL         string
	ClusterID   string
	SubID       string
	NatsPubID   string
	NatsDurable string
	Subject     string
}

type PostgresSQL struct {
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDBName   string
	PostgresqlSslmode  string
}

func readConfig() error {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.SetConfigName("config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("[Error] .env file didn't load: %s", err.Error())
	}

	return nil
}

func ParseConfig() (*Config, error) {
	if err := readConfig(); err != nil {
		return nil, err
	}

	var c Config

	err := viper.Unmarshal(&c)
	if err != nil {
		log.Printf("[Error] unable to decode into struct: %v\n", err)
		return nil, err
	}

	natsUrl := os.Getenv(constants.NATS_URL)
	if natsUrl != "" {
		c.Nats.URL = natsUrl
	}

	natsSubID := os.Getenv(constants.NATS_SUB_ID)
	if natsSubID != "" {
		c.Nats.SubID = natsSubID
	}

	natsDurable := os.Getenv(constants.NATS_DURABLE)
	if natsDurable != "" {
		c.Nats.NatsDurable = natsDurable
	}

	natsPubID := os.Getenv(constants.NATS_PUB_ID)
	if natsPubID != "" {
		c.Nats.NatsPubID = natsPubID
	}

	natsClusterID := os.Getenv(constants.NATS_CLUSTER_ID)
	if natsClusterID != "" {
		c.Nats.ClusterID = natsClusterID
	}

	nutsSubject := os.Getenv(constants.NATS_SUBJECT)
	if nutsSubject != "" {
		c.Nats.Subject = nutsSubject
	}

	postgresqlHost := os.Getenv(constants.POSTGRES_HOST)
	if postgresqlHost != "" {
		c.PostgresSQL.PostgresqlHost = postgresqlHost
	}

	postgresqlPort := os.Getenv(constants.POSTGRES_PORT)
	if postgresqlPort != "" {
		c.PostgresSQL.PostgresqlPort = postgresqlPort
	}

	postgresqlUser := os.Getenv(constants.POSTGRES_USER)
	if postgresqlUser != "" {
		c.PostgresSQL.PostgresqlUser = postgresqlUser
	}

	postgresqlPassword := os.Getenv(constants.POSTGRES_PASSWORD)
	if postgresqlPassword != "" {
		c.PostgresSQL.PostgresqlPassword = postgresqlPassword
	}

	postgresqlDBName := os.Getenv(constants.POSTGRES_DBNAME)
	if postgresqlDBName != "" {
		c.PostgresSQL.PostgresqlDBName = postgresqlDBName
	}

	postgresqlSslmode := os.Getenv(constants.POSTGRES_SSLMODE)
	if postgresqlSslmode != "" {
		c.PostgresSQL.PostgresqlSslmode = postgresqlSslmode
	}

	return &c, nil
}

// func getEnv() error {
// 	err := os.Setenv("NATS_URL", "localhost:4222")
// 	if err != nil {
// 		return err
// 	}
// 	err = os.Setenv("NATS_CLUSTER_ID", "test-cluster")
// 	if err != nil {
// 		return err
// 	}
// 	err = os.Setenv("NATS_SUB_ID", "test-sub")
// 	if err != nil {
// 		return err
// 	}
// 	err = os.Setenv("NATS_PUB_ID", "test-pub")
// 	if err != nil {
// 		return err
// 	}
// 	err = os.Setenv("NATS_DURABLE_ID", "test-durable-sub")
// 	if err != nil {
// 		return err
// 	}
// 	err = os.Setenv("NATS_SUBJECT", "orders")
// 	if err != nil {
// 		return err
// 	}

// 	// POSTGRES_HOST
// 	// POSTGRES_PORT
// 	// POSTGRES_USER
// 	// POSTGRES_PASSWORD
// 	// POSTGRES_DBNAME
// 	// POSTGRES_SSLMODE
// 	return nil
// }
