package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/bitly/go-nsq"
	"github.com/sharring_session/nsq/api/ovo"
	handlerhttp "github.com/sharring_session/nsq/http"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer(handlerhttp.NSQ_TOPIC_PUBLISH_GIVE_BENEFIT, "consumer_give_ovo", config)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		userID, err := strconv.Atoi(string(message.Body[:]))
		if err != nil {
			fmt.Println(err)
		}
		ovo.GiveBenefit(userID)
		return nil
	}))

	err := q.ConnectToNSQLookupd("172.18.59.254:4161")
	if err != nil {
		log.Panic("Could not connect")
	}
	wg.Wait()
}
