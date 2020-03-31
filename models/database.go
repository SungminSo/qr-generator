package models

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var QRCodes 	*mongo.Collection

func InitDB(db_host string, db_port string) {
	mongo_db_host := db_host
	if mongo_db_host == "" {
		mongo_db_host = "localhost"
	}

	mongo_db_port := db_port
	if mongo_db_port == "" {
		mongo_db_port = ":27017"
	}

	// Set mongoDB client options
	clientOptions := options.Client().ApplyURI("mongodb://" + mongo_db_host + mongo_db_port)

	mongoClient, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
	}

	QRCodes = mongoClient.Database("qr").Collection("qrcodes")
}