package nsq

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func InitConsumer() {
	// initiate consumer
	cfg := ConsumerConfig{
		Channel:       ChannelGiveOVO,
		LookupAddress: "127.0.0.1:4161",
		Topic:         TopicGiveOVO,
		MaxAttempts:   ConsumerMaxAttempts,
		MaxInFlight:   ConsumerMaxInFlight,
		Handler:       GiveOvo,
	}
	consumer := NewConsumer(cfg)

	go RunConsumer(consumer)
}

func InitProducer() {
	tempProducer := ProducerConfig{
		NsqdAddress: "127.0.0.1:4150",
	}
	producer = NewProducer(tempProducer)
}

func RunConsumer(consumer Consumer) {
	consumer.Run()

	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	select {
	case <-term:
		log.Println("gracefully terminated")
	}
}
