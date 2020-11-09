package nsqu

import (
	"log"
	"strconv"
	"sync"

	"github.com/bitly/go-nsq"
	"github.com/sharring_session/nsq/api/ovo"
)

func InitConsumer() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer("myapp__give_ovo", "give_ovo", config)
	q.AddHandler(nsq.HandlerFunc(GiveOVOConsumer))

	err := q.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Panic("Could not connect")
	}

	wg.Wait()
}

func GiveOVOConsumer(message *nsq.Message) error {
	userID, err := strconv.Atoi(string(message.Body))
	if err != nil {
		return err
	}

	err = ovo.GiveBenefit(userID)
	if err != nil {
		return err
	}

	return err
}
