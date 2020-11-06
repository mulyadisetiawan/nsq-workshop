package handlerhttp

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/sharring_session/nsq/nsq-workshop/api/ovo"
	"github.com/sharring_session/nsq/nsq-workshop/consumer"
)

func giveBenefit(w http.ResponseWriter, r *http.Request) {

	userIDStr := r.URL.Query().Get("user_id")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = ovo.GiveBenefit(userID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(string("success"))
	return
}

func giveOVO(w http.ResponseWriter, r *http.Request) {

	userIDStr := r.URL.Query().Get("user_id")

	var response ovo.Response

	defer func() {
		resp, _ := json.Marshal(response)
		w.Write(resp)
	}()

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		response.Code = "300"
		response.Error = err.Error()
		return
	}

	if userID == 0 {
		response.Code = "300"
		response.Error = "user id is empty"
		return
	}

	response.Code = "200"
	return
}

func HandleRequests() {
	http.HandleFunc("/givebenefit", giveBenefit)

	http.HandleFunc("/giveovo", giveOVO)

	consumer.ConsumeGiveBenefit()

	log.Fatal(http.ListenAndServe(":10000", nil))
}
