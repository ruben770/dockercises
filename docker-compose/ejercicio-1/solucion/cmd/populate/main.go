package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"

	"github.com/ruben770/dockercises/docker-compose/ejercicio-1/solucion/pkg/model"
	"github.com/ruben770/dockercises/docker-compose/ejercicio-1/solucion/pkg/mongo"
)

func main() {
	err := mongo.NewConnection()
	if err != nil {
		panic(err)
	}

	defer mongo.Conn.GetClient().Disconnect(context.TODO())

	var p model.People
	xmlToPeople("../people.xml", &p)

	// Made a slice of type interface{} bc bellow insertMany() won't let me use a slice of a custom struct Person
	docs := make([]interface{}, len(p.Person))
	for i, v := range p.Person {
		docs[i] = v
	}

	insertManyResult, err := mongo.Conn.InsertMany(docs)
	if err != nil {
		panic(err)
	}

	fmt.Println("Inserted", len(insertManyResult.InsertedIDs), "documents!")
}

func xmlToPeople(xmlRoute string, people *model.People) {
	file, err := ioutil.ReadFile(xmlRoute)
	if err != nil {
		panic(err)
	}

	xml.Unmarshal(file, people)
}
