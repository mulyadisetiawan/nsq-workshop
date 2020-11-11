package main

import (
	"fmt"

	consumer "github.com/zulfahmi14/nsq-workshop/nsq/consumer"
)

func main() {
	fmt.Println("RUNNING")
	consumer.InitConsumer()
}
