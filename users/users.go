package users

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/abhishekhsheoran/library-management/models"
	"github.com/abhishekhsheoran/library-management/utils"
)

func CreateUsers(w http.ResponseWriter, r *http.Request) {
	var user models.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		log.Println("error occured during decoding the input", err)
		http.Error(w, "error occured during decoding the input", http.StatusBadRequest)
		return
	}
	if user.ID == "" {
		http.Error(w, "user id cannot be empty", http.StatusBadRequest)
		return
	}
	if user.Name == "" {
		http.Error(w, "user name cannot be empty", http.StatusBadRequest)
		return
	}
	if user.Contact == 0 {
		http.Error(w, "user contact cannot be empty", http.StatusBadRequest)
		return
	}

	utils.InitialiseDatabase()
	collection := utils.Connection.Database(utils.LMDB).Collection(utils.Users)
	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Println("error occured during insertion of data", err)
		http.Error(w, "error occured during insertion of data", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
