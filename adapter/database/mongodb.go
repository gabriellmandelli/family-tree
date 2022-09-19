package database

import (
	"context"
	"fmt"

	"github.com/gabriellmandelli/family-tree/adapter/config"
	"github.com/joomcode/errorx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dataBaseName = "admin"
)

type MongoDB struct {
	client *mongo.Database
}

func NewMongoDbClient(ctx context.Context, config config.Configuration) (*mongo.Database, *errorx.Error) {
	var errx *errorx.Error

	clientOptions := options.Client().ApplyURI(config.MongoDB.URI)
	clientOptions.Auth = &options.Credential{
		Username: config.MongoDB.UserName,
		Password: config.MongoDB.Password,
	}

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		errx = errorx.Decorate(err, "Erro to connect MongoDb")
	}

	// Check the connection
	err = client.Ping(ctx, nil)

	if err != nil {
		errx = errorx.Decorate(err, "Erro to check conection MongoDb")
	}

	fmt.Println("Connected to MongoDB")
	mongoDb := client.Database(dataBaseName)

	return mongoDb, errx
}
