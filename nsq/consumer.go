package nsq

import (
	"log"

	gonsq "github.com/nsqio/go-nsq"
)

func NewConsumer(config ConsumerConfig) Consumer {
	nsqConf := gonsq.NewConfig()
	nsqConf.MaxAttempts = config.MaxAttempts
	nsqConf.MaxInFlight = config.MaxInFlight

	topic := NSQPrefix + config.Topic
	c, err := gonsq.NewConsumer(topic, config.Channel, gonsq.NewConfig())
	if err != nil {
		log.Fatal(err)
	}
	return Consumer{
		Consumer:      c,
		LookupAddress: config.LookupAddress,
		Handler:       config.Handler,
	}
}

func (c *Consumer) Run() {
	c.Consumer.AddHandler(c.Handler)
	err := c.Consumer.ConnectToNSQLookupd(c.LookupAddress)
	if err != nil {
		log.Fatal(err)
	}
}
