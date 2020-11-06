package server

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/nsqio/go-nsq"
)

var (
	Producer *nsq.Producer
)

const (
	ProducerHost = "127.0.0.1:4150"
	NsqdAddr     = "127.0.0.1:4150"
)

func GetProducer() *nsq.Producer {

	var err error

	cfg := nsq.NewConfig()

	if Producer != nil {
		return Producer
	}

	Producer, err = nsq.NewProducer(ProducerHost, cfg)
	if err != nil {
		log.Fatal("error init producer")
	}

	return Producer
}

func InitConsumer() {
	defaultCfg := nsq.NewConfig()

	RegisterConsumerGiveOVO(defaultCfg)
}

func RegisterConsumerGiveOVO(cfg *nsq.Config) {

	c, err := nsq.NewConsumer("topic", "channel", cfg)
	if err != nil {
		log.Fatal("error create consumer RegisterConsumerGiveOVO", err)
	}

	c.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {

		fmt.Println("INCOMING MESSAGE RegisterConsumerGiveOVO: ", message)

		userIDStr := string(message.Body)

		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			return err
		}

		if userID == 0 {
			return errors.New("invalid user id")
		}

		return nil
	}))

	if err := c.ConnectToNSQD(NsqdAddr); err != nil {
		log.Fatal("failed connect to nsqd")
	}
}
