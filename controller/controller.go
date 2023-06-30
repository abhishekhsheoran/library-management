package controller

import (
	"github.com/abhishekhsheoran/library-management/users"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func StartServer() {
	router := mux.NewRouter()
	router.HandleFunc("api/v1/user/", users.CreateUsers).Methods(http.MethodPost)
	log.Println(http.ListenAndServe(":8080", router))
}
