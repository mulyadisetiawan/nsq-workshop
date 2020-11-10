package main

import (
	"fmt"
	mnsq "github.com/sharring_session/nsq/nsq"
)

func main() {
	fmt.Println("RUNNING")
	mnsq.InitConsumer()
}
