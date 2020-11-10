package mnsq

import (
	"github.com/nsqio/go-nsq"
	"sync"
)

func InitConsumer() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	q, _ := nsq.NewConsumer(TopicGiveOvoBenefit, "give_benefit", nsq.NewConfig())
	q.AddHandler(nsq.HandlerFunc(GiveBenefitOvo))

	err := q.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		panic("Could not connect")
	}

	wg.Wait()
}
