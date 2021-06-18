package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConnection holds the client iformation to connect to DB
var MongoConnection = ConnectDataBase()

var connectionString string = "mongodb+srv://rodrijara:cllamaL.1@cluster0.lmdlc.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
var clientOptions = options.Client().ApplyURI(connectionString)

// ConnectDataBase sets the connection to the mongo DB,
// returns a mongo.Client object
func ConnectDataBase() *mongo.Client {

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Connection SUCCESSFUL")
	return client

}

// CheckConnection evaluates a successful connection to DB
func CheckConnection() int {
	err := MongoConnection.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1

}