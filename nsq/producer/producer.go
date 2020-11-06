package nsqproducer

import (
	"log"

	jsoniter "github.com/json-iterator/go"
	nsq "github.com/nsqio/go-nsq"
)

var producer *nsq.Producer

func Init() {
	newProducer, err := nsq.NewProducer("localhost:4150", nsq.NewConfig())
	if err != nil {
		log.Fatal(err)
	}

	producer = newProducer
}

func Publish(topic string, data interface{}) (err error) {
	payload, err := jsoniter.Marshal(data)
	if err != nil {
		return
	}

	if producer == nil {
		return
	}

	return producer.Publish(topic, payload)
}
