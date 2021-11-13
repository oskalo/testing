package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func AddComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	comment := new(Comment)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	comment.ID = 4
	comment.ProductID = mux.Vars(r)["productID"]

	repo := commentsRepository{}
	err = repo.SaveComment(*comment)

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	repo := commentsRepository{}
	comments, err := repo.GetComments(mux.Vars(r)["productID"])
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
