package nsqhandler

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	defaultConsumerMaxAttempts = 10
	defaultConsumerMaxInFlight = 100
)

func InitConsumer() {
	// initiate consumer
	cfg := ConsumerConfig{
		Channel:       "give-ovo",
		LookupAddress: "127.0.0.1:4161",
		Topic:         "give-ovo",
		MaxAttempts:   defaultConsumerMaxAttempts,
		MaxInFlight:   defaultConsumerMaxInFlight,
		Handler:       GiveOvo,
	}
	consumer := NewConsumer(cfg)

	go RunConsumer(consumer)
}

func InitProducer() {
	// initiate producer
	prodConf := ProducerConfig{
		NsqdAddress: "127.0.0.1:4150", // TODO: update to nsqd address
	}
	producer = NewProducer(prodConf)
}

func RunConsumer(c Consumer) {
	// run consumer
	c.Run()

	// keep app alive until terminated
	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	select {
	case <-term:
		log.Println("Application terminated")
	}
}
