package main

import (
	"fmt"

	handlerhttp "github.com/sharring_session/nsq/http"
	nsq_publisher "github.com/sharring_session/nsq/nsq"
)

func main() {
	nsq_publisher.InitNSQ()
	fmt.Println("RUNNING")
	handlerhttp.HandleRequests()
}
