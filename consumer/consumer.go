package consumer

import (
	"fmt"

	"github.com/bitly/go-nsq"
	"github.com/sharring_session/nsq/config"
)

func Consume() {
	go func() {
		cfg := nsq.NewConfig()

		q, err := nsq.NewConsumer(config.NSQ_TOPIC, config.NSQ_CHANNEL, cfg)
		if err != nil {
			panic(err)
		}

		q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
			fmt.Println("[OVO]", string(message.Body))
			return nil
		}))

		if err := q.ConnectToNSQD(config.NSQLOOKUPD_URL); err != nil {
			panic(err)
		}

		c := make(chan struct{})
		<-c
	}()
}
