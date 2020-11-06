package benefit

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/sharring_session/nsq/config"
)

type publisher interface {
	Publish(topic string, msg []byte) error
}

type Usecase struct {
	publisher publisher
}

func NewUsecase(publisher publisher) *Usecase {
	return &Usecase{
		publisher,
	}
}

func (uc *Usecase) PublishGiveBenefit(userIDStr string) error {
	paramBytes, err := jsoniter.Marshal(parsePublishBenefitParam(userIDStr))
	if err != nil {
		return err
	}

	return uc.publisher.Publish(config.NSQTopic, paramBytes)
}
