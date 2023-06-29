package main

import (
	"log"
	"net/http"

	"github.com/abhishekhsheoran/library-management/users"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/kuch/bhi/",users.CreateUsers).Methods(http.MethodPost)
	log.Println(http.ListenAndServe(":8080", router))
}

func defaultHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello from library-management"))
}
