package nsq

import (
	"encoding/json"
	"log"

	gonsq "github.com/nsqio/go-nsq"
)

var producer *gonsq.Producer

func NewProducer(config ProducerConfig) *gonsq.Producer {
	producer, err := gonsq.NewProducer(config.NsqdAddress, gonsq.NewConfig())
	if err != nil {
		log.Fatal(err)
	}

	return producer
}

func Publish(topic string, msg interface{}) error {
	payload, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	topic = NSQPrefix + topic

	return producer.Publish(topic, payload)
}
