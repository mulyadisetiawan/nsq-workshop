package consumer

import (
	"fmt"

	"github.com/nsqio/go-nsq"
	config "github.com/sharring_session/nsq/nsq-workshop/config"
	handler "github.com/sharring_session/nsq/nsq-workshop/handler"
)

func ConsumeGiveBenefit() error {
	cfg := nsq.NewConfig()

	c, err := nsq.NewConsumer(config.NSQTopic, config.NSQChannel, cfg)
	if err != nil {
		fmt.Println(err)
		return err
	}

	c.AddHandler(nsq.HandlerFunc(handler.HandlerGiveBenefitOVO))

	err = c.ConnectToNSQD(config.NSQLookupdURL)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
