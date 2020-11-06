package handlerhttp

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/sharring_session/nsq/api/ovo"
	nsq_publisher "github.com/sharring_session/nsq/nsq"
)

const (
	//TOPIC
	NSQ_TOPIC_PUBLISH_GIVE_BENEFIT = "workshop_nsq_publish_ovo_henrys"
)

func giveBenefitNSQ(w http.ResponseWriter, r *http.Request) {

	userIDStr := r.URL.Query().Get("user_id")

	nsq_publisher.Producer.Publish(NSQ_TOPIC_PUBLISH_GIVE_BENEFIT, []byte(userIDStr))

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
	http.HandleFunc("/givebenefit", giveBenefitNSQ)

	http.HandleFunc("/giveovo", giveOVO)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
