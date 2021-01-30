package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoClient() (client *mongo.Client, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)

	defer cancel()

	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo_compose:27017"))
	return
}
