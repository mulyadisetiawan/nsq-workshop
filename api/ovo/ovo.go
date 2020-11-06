package ovo

import (
	"encoding/json"

	"github.com/sharring_session/nsq/producer"
)

type Response struct {
	Code  string `json:"code"`
	Error string `json:"error"`
}

type RequestOVO struct {
	UserID int `json:"user_id"`
}

func GiveBenefit(userID int) error {
	req := RequestOVO{UserID: userID}
	bytes, err := json.Marshal(req)
	if err != nil {
		return err
	}
	return producer.Publish(bytes)
}
