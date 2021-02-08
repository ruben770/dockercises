package main

import (
	"context"
	"log"
	"net/http"

	"github.com/ruben770/dockercises/docker-compose/ejercicio-1/solucion/pkg/api"
	"github.com/ruben770/dockercises/docker-compose/ejercicio-1/solucion/pkg/mongo"
)

func main() {
	err := mongo.NewConnection()
	if err != nil {
		panic(err)
	}
	defer mongo.Conn.GetClient().Disconnect(context.TODO())

	r := api.SetRoutes()

	log.Fatal(http.ListenAndServe(":7777", r))
}
