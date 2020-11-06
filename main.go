package main

import (
	"fmt"

	handlerhttp "github.com/sharring_session/nsq/http"
	nsqhandler "github.com/sharring_session/nsq/nsq"
)

func main() {
	fmt.Println("RUNNING")
	handlerhttp.HandleRequests()
	nsqhandler.InitConsumer()
	nsqhandler.InitProducer()
}
