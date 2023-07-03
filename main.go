package main

import (
	"log"
	"net/http"

	"github.com/abhishekhsheoran/library-management/users"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", defaultHandler).Methods(http.MethodPost)
	router.HandleFunc("/api/signUp/users/", users.CreateUsers).Methods(http.MethodPost)
	router.HandleFunc("/login", users.SignInUser).Methods(http.MethodConnect)
	router.HandleFunc("/update/user/users", users.UpdateUser).Methods(http.MethodPatch)
	log.Println(http.ListenAndServe(":8080", router))
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from library-management"))
}
