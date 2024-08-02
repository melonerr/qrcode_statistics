package repositories

import (
	"context"
	"log"
	"qrcode_statistics/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectDB() {
	config := config.Load()

	client, err := mongo.NewClient(options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	Client = client
	log.Println("MongoDB connected successfully")
}

func GetCollection(collectionName string) *mongo.Collection {
	config := config.Load()
	database := Client.Database(config.MongoURI)
	return database.Collection(collectionName)
}
