package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalln(err)
	}
	defer nc.Close()

	_, err = nc.Subscribe("hello", func(msg *nats.Msg) {
		log.Printf("Received a message: %s\n", string(msg.Data))
	})

	if err != nil {
		log.Fatalln(err)
	}

	select {}
}
