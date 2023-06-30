package main

import (
	"log"
	"net/http"

	"github.com/abhishekhsheoran/library-management/users"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/users/", users.CreateUsers).Methods(http.MethodPost)
	router.HandleFunc("/", defaultHandler).Methods(http.MethodPost)

	log.Println(http.ListenAndServe(":8080", router))
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from library-management"))
}
