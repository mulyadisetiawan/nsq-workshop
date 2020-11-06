package main

import (
	"fmt"

	server "github.com/sharring_session/nsq/server"
)

func main() {
	fmt.Println("RUNNING")

	server.InitConsumer()
	server.HandleRequests()
}
