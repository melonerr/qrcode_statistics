package repositories

import (
	"context"
	"log"
	"time"

	"qrcode_statistics/internal/config"
	"qrcode_statistics/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateMember(user models.Members) (*mongo.InsertOneResult, error) {
	collection := config.GetCollection("members")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user.ID = primitive.NewObjectID()
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Println("InsertUser Error:", err)
		return nil, err
	}

	return result, nil
}
