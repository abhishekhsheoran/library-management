package books

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_createBook(t *testing.T) {
	body :=
		`
	{
	"Name" : "Two States",
	"Auther" : "Chetan Bhagat",
	"IssuedBy" :["rughbir"],
	"NumOfCopy" : 1
	}
	`
	req, err := http.NewRequest(http.MethodPost, "http.localhost:8080/create/book/books", bytes.NewBuffer([]byte(body)))
	if err != nil {
		log.Println("error occured during making request")
	t.Fail()
	}
	response := httptest.NewRecorder()
	CreateBook(response,req)
	if response.Result().StatusCode!= http.StatusCreated{
		log.Printf("got different status code, which is %v", response.Result().StatusCode)
		t.Fail()
	}
}


func Test_deleteBook(t *testing.T){
	
}