package nsq

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/bitly/go-nsq"
	"github.com/sharring_session/nsq/api/ovo"
)

func InitConsumer() {
	decodeConfig := nsq.NewConfig()
	c, err := nsq.NewConsumer(GiveOVOTopic, GiveOVOChannel, decodeConfig)
	if err != nil {
		log.Fatal("error when creating consumer ")
	}

	c.AddHandler(nsq.HandlerFunc(GiveBenefitOvo))

	err = c.ConnectToNSQLookupd("127.0.0.1:4161")
	if err != nil {
		log.Fatal("error when connecting to nsqlookupd")
	}

	if err := c.ConnectToNSQD("127.0.0.1:4150"); err != nil {
		log.Fatal("error when connecting to nsqd")
	}
}

func GiveBenefitOvo(message *nsq.Message) error {
	var userID int

	err := json.Unmarshal(message.Body, &userID)
	if err != nil {
		return err
	}

	err = ovo.GiveBenefit(userID)
	if err != nil {
		fmt.Println("error when give benefit")
		return err
	}

	message.Finish()
	return nil
}
