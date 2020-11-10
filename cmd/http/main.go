package main

import (
	"fmt"
	mnsq "github.com/sharring_session/nsq/nsq"

	handlerhttp "github.com/sharring_session/nsq/http"
)

func main() {
	fmt.Println("RUNNING")
	mnsq.InitProducer()
	handlerhttp.HandleRequests()
}
