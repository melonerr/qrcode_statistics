package repositories

import (
	"context"
	"log"
	"time"

	"qrcode_statistics/internal/config"
	"qrcode_statistics/internal/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetMemberById(id string) (*models.Members, error) {
	collection := config.GetCollection("members")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid ID:", err)
		return nil, err
	}

	var member models.Members
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&member)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("No member found with the given ID")
			return nil, nil
		}
		log.Println("FindOne Error:", err)
		return nil, err
	}

	return &member, nil
}
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
func UpdateMember(id string, user models.Members) (*mongo.UpdateResult, error) {
	collection := config.GetCollection("members")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid ID:", err)
		return nil, err
	}

	update := bson.M{
		"$set": user,
	}

	result, err := collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		log.Println("Update Error:", err)
		return nil, err
	}

	return result, nil
}
func DeleteMember(id string) (*mongo.UpdateResult, error) {
	// soft delete
	collection := config.GetCollection("members")

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
