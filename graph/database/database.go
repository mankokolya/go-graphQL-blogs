package database

import (
	"context"
	"go-graphql-blog/graph/utils"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client

	Database *mongo.Database
}

var DB MongoInstance

func Connect(dbName string) error {
	client, _ := mongo.NewClient(options.Client().ApplyURI(utils.GetValue("MONGO_URI")))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := client.Connect(ctx)

	if err != nil {
		return err
	}

	var db *mongo.Database = client.Database(dbName)

	DB = MongoInstance{
		Client:   client,
		Database: db,
	}

	return nil
}

func GetCollection(name string) *mongo.Collection {
	return DB.Database.Collection(name)
}
