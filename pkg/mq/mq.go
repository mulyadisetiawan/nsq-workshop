package mq

import (
	"github.com/bitly/go-nsq"
)

type (
	Config struct {
		ListenAddress []string `yaml:"listenaddress"`
	}

	SubscriberOptions struct {
		Handler nsq.Handler
	}

	Consumer struct {
		topic   string
		channel string
		handler nsq.Handler
	}

	Subscribers map[string]map[string]SubscriberOptions
)

func NewNSQLHandler(h func(*nsq.Message) error) nsq.HandlerFunc {
	return nsq.HandlerFunc(h)
}
