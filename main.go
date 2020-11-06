package main

import (
	"fmt"

	"github.com/sharring_session/nsq-workshop/consumer"
	handlerhttp "github.com/sharring_session/nsq-workshop/http"
)

func main() {
	fmt.Println("RUNNING")
	consumer.Consumer()
	handlerhttp.HandleRequests()
}
