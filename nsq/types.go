package nsq

import (
	gonsq "github.com/nsqio/go-nsq"
)

type ProducerConfig struct {
	NsqdAddress string
}

type ConsumerConfig struct {
	Topic         string
	Channel       string
	LookupAddress string
	MaxAttempts   uint16
	MaxInFlight   int
	Handler       gonsq.HandlerFunc
}

type Consumer struct {
	Consumer      *gonsq.Consumer
	LookupAddress string
	Handler       gonsq.HandlerFunc
}
