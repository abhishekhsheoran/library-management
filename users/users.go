package users

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/abhishekhsheoran/library-management/models"
	"github.com/abhishekhsheoran/library-management/utils"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func verifyUserAttributes(user models.User) string {
	var error string
	if user.Name == "" {
		error = "user name cannot be empty"
		return error
	}
	if user.Contact == 0 {
		error = "user contact cannot be empty"
		return error
	}
	if user.ID == "" {
		error = "user iD cannot be empty"
		return error
	}
	error = ""
	return error
}

func CreateUsers(w http.ResponseWriter, r *http.Request) {
	var user models.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		log.Println("error occured during decoding the input", err)
		http.Error(w, "error occured during decoding the input", http.StatusBadRequest)
		return
	}

	verifyResult := verifyUserAttributes(user)
	if verifyResult != "" {
		http.Error(w, verifyResult, http.StatusBadRequest)
		return
	}
	collection := utils.GetCollection()
	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		http.Error(w, "error occured during insertion of data", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func SignInUser(w http.ResponseWriter, r *http.Request) {
	input := mux.Vars(r)
	email := input["email"]
	password := input["password"]
	filter := bson.M{email: "email", password: "Password"}
	collection := utils.GetCollection()
	findEmail := collection.FindOne(context.TODO(), email)
	if findEmail.Err() != nil {
		http.Error(w, "Invalid e-mail address", http.StatusBadRequest)
		return
	}
	findPassword := collection.FindOne(context.TODO(), password)
	if findPassword.Err() != nil {
		http.Error(w, "Invalid password", http.StatusBadRequest)
		return
	}
	findResult := collection.FindOne(context.TODO(), filter)
	if findResult.Err() != nil {
		http.Error(w, "email and password did not match ", http.StatusBadRequest)
		return
	}
	decodedResult, err := findResult.DecodeBytes()
	if err != nil {
		http.Error(w, "error occured into internal process", http.StatusInternalServerError)
		return
	}
	w.Write(decodedResult)
}
