package main

import (
	"fmt"

	handlerhttp "github.com/sharring_session/nsq/http"
	"github.com/sharring_session/nsq/nsq"
)

func main() {
	nsq.InitProducer()
	nsq.InitConsumer()
	fmt.Println("RUNNING")
	handlerhttp.HandleRequests()
}
