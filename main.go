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
	"go.mongodb.org/mongo-driver/bson/primitive"
) 

var collection = helper.Connection()

func main(){
	r := mux.NewRouter()

	r.HandleFunc("/todolist", getLists).Methods("GET")
	r.HandleFunc("/todolist/{id}", getList).Methods("GET")
	r.HandleFunc("/todolist", createList).Methods("POST")
	r.HandleFunc("/todolist/{id}", updateList).Methods("PUT")
	r.HandleFunc("/todolist/{id}", deleteList).Methods("DELETE")

	log.Println("Running in localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getLists(w http.ResponseWriter, r *http.Request){
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

func getList(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var list model.List
	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&list)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(list)
}

func createList(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var list model.List

	_ = json.NewDecoder(r.Body).Decode(&list)

	result, err := collection.InsertOne(context.TODO(), list)

	if err != nil{
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func updateList(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Conten-Type", "application/json")

	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	var list model.List

	filter := bson.M{"_id": id}

	_ = json.NewDecoder(r.Body).Decode(&list)

	update := bson.D{
		{"$set", bson.D{
			{"name", list.Name},
			{"status", list.Status},
		}},
	}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&list)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	list.ID = id

	json.NewEncoder(w).Encode(list)
}

func deleteList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
}