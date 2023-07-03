package users

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os/user"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestUsers(t *testing.T) {
	body :=
		`
	{
		"NAME" :  "abhishek",
		"ID" : "Sheoran",
		"contct": 0000
	
	}
	`

	request, err := http.NewRequest(http.MethodPost, "http.localhost:8080/api/signUp/users/", bytes.NewBuffer([]byte(body)))
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

func Test_signIn(t *testing.T) {
	req, err := http.NewRequest(http.MethodConnect, "http.localhost:8080/login?email&password", bytes.NewBufferString(""))
	if err != nil {
		log.Println("error occured during makin request body")
	}
	response := httptest.NewRecorder()
	SignInUser(response, req)
	var user user.User
	bson.UnmarshalExtJSON([]byte(user), user)
	if response.Header() != user {

	}
}

func Test_updateUser(t *testing.T) {
	body :=
		`
		{
	"password" : "kuchbhi",
	"contact": 0000,
	"email": "abhishek@sheoran.com"
		}
	`
	req, err := http.NewRequest(http.MethodPatch, "http.localhost:8080/update/user/users", bytes.NewBuffer([]byte(body)))
	if err != nil {
		log.Println("error occured during making req body")
	}
	response := httptest.NewRecorder()
	UpdateUser(response, req)
	if response.Result().StatusCode != http.StatusOK {
		log.Println("expecting status code statusOK but got %v", response.Result().Status)
		t.Fail()
	}
}
