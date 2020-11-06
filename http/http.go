package handlerhttp

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/sharring_session/nsq/api/ovo"
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

func HandleRequests() {
	http.HandleFunc("/givebenefit", giveBenefit)

	log.Fatal(http.ListenAndServe(":10000", nil))
}
