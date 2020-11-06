package publisher

import (
	"fmt"

	"github.com/nsqio/go-nsq"
	"github.com/sharring_session/nsq/nsq-workshop/config"
)

func PublishGiveBenefit(message []byte) error {
	cfg := nsq.NewConfig()

	p, _ := nsq.NewProducer(config.NSQLookupdURL, cfg)

	err := p.Publish(config.NSQTopic, message)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
