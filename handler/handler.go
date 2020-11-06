package handler

import (
	"encoding/json"
	"fmt"

	"github.com/nsqio/go-nsq"
	"github.com/sharring_session/nsq/nsq-workshop/api/ovo"
)

func HandlerGiveBenefitOVO(message *nsq.Message) error {
	fmt.Println("[Incoming Message HandlerGiveBenefitOVO]", string(message.Body))

	var req ovo.RequestGiveBenefit
	err := json.Unmarshal(message.Body, &req)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = ovo.GetBenefitFromOvoNSQ(req.UserID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
