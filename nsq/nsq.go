package nsq

import (
	"encoding/json"
	"fmt"
	"log"

	nsq "github.com/nsqio/go-nsq"
	"github.com/sharing_session/nsq/nsq-workshop/api/ovo"
)

var producer *nsq.Producer

func InitNsq() {
	var err error
	config := nsq.NewConfig()
	producer, err = nsq.NewProducer(NSQDUrl, config)
	if err != nil {
		log.Fatal(err)
	}
}

func Publish(topic string, message interface{}) error {
	messageByte, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = producer.Publish(topic, messageByte)
	if err != nil {
		fmt.Println("Error publish ", err)
		return err
	}

	return nil
}

func InitConsumer() {
	//create consumer
	decodeConfig := nsq.NewConfig()
	c, err := nsq.NewConsumer(GiveOVOTopic, GiveOVOChannel, decodeConfig)
	if err != nil {
		log.Fatal("error when creating consumer ")
	}

	c.AddHandler(nsq.HandlerFunc(HandlerGiveOVO))

	err = c.ConnectToNSQLookupd(NSQLookupdURL)
	if err != nil {
		log.Fatal("error when connecting to nsqlookupd")
	}

	if err := c.ConnectToNSQD(NSQDUrl); err != nil {
		log.Fatal("error when connecting to nsqd")
	}
}

func HandlerGiveOVO(message *nsq.Message) error {
	fmt.Println("CONSUMED")

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
