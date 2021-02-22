package helper

import (
	"context"
	"encoding/json"
	"net/http"
	"log"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
) 

func Connection() *mongo.Collection {

	clientOptions := options.Client().ApplyURI("mongodb+srv://root:root@maincluster.0j63d.mongodb.net/todolist_go?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	collection := client.Database("todolist_go").Collection("list")

	return collection
}

type ErrorResponse struct {
	StatusCode int `json:"status"`
	ErrorMessage string  `json:"message"`
}

func GetError(err error, w http.ResponseWriter){
	log.Fatal(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode: http.StatusInternalServerError,
	}

	message, _:= json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}