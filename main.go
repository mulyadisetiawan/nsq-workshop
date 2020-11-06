package main

import (
	"fmt"

	handlerhttp "github.com/sharing_session/nsq/nsq-workshop/http"
	"github.com/sharing_session/nsq/nsq-workshop/nsq"
)

func main() {
	nsq.InitNsq()
	nsq.InitConsumer()
	fmt.Println("RUNNING")
	handlerhttp.HandleRequests()
}
