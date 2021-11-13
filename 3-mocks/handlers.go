package handlers

import (
	"encoding/json"
	"net/http"
	"testing/3-mocks/repository"

	"github.com/gorilla/mux"
)

type headerComment struct {
	repository repository.CommentsRepository
}

func (hc headerComment) AddComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	comment := new(repository.Comment)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	comment.ID = 4
	comment.ProductID = mux.Vars(r)["productID"]

	err = hc.repository.SaveComment(*comment)

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (hc headerComment) GetComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	comments, err := hc.repository.GetComments(mux.Vars(r)["productID"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(comments)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
