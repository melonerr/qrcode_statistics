package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Statistics struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Qr_id      string             `bson:"qr_id" json:"qr_id"`
	Event_id   string             `bson:"event_id" json:"event_id"`
	Ip_address string             `bson:"ip_address" json:"ip_address"`
	Location   string             `bson:"location" json:"location"`
	Time_stamp time.Time          `bson:"time_stamp" json:"time_stamp"`
}
