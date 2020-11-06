package nsqhandler

import (
	"encoding/json"
	"log"

	"github.com/nsqio/go-nsq"
)

const (
	defaultProducerPrefix = "bgp_tech_curriculum"
)

type (
	ProducerConfig struct {
		NsqdAddress string
	}
)

var producer *nsq.Producer

func NewProducer(cfg ProducerConfig) *nsq.Producer {
	producer, err := nsq.NewProducer(cfg.NsqdAddress, nsq.NewConfig())
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
	topic = defaultProducerPrefix + topic

	return producer.Publish(topic, payload)
}
