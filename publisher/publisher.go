package publisher

import (
	"fmt"

	"github.com/nsqio/go-nsq"
	"github.com/sharring_session/nsq-workshop/util"
)

//Publish is function for publish
func Publish(topic string, data []byte) (err error) {
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer(util.NSQLookupdURL, config)

	err = w.Publish(topic, data)
	if err != nil {
		fmt.Println("error from func publish :", err)
		return err
	}

	return nil
}
