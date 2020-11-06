package nsq

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/nsqio/go-nsq"
	"github.com/sharring_session/nsq-workshop/api/ovo"
)

func GiveBenefitNSQ(message *nsq.Message) error {
	var req ovo.GiveBenefitNSQRequest
	err := jsoniter.Unmarshal(message.Body, &req)
	if err != nil {
		fmt.Println("error unmarshal")
		nsq.Finish(message.ID)

		return err
	}

	err = ovo.GiveBenefit(req.UserID)
	if err != nil {
		return err
	}

	return nil
}
