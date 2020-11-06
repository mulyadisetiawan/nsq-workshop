package producer

import (
	"sync"

	"github.com/bitly/go-nsq"
	"github.com/sharring_session/nsq/config"
)

var producer *nsq.Producer
var once sync.Once

func Publish(msg []byte) error {
	once.Do(func() {
		cfg := nsq.NewConfig()
		_producer, err := nsq.NewProducer(config.NSQLOOKUPD_URL, cfg)
		if err != nil {
			panic(err)
		}
		producer = _producer
	})

	if err := producer.Publish(config.NSQ_TOPIC, msg); err != nil {
		return err
	}

	return nil
}
