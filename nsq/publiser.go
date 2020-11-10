package mnsq

import (
	jsoniter "github.com/json-iterator/go"

	nsq "github.com/nsqio/go-nsq"
)

var producer *nsq.Producer

func InitProducer() {
	var err error
	config := nsq.NewConfig()
	producer, err = nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		panic("Failed to create NSQ producer")
	}
}

func Publish(topic string, data interface{}) (err error) {
	var payload []byte
	payload, err = jsoniter.Marshal(data)
	if err != nil {
		return
	}

	if producer == nil {
		return
	}

	return producer.Publish(topic, payload)
}
