package publisher

import (
	jsoniter "github.com/json-iterator/go"
	gonsq "github.com/nsqio/go-nsq"
	"github.com/sharing_session/nsq-workshop/nsq"
)

func Publish(topic string, data interface{}) error {
	config := gonsq.NewConfig()
	p, _ := gonsq.NewProducer(nsq.Nsqd_address, config)

	payload, err := jsoniter.Marshal(data)
	if err != nil {
		return err
	}

	err = p.Publish(topic, payload)
	if err != nil {
		return err
	}

	return nil
}
