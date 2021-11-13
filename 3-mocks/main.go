package handlers

import (
	"github.com/gorilla/mux"
	"testing/3-mocks/repository"

	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	hc := headerComment{
		repository: repository.NewCommentsRepository(),
	}

	r.HandleFunc("comments/{productID}", hc.AddComment).Methods(http.MethodPost)
	r.HandleFunc("comments/{productID}", hc.GetComments).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}
