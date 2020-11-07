package benefit

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	publishBenefit publishBenefit
}

type publishBenefit interface {
	PublishGiveBenefit(userIDStr string) error
}

func NewHandler(publishBenefit publishBenefit) {
	handler := &Handler{
		publishBenefit,
	}

	http.HandleFunc("/givebenefit", handler.giveBenefit)
}

func (h *Handler) giveBenefit(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")

	err := h.publishBenefit.PublishGiveBenefit(userIDStr)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(string("success"))
	return
}
