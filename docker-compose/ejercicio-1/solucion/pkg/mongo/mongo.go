package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Conn *connection

type connection struct {
	Client *mongo.Client
}

func NewConnection() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo_compose:27017"))
	if err != nil {
		return err
	}

	Conn = &connection{Client: client}

	return nil
}

func (c *connection) GetClient() (client *mongo.Client) {
	return c.Client
}
