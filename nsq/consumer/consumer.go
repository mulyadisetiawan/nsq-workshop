package consumer

import (
	"errors"
	"sync"

	jsoniter "github.com/json-iterator/go"
	"github.com/nsqio/go-nsq"
)

func InitConsumer() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	q, _ := nsq.NewConsumer("Topic_GiveOvo_Benefit", "give_benefit", nsq.NewConfig())
	q.AddHandler(nsq.HandlerFunc(GiveBenefitOvo))

	err := q.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		panic("Could not connect")
	}

	wg.Wait()
}

func GiveBenefitOvo(msg *nsq.Message) (err error) {
	payload := OvoPayload{}
	err = jsoniter.Unmarshal(msg.Body, &payload)
	if err != nil {
		msg.Finish()
		return err
	}

	if payload.UserID <= 0 {
		msg.Finish()
		return errors.New("userID  = 0")
	}

	msg.Finish()
	return nil
}
