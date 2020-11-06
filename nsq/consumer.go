package main

import (
	"fmt"
	"strconv"
	"time"

	nsq "github.com/nsqio/go-nsq"
)

func giveOVO() {
	decodeConfig := nsq.NewConfig()
	c, err := nsq.NewConsumer("give_ovo", "give", decodeConfig)
	if err != nil {
		fmt.Println("Could not create consumer")
	}

	c.AddHandler(nsq.HandlerFunc(func(msg *nsq.Message) error {
		fmt.Println("NSQ msg received:")
		fmt.Println(string(msg.Body))

		userID, err := strconv.Atoi(string(msg.Body))
		if err != nil {
			if msg.Attempts > 10 {
				msg.Finish()
			} else if msg.Attempts > 3 {
				msg.RequeueWithoutBackoff(time.Minute * 20) // can handle down time 140 minutes
			} else {
				msg.RequeueWithoutBackoff(time.Second * 10)
			}
			return err
		}

		if userID == 0 {
			msg.Finish()
			return fmt.Errorf("user id is empty. userID:%d", userID)
		}

		msg.Finish()
		return nil
	}))

	// connect by nsqlookupd
	err = c.ConnectToNSQLookupd("127.0.0.1:4161")
	// connect directly to nsqd
	// err = c.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		fmt.Println("Could not connect")
	}
}
