package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var people *mongo.Collection

func main() {
	client, err := GetMongoClient()
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.TODO())

	people = client.Database("compose").Collection("people")

	r := setRoutes()

	http.ListenAndServe(":7777", r)
}

func setRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/people", func(r chi.Router) {
		r.Get("/", getPeople)
		r.Get("/{personId}", getPerson)
	})
	return r
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	personId := chi.URLParam(r, "personId")
	if personId == "" {
		http.Error(w, "Resource not found", 404)
		return
	}
	person := &Person{}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	id, err := strconv.Atoi(personId)
	err = people.FindOne(ctx, bson.M{"_id": id}).Decode(person)
	if err != nil {
		http.Error(w, fmt.Sprintf("Person with the id %v was not found.", personId), 404)
		return
	}

	json.NewEncoder(w).Encode(person)
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	cur, err := people.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	defer cur.Close(ctx)

	var people []*Person

	for cur.Next(ctx) {
		person := &Person{}
		err := cur.Decode(person)
		if err != nil {
			log.Fatal(err)
		}
		people = append(people, person)
	}

	json.NewEncoder(w).Encode(people)
}
