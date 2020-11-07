package ovo

import (
	"fmt"

	"github.com/bitly/go-nsq"
	jsoniter "github.com/json-iterator/go"
	"github.com/sharring_session/nsq/repository/benefit"
	"github.com/sharring_session/nsq/repository/ovo"
)

type giveOVOItf interface {
	GiveOVO(userIDStr string) (ovo.Response, error)
}

type Handler struct {
	giveOVO giveOVOItf
}

func NewHandler(giveOVO giveOVOItf) *Handler {
	return &Handler{
		giveOVO,
	}
}

func (h *Handler) EventGive(message *nsq.Message) error {
	var (
		result ovo.Response
		err    error
		param  benefit.PublishParam
	)

	err = jsoniter.Unmarshal(message.Body, &param)
	if err != nil {
		return err
	}

	defer func() {
		fmt.Println("response code: ", result.Code)
		fmt.Println("response error: ", result.Error)
		message.Finish()
	}()

	result, err = h.giveOVO.GiveOVO(param.UserIDStr)
	if err != nil {
		return err
	}

	return nil
}
