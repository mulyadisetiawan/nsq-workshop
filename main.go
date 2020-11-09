package main

import (
	"fmt"

	nsqu "github.com/sharring_session/nsq/nsq"

	handlerhttp "github.com/sharring_session/nsq/http"
)

func main() {
	fmt.Println("RUNNING")
	nsqu.InitProducer()
	go nsqu.InitConsumer()
	handlerhttp.HandleRequests()
	nsqu.Stop()
}
