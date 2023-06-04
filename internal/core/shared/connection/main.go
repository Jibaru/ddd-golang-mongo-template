package connection

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const uri string = "mongodb://golang:password@mongo/test?retryWrites=true&w=majority&authSource=admin"

func NewClient() (*mongo.Client, error) {
	return mongo.NewClient(options.Client().ApplyURI(uri))
}

func Connect(client *mongo.Client) (context.Context, error) {
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	return ctx, client.Connect(ctx)
}
