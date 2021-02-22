package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	// "fmt"

	"github.com/faizinkholiq/todolist/helper"
	"github.com/faizinkholiq/todolist/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
) 

var collection = helper.Connection()

func main(){
	r := mux.NewRouter()

	r.HandleFunc("/todolist", getList).Methods("GET")

	log.Println("Running in localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getList(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var list []model.List

	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err,w)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var item model.List

		err := cur.Decode(&item)
		if err != nil{
			log.Fatal(err)
		}

		list = append(list, item)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(list)

}