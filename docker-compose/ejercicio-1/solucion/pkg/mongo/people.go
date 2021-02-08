package mongo

import (
	"context"
	"time"

	"github.com/ruben770/dockercises/docker-compose/ejercicio-1/solucion/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (c *connection) FindPerson(id int, person *model.Person) error {
	ctx, _ := context.WithTimeout(context.TODO(), time.Second*5)
	people := c.Client.Database("compose").Collection("people")
	err := people.FindOne(ctx, bson.M{"_id": id}).Decode(person)
	if err != nil {
		return err
	}
	return nil
}

func (c *connection) FindAll() (cur *mongo.Cursor, err error) {
	ctx, _ := context.WithTimeout(context.TODO(), time.Second*5)
	people := c.Client.Database("compose").Collection("people")
	cur, err = people.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	return cur, nil
}

func (c *connection) InsertMany(docs []interface{}) (r *mongo.InsertManyResult, err error) {
	ctx, _ := context.WithTimeout(context.TODO(), time.Second*5)
	people := c.Client.Database("compose").Collection("people")
	r, err = people.InsertMany(ctx, docs)
	if err != nil {
		return nil, err
	}
	return r, nil
}
