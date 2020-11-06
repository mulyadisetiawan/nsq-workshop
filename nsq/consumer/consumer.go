package consumer

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
	gonsq "github.com/nsqio/go-nsq"
	"github.com/sharing_session/nsq-workshop/api/ovo"
	"github.com/sharing_session/nsq-workshop/nsq"
)

func Consume() {
	decodeConfig := gonsq.NewConfig()
	c, err := gonsq.NewConsumer(nsq.Topic_give_ovo, "give_ovo", decodeConfig)
	if err != nil {
		fmt.Println("error creating consumer")
	}

	c.AddHandler(gonsq.HandlerFunc(func(message *gonsq.Message) error {
		var userID int
		err = jsoniter.Unmarshal(message.Body, &userID)
		if err != nil {
			message.Finish()
			return err
		}

		return ovo.GiveBenefit(userID)
	}))

	err = c.ConnectToNSQLookupd(nsq.Nsqlookupd_address)
	if err != nil {
		fmt.Println("could not connect to nsq lookupd")
	}
}
