package main

import (
	"log"

	"github.com/nats-io/nats.go"
)


func main() {
	// Connect to NATS server
	nc, err := nats.Connect("nats://127.0.0.1:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Subscribe to a subject
	_, err = nc.Subscribe("foo", func(msg *nats.Msg) {
		log.Printf("Received message: %s", string(msg.Data))
	})
	if err != nil {
		log.Fatal(err)
	}
}
