package server

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	nsqmodule "github.com/mulyadisetiawan/nsq-workshop/nsq"
	"github.com/nsqio/go-nsq"
)

var ProducerClient nsqmodule.Producer
var ConsumerClient nsqmodule.Consumer

func InitProducer() {
	prodConf := nsqmodule.ProducerConfig{
		NsqdAddress: NSQdAddress,
	}
	ProducerClient = nsqmodule.NewProducer(prodConf)
}

func InitConsumer(function nsq.HandlerFunc) {
	cfg := nsqmodule.ConsumerConfig{
		Channel:       NSQPrefix + NSQChannelGiveBenefit,
		LookupAddress: NSQLookupdAddress,
		Topic:         NSQPrefix + NSQTopicGiveBenefit,
		MaxAttempts:   nsqmodule.DefaultConsumerMaxAttempts,
		MaxInFlight:   nsqmodule.DefaultConsumerMaxInFlight,
		Handler:       function,
	}
	ConsumerClient := nsqmodule.NewConsumer(cfg)

	ConsumerClient.Run()

	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	select {
	case <-term:
		log.Println("Application terminated")
	}
}
