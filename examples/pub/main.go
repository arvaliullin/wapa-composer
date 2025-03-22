package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Connection error %v", err)
	}
	defer nc.Close()

	nc.Publish("hello", []byte("Message"))

}
