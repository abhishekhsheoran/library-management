package users

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUsers(t *testing.T) {
	body :=
		`
	{
		"NAME" :  "abhishek",
		"ID" : "Sheoran",
		"contct": "0000"
	
	}
	`

	request, err := http.NewRequest(http.MethodPost, "http.localhost:8080/kuch/bhi/", bytes.NewBuffer([]byte(body)))
	if err != nil {
		log.Println("error occured in making request")
	}
	response := httptest.NewRecorder()
	CreateUsers(response, request)
	if response.Result().StatusCode != http.StatusOK {
		log.Println("response does not match", response.Result().StatusCode)
		t.Fail()
	}
}
