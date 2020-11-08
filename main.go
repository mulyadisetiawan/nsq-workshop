package main

import (
	"fmt"

	handlerhttp "github.com/sharring_session/nsq/http"
	"github.com/sharring_session/nsq/nsq"
)

func main() {
	fmt.Println("RUNNING")
	handlerhttp.HandleRequests()
	nsq.InitConsumer()
	nsq.InitProducer()
}
