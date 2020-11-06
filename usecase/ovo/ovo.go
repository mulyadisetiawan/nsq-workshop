package ovo

import (
	"github.com/sharring_session/nsq/repository/ovo"
)

type giveItf interface {
	Give(userIDStr string) (string, error)
}

type Usecase struct {
	give giveItf
}

func NewUsecase(give giveItf) *Usecase {
	return &Usecase{
		give,
	}
}

func (uc *Usecase) GiveOVO(userIDStr string) (ovo.Response, error) {
	code, err := uc.give.Give(userIDStr)
	return parseOVOResponse(code, err), err
}
