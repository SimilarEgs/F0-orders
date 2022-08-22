package main

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/SimilarEgs/L0-orders/config"
	"github.com/SimilarEgs/L0-orders/nats"
)

const (
	path = "./publisher/json/"
)


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

	data, err := ReadFiles(path)
	if err != nil {
		log.Fatalf("[Error] occurred while reading files: %v\n", err)
	}

	for i, v := range *data {

		err = con.Publish(cfg.Nats.Subject, v)

		if err != nil {
			log.Printf("[Error] occurred while publishing the message: %v", err)
		}
		log.Printf("[Info] message %d was successfully sent\n", i+1)

		time.Sleep(time.Second * 1)
	}

}

func ReadFiles(dir string) (*[][]byte, error) {

	res := new([][]byte)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if !f.IsDir() {
			data, err := ioutil.ReadFile(path + f.Name())
			if err != nil {
				return nil, err
			}
			*res = append(*res, data)
		}

	}

	return res, nil
}
