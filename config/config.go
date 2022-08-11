package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	HTTP        HTTP
	Nats        Nats
	PostgresSQL PostgresSQL
}

type HTTP struct {
	Port              string
	Timeout           time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	MaxConnectionIdle time.Duration
	MaxConnectionAge  time.Duration
}

type Nats struct {
	URL       string
	ClusterID string
	ClientID  string
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

	return &c, nil
}
