package repositories

import (
	"context"
	"fmt"
	"log"
	"time"

	"qrcode_statistics/internal/config"
	"qrcode_statistics/internal/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetEventById(id string) (*models.Events, error) {
	collection := config.GetCollection("events")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid ID:", err)
		return nil, err
	}

	var event models.Events
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&event)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("No event found with the given ID")
			return nil, nil
		}
		log.Println("FindOne Error:", err)
		return nil, err
	}

	return &event, nil
}
func CreateEvent(event models.Events) (*mongo.InsertOneResult, error) {
	eventCollection := config.GetCollection("events")
	memberCollection := config.GetCollection("members")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ID, err := primitive.ObjectIDFromHex(event.U_id)
	if err != nil {
		log.Println("Invalid ID:", err)
		return nil, err
	}

	var User models.Members
	err = memberCollection.FindOne(ctx, bson.M{"_id": ID}).Decode(&User)
	if err != nil {
		return nil, fmt.Errorf("user with ID %v does not exist", event.U_id)
	}

	// User exists, create the event
	event.ID = primitive.NewObjectID()
	result, err := eventCollection.InsertOne(ctx, event)
	if err != nil {
		log.Println("InsertEvent Error:", err)
		return nil, err
	}
	return result, nil
}

func UpdateEvent(id string, event models.Events) (*mongo.UpdateResult, error) {
	collection := config.GetCollection("events")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid ID:", err)
		return nil, err
	}

	update := bson.M{
		"$set": event,
	}

	result, err := collection.UpdateOne(ctx, bson.M{"_id": objectID, "u_id": event.U_id}, update)
	if err != nil {
		log.Println("Update Error:", err)
		return nil, err
	}

	return result, nil
}
func DeleteEvent(id string) (*mongo.UpdateResult, error) {
	// soft delete
	collection := config.GetCollection("events")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid ID:", err)
		return nil, err
	}

	update := bson.M{
		"$set": bson.M{
			"status": false,
		},
	}

	result, err := collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		log.Println("Delete Error:", err)
		return nil, err
	}

	return result, nil
}
