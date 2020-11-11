package main

import (
	"fmt"

	handlerhttp "github.com/zulfahmi14/nsq-workshop/http"
	publisher "github.com/zulfahmi14/nsq-workshop/nsq/publisher"
)

func main() {
	fmt.Println("RUNNING")
	handlerhttp.HandleRequests()
	publisher.InitProducer()
}
