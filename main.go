package main

import (
	"flag"
	"fmt"

	handler "github.com/mulyadisetiawan/nsq-workshop/handler"
	"github.com/mulyadisetiawan/nsq-workshop/server"
)

func main() {
	appType := flag.String("type", "http", "App Type (http/nsq)")
	flag.Parse()

	switch *appType {
	case "http":
		fmt.Println("RUNNING HTTP")
		server.InitProducer()
		handler.HandleRequests()
	case "nsq":
		fmt.Println("RUNNING NSQ")
		server.InitConsumer(handler.GiveBenefitHandler)
	}

}
