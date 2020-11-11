package ovo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Code  string `json:"code"`
	Error string `json:"error"`
}

// deprecated
func GiveBenefit(userID int) error {

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
