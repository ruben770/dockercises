package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/ruben770/dockercises/docker-compose/ejercicio-1/solucion/pkg/model"
	"github.com/ruben770/dockercises/docker-compose/ejercicio-1/solucion/pkg/mongo"
)

func getPerson(w http.ResponseWriter, r *http.Request) {
	personId := chi.URLParam(r, "personId")
	if personId == "" {
		http.Error(w, "Resource not found", 404)
		return
	}
	person := &model.Person{}
	id, err := strconv.Atoi(personId)
	err = mongo.Conn.FindPerson(id, person)
	if err != nil {
		http.Error(w, fmt.Sprintf("Person with the id %v was not found.", personId), 404)
		return
	}

	json.NewEncoder(w).Encode(person)
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	cur, err := mongo.Conn.FindAll()
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	defer cur.Close(ctx)

	var people []*model.Person

	for cur.Next(ctx) {
		person := &model.Person{}
		err := cur.Decode(person)
		if err != nil {
			log.Fatal(err)
		}
		people = append(people, person)
	}

	json.NewEncoder(w).Encode(people)
}
