package main

import (
	"fmt"

	handlerhttp "github.com/sharing_session/nsq-workshop/http"
	"github.com/sharing_session/nsq-workshop/nsq/consumer"
)

func main() {
	consumer.Consume()
	fmt.Println("RUNNING")
	handlerhttp.HandleRequests()

}
