package controller

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var dbName string = "Go-Blog"
var collectionName string = "blogs"

var collection *mongo.Collection

func getConnectionUrl() string{
    err := godotenv.Load()
    if err != nil{
        log.Fatal("unable to access env file")
    }

    connectionUrl := os.Getenv("MONGO_URL")
    return connectionUrl
}

func ConnectDB() (*mongo.Client, error){
    dbUrl := getConnectionUrl()

    clientOption := options.Client().ApplyURI(dbUrl)
    client, err := mongo.Connect(clientOption)
    checkNilErr(err, "unable to connect to db")

    collection = client.Database(dbName).Collection(collectionName)
    fmt.Println("connection established")

    return client, nil
}

