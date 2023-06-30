package utils

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Connection *mongo.Client
)

func InitialiseDatabase() {
	options := options.Client().ApplyURI("mongodb://localhost:27017/")
	var err error
	Connection, err = mongo.Connect(context.TODO(), options)
	if err != nil {
		log.Println("error occured in mongo connection", err)
		return
	}
}
