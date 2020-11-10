package mnsq

import (
	jsoniter "github.com/json-iterator/go"
	nsq "github.com/nsqio/go-nsq"
	"github.com/sharring_session/nsq/api/ovo"
)

const (
	TopicGiveOvoBenefit = "give_ovo_benefit"
)

type OvoPayload struct {
	UserID int `json:"user_id"`
}

func GiveBenefitOvo(msg *nsq.Message) (err error) {
	payload := OvoPayload{}
	err = jsoniter.Unmarshal(msg.Body, &payload)
	if err != nil {
		msg.Finish()
		return err
	}

	err = ovo.GiveBenefit(payload.UserID)
	if err != nil {
		return err
	}

	msg.Finish()

	return nil
}
