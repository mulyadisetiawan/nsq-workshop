package main

import (
	"fmt"

	handlerhttp "github.com/natasharamdani/nsq-workshop/http"
	nsqconsumer "github.com/natasharamdani/nsq-workshop/nsq/consumer"
	nsqproducer "github.com/natasharamdani/nsq-workshop/nsq/producer"
)

func main() {
	fmt.Println("RUNNING")
	handlerhttp.HandleRequests()
	nsqproducer.Init()
	nsqconsumer.Init()
}
