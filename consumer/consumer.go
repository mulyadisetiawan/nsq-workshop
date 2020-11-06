package consumer

import (
	"fmt"

	"github.com/nsqio/go-nsq"
	nsqku "github.com/sharring_session/nsq-workshop/nsq"
	"github.com/sharring_session/nsq-workshop/util"
)

func Consumer() {
	decodeConfig := nsq.NewConfig()
	c, err := nsq.NewConsumer(util.TopicGiveBenefit, util.ChannelGiveBenefit, decodeConfig)
	if err != nil {
		fmt.Println(err)
	}

	c.AddHandler(nsq.HandlerFunc(nsqku.GiveBenefitNSQ))

	err = c.ConnectToNSQLookupd(util.NSQLookupURL)
	if err != nil {
		fmt.Println("any error")
	}
}
