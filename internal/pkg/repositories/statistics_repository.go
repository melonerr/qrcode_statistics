package repositories

import (
	"context"
	"log"
	"qrcode_statistics/internal/config"
	"qrcode_statistics/internal/pkg/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddStatistics(data models.Statistics) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	Collection := config.GetCollection("statistics")

	// Create and insert QR code
	data = models.Statistics{
		Qr_id:      data.Qr_id,
		Ip_address: data.Ip_address,
		Time_stamp: data.Time_stamp,
		Status:     true,
	}

	data.ID = primitive.NewObjectID()
	result, err := Collection.InsertOne(ctx, data)
	if err != nil {
		log.Println("InsertQrcode Error:", err)
		return nil, err
	}

	return result, nil
}
