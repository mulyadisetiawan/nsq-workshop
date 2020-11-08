package nsq

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	gonsq "github.com/nsqio/go-nsq"
	"github.com/sharring_session/nsq/api/ovo"
)

func GiveOvo(message *gonsq.Message) error {
	payload := new(ovo.Request)
	err := json.Unmarshal(message.Body, &payload)
	if err != nil {
		return err
	}

	resp, err := http.Get(fmt.Sprintf("http://localhost:10000/giveovo?user_id=%d", payload.UserID))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var response ovo.Response
	err = json.Unmarshal(data, &response)
	if err != nil {
		return err
	}

	if response.Code != "200" {
		return fmt.Errorf("Error: " + response.Error)
	}

	message.Finish()
	return nil
}
