package controller

import (
	"log"
	"net/http"

	"github.com/abhishekhsheoran/library-management/books"
	"github.com/abhishekhsheoran/library-management/users"
	"github.com/gorilla/mux"
)

func Path() {
	router := mux.NewRouter()
	router.HandleFunc("/",main.defaultHandler).Methods(http.MethodPost)
	router.HandleFunc("/api/signUp/users/", users.CreateUsers).Methods(http.MethodPost)
	router.HandleFunc("/login", users.SignInUser).Methods(http.MethodConnect)
	router.HandleFunc("/update/user/users", users.UpdateUser).Methods(http.MethodPatch)
	router.HandleFunc("/create/book/books",books.CreateBook).Methods(http.MethodPost)
	router.HandleFunc("/delete/book/books",books.DeleteBook).Methods(http.MethodDelete)
	log.Println(http.ListenAndServe(":8080", router))

}
