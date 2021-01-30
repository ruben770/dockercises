package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type People struct {
	XMLName xml.Name `xml:"people"`
	Person  []Person `xml:"person"`
}

type Person struct {
	ID          int    `bson:"_id,omitempty" xml:"id"`
	FirstName   string `bson:"first_name,omitempty" xml:"first_name"`
	LastName    string `bson:"last_name,omitempty" xml:"last_name"`
	Company     string `bson:"company,omitempty" xml:"company"`
	Email       string `bson:"email,omitempty" xml:"email"`
	IPAddress   string `bson:"ip_address,omitempty" xml:"ip_address"`
	PhoneNumber string `bson:"phone,omitempty" xml:"phone_number"`
}

func main() {
	client, err := GetMongoClient()
	if err != nil {
		panic(err)
	}

	defer client.Disconnect(context.TODO())

	people := client.Database("compose").Collection("people")

	var p People
	xmlToPeople("../data/people.xml", &p)

	// Made a slice of type interface{} bc bellow insertMany() won't let me use a slice of a custom struct Person
	docs := make([]interface{}, len(p.Person))
	for i, v := range p.Person {
		docs[i] = v
	}

	insertManyResult, err := people.InsertMany(context.TODO(), docs)
	if err != nil {
		panic(err)
	}

	fmt.Println("Inserted", len(insertManyResult.InsertedIDs), "documents!")
}

func xmlToPeople(xmlRoute string, people *People) {
	file, err := ioutil.ReadFile(xmlRoute)
	if err != nil {
		panic(err)
	}

	xml.Unmarshal(file, people)
}
