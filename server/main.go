package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Concerned-Doggo/Go-React-Blog/controller"
	"github.com/Concerned-Doggo/Go-React-Blog/router"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var client *mongo.Client

func init(){
    client, _ = controller.ConnectDB()
}

func main(){
    fmt.Println("ready to go")
    router := router.Router()

    defer func() {
        fmt.Println("disconnected from db")
        if err := client.Disconnect(context.Background()); err != nil {
            panic(err)
        }
    }()

    fmt.Println("server is listening on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
