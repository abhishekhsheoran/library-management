package main

import (
	"net/http"

	"github.com/abhishekhsheoran/library-management/controller"
)

func main() {
	controller.Path()

}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from library-management"))
}
