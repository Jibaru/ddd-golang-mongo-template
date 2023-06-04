package repositories

import (
	"context"
	"errors"
	"gitub.com/jibaru/ddd-golang-mongodb-template/internal/core/shared/connection"
	"gitub.com/jibaru/ddd-golang-mongodb-template/internal/core/version/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var (
	ErrConnectionFailed = errors.New("connection failed")
)

type mongoVersionRepository struct{}

type versionDocument struct {
	ID        string    `bson:"_id"`
	Value     string    `bson:"Value"`
	CreatedAt time.Time `bson:"CreatedAt"`
}

func (r *mongoVersionRepository) FindLatest() (*entities.Version, error) {
	client, err := connection.NewClient()
	if err != nil {
		log.Println(err)
		return nil, ErrConnectionFailed
	}

	ctx, err := connection.Connect(client)
	if err != nil {
		log.Print(err)
		return nil, ErrConnectionFailed
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err = client.Disconnect(ctx)
		if err != nil {
			log.Print(err)
		}
	}(client, ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Print(err)
		return nil, ErrConnectionFailed
	}

	db := client.Database("test")
	collection := db.Collection("versions")

	findOptions := options.Find()
	findOptions.SetSort(map[string]int{"CreatedAt": -1})
	findOptions.SetLimit(1)

	cursor, err := collection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		log.Print(err)
		return nil, ErrConnectionFailed
	}

	var results []versionDocument
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, ErrConnectionFailed
	}

	if len(results) == 1 {
		return entities.NewVersion(results[0].Value), nil
	}

	return entities.NewVersion("0.0.0"), nil
}
