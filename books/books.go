package books

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/abhishekhsheoran/library-management/models"
	"github.com/abhishekhsheoran/library-management/utils"
	"github.com/gorilla/mux"
)

func verifyBook(book models.Book) string {
	var returnStr = ""
	if book.Name == "" {
		returnStr = "Book Name can not be empty"
		return returnStr
	}
	if book.Auther == "" {
		returnStr = "Book Auther can not be empty"
		return returnStr
	}
	if book.IssuedBy == nil {
		returnStr = "Must have to tell who issued the book"
		return returnStr
	}
	if book.NumOfCopy == 0 {
		returnStr = "Number of copy must be mentioned"
		return returnStr
	}
	returnStr = ""
	return returnStr
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "error occured during getting the input", http.StatusBadRequest)
		return
	}
	verifyResult := verifyBook(book)
	if verifyResult != "" {
		http.Error(w, verifyResult, http.StatusBadRequest)
		return
	}
	collection := utils.GetCollection(utils.Books)
	_, err = collection.InsertOne(context.TODO(), book)
	if err != nil {
		http.Error(w, "error occured during insertion of data", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	input := mux.Vars(r)
	bookName := input["bookName"]
	collection := utils.GetCollection(utils.Books)
	findResult := collection.FindOne(context.TODO(), bookName)

	if findResult.Err() != nil {
		http.Error(w, "error occured during finding the document", http.StatusInternalServerError)
		return
	}
	_,err := collection.DeleteOne(context.TODO(), bookName)
	if err != nil {
		http.Error(w, "error occured durin deletion of record", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
