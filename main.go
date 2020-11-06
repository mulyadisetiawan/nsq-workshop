package main

import (
	"fmt"

	"github.com/sharring_session/nsq/consumer"
	handlerhttp "github.com/sharring_session/nsq/http"
)

func main() {
	fmt.Println("RUNNING")
	consumer.Consume()
	handlerhttp.HandleRequests()
}
