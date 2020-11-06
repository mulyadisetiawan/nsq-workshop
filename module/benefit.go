package module

import (
	"github.com/mulyadisetiawan/nsq-workshop/api/ovo"
	"github.com/mulyadisetiawan/nsq-workshop/server"
)

func GiveOVO(userID int) error {
	err := ovo.GiveBenefit(userID)
	if err != nil {
		return err
	}

	return nil
}

func PublishGiveBenefit(userID int) error {
	err := server.ProducerClient.Publish(server.NSQPrefix+server.NSQTopicGiveBenefit, userID)
	if err != nil {
		return err
	}

	return nil
}
