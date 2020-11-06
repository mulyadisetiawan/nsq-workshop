package nsq_publisher

import "github.com/bitly/go-nsq"

var Producer *nsq.Producer

func InitNSQ() {
	config := nsq.NewConfig()
	Producer, _ = nsq.NewProducer("172.18.59.254:4150", config)
}
