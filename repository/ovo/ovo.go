package ovo

import (
	"errors"
	"strconv"
)

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Give(userIDStr string) (string, error) {
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return "300", err
	}

	if userID == 0 {
		return "300", errors.New("user id is empty")
	}

	return "200", nil
}
