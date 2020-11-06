package main

import (
	"fmt"

	handlerhttp "github.com/sharring_session/nsq/nsq-workshop/http"
)

func main() {
	fmt.Println("RUNNING")
	handlerhttp.HandleRequests()
}
