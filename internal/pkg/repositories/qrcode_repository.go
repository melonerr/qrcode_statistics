package repositories

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"qrcode_statistics/internal/config"
	"qrcode_statistics/internal/pkg/models"
	"qrcode_statistics/internal/pkg/service"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetQrcodeById(id string) (*models.QrcodeRes, error) {
	collection := config.GetCollection("qrcode")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid ID:", err)
		return nil, err
	}

	var qrcode models.Qrcode
	err = collection.FindOne(ctx, bson.M{"_id": objectID, "status": true}).Decode(&qrcode)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("No event found with the given ID")
			return nil, nil
		}
		log.Println("FindOne Error:", err)
		return nil, err
	}

	BaseUrl, err := getEnv("APP_URI")
	if err != nil {
		return nil, err
	}

	ShortUrl := BaseUrl + "/" + qrcode.ShortUrl
	QrcodeGenerator, err := service.QecodeGenerator(ShortUrl)
	if err != nil {
		return nil, fmt.Errorf("error generating QR code: %v", err)
	}

	resData := &models.QrcodeRes{
		ID:       qrcode.ID,
		Title:    qrcode.Title,
		Target:   qrcode.Target,
		ShortUrl: ShortUrl,
		Qrcode:   QrcodeGenerator,
		Status:   qrcode.Status,
	}
	return resData, nil
}
func CreateQrcode(event models.Qrcode) (any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	eventCollection := config.GetCollection("events")
	memberCollection := config.GetCollection("members")
	qrcodeCollection := config.GetCollection("qrcode")

	// Validate and convert IDs
	U_ID, err := primitive.ObjectIDFromHex(event.U_id)
	if err != nil {
		return nil, fmt.Errorf("invalid U_id: %s", event.U_id)
	}

	E_ID, err := primitive.ObjectIDFromHex(event.E_id)
	if err != nil {
		return nil, fmt.Errorf("invalid E_id: %s", event.E_id)
	}

	// Check if user and event exist
	if err := checkExistence(ctx, memberCollection, U_ID, "user", event.U_id); err != nil {
		return nil, err
	}

	if err := checkExistence(ctx, eventCollection, E_ID, "event", event.E_id); err != nil {
		return nil, err
	}

	// Generate random string
	RandStr, err := service.RandomString(8)
	if err != nil {
		return nil, fmt.Errorf("error generating random string: %v", err)
	}

	// Create and insert QR code
	data := models.Qrcode{
		U_id:     event.U_id,
		E_id:     event.E_id,
		Title:    event.Title,
		Target:   event.Target,
		ShortUrl: RandStr,
		Status:   true,
	}

	event.ID = primitive.NewObjectID()
	result, err := qrcodeCollection.InsertOne(ctx, data)
	if err != nil {
		log.Println("InsertQrcode Error:", err)
		return nil, err
	}

	return result, nil
}

func UpdateQrcode(id string, qrcode models.Qrcode) (*mongo.UpdateResult, error) {
	collection := config.GetCollection("qrcode")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid ID:", err)
		return nil, err
	}

	update := bson.M{
		"$set": bson.M{
			"target": qrcode.Target,
			"title":  qrcode.Title,
			"status": qrcode.Status,
		},
	}

	result, err := collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		log.Println("Update Error:", err)
		return nil, err
	}

	return result, nil
}
func DeleteQrcode(id string) (*mongo.UpdateResult, error) {
	// soft delete
	collection := config.GetCollection("qrcode")

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

func checkExistence(ctx context.Context, collection *mongo.Collection, id primitive.ObjectID, entityType, idStr string) error {
	var entity interface{}
	if err := collection.FindOne(ctx, bson.M{"_id": id, "status": true}).Decode(&entity); err != nil {
		return fmt.Errorf("%s with ID %s does not exist", entityType, idStr)
	}
	return nil
}

func getEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("%s environment variable is not set", key)
	}
	return value, nil
}
