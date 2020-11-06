package nsqmodule

import (
	"log"

	nsq "github.com/nsqio/go-nsq"
)

func NewConsumer(cfg ConsumerConfig) Consumer {
	nsqConf := nsq.NewConfig()
	nsqConf.MaxAttempts = cfg.MaxAttempts
	nsqConf.MaxInFlight = cfg.MaxInFlight

	topic := cfg.Topic
	c, err := nsq.NewConsumer(topic, cfg.Channel, nsq.NewConfig())
	if err != nil {
		log.Fatal(err)
	}
	return Consumer{
		cons:          c,
		lookupAddress: cfg.LookupAddress,
		handler:       cfg.Handler,
	}
}

func (c *Consumer) Run() {
	c.cons.AddHandler(c.handler)
	err := c.cons.ConnectToNSQLookupd(c.lookupAddress)
	if err != nil {
		log.Fatal(err)
	}
}
