package nsqconsumer

import (
	"log"

	jsoniter "github.com/json-iterator/go"
	nsq "github.com/nsqio/go-nsq"

	"github.com/natasharamdani/nsq-workshop/api/ovo"
	handlerhttp "github.com/natasharamdani/nsq-workshop/http"
)

func Init() {
	consumer, err := nsq.NewConsumer("topic_give_benefit", "channel_give_benefit", nsq.NewConfig())
	if err != nil {
		log.Fatal(err)
	}

	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		var payload handlerhttp.GiveBenefitPayload
		err = jsoniter.Unmarshal(message.Body, &payload)
		if err != nil {
			message.Finish()
			return err
		}

		err = ovo.GiveBenefit(payload.UserID)
		if err != nil {
			message.Finish()
			return err
		}

		return nil
	}))

	err = consumer.ConnectToNSQLookupd("localhost:4161")
	if err != nil {
		log.Fatal(err)
	}
}
