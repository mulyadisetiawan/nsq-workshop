package nsq

import (
	"log"

	"github.com/bitly/go-nsq"
	jsoniter "github.com/json-iterator/go"
)

var producer *nsq.Producer

func InitProducer() {
	var err error
	config := nsq.NewConfig()
	producer, err = nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal("error when connecting to nsqd")
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
