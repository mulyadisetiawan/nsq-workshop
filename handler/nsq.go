package handler

import (
	"strconv"

	"github.com/mulyadisetiawan/nsq-workshop/module"
	"github.com/nsqio/go-nsq"
)

func GiveBenefitHandler(message *nsq.Message) error {
	userID, err := strconv.Atoi(string(message.Body))
	if err != nil {
		return err
	}

	err = module.GiveOVO(userID)
	if err != nil {
		return err
	}

	message.Finish()
	return nil
}
