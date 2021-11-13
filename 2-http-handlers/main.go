package handlers

import (
	"github.com/gorilla/mux"

	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("comments/{productID}", AddComment).Methods(http.MethodPost)
	r.HandleFunc("comments/{productID}", GetComments).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}
