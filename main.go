package main

import (
	"fmt"

	handlerhttp "github.com/sharring_session/nsq/http"
)

func main() {
	fmt.Println("RUNNING")
	handlerhttp.HandleRequests()
	giveOVO()
}
