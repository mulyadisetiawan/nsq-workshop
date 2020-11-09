package nsqu

import (
	"fmt"

	"github.com/bitly/go-nsq"
)

var w *nsq.Producer

func InitProducer() {
	config := nsq.NewConfig()
	w, _ = nsq.NewProducer("127.0.0.1:4150", config)
}

func Stop() {
	w.Stop()
}

func Publish(topic, message string) error {
	fmt.Printf("%+v\n", w)

	err := w.Publish(topic, []byte(message))

	return err
}
