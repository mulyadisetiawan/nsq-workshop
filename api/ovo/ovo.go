package ovo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sharring_session/nsq/nsq-workshop/publisher"
)

type Response struct {
	Code  string `json:"code"`
	Error string `json:"error"`
}

type RequestGiveBenefit struct {
	UserID int
}

func GiveBenefit(userID int) error {

	req := RequestGiveBenefit{
		UserID: userID,
	}

	payload, err := json.Marshal(req)
	if err != nil {
		return err
	}

	return publisher.PublishGiveBenefit(payload)

}

func GetBenefitFromOvoNSQ(userID int) error {
	fmt.Println("GetBenefitFromOvo....")
	fmt.Println("UserID: ", userID)

	resp, err := http.Get(fmt.Sprintf("http://localhost:10000/giveovo?user_id=%d", userID))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var response Response
	err = json.Unmarshal(data, &response)
	if err != nil {
		return err
	}

	if response.Code != "200" {
		return fmt.Errorf("Error give ovo: " + response.Error)
	}

	return nil
}
