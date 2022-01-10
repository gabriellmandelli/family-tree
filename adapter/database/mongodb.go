package database

import (
	"context"
	"fmt"
	"log"

	"github.com/gabriellmandelli/family-tree/adapter/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dataBaseName = "admin"
)

type MongoDB struct {
	client *mongo.Database
}

var mongoDb *mongo.Database

func NewMongoDbClient(ctx context.Context, config config.Configuration) (*mongo.Database, error) {
	if mongoDb == nil {
		clientOptions := options.Client().ApplyURI(config.MongoDB.URI)
		clientOptions.Auth = &options.Credential{
			Username: config.MongoDB.UserName,
			Password: config.MongoDB.Password,
		}

		// Connect to MongoDB
		client, err := mongo.Connect(ctx, clientOptions)

		if err != nil {
			log.Fatal(err)
		}

		// Check the connection
		err = client.Ping(ctx, nil)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Connected to MongoDB!")

		mongoDb = client.Database(dataBaseName)
	}

	return mongoDb, nil
}
