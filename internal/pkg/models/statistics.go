package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Statistics struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Qr_id      string             `bson:"qr_id" json:"qr_id"`
	Ip_address string             `bson:"ip_address" json:"ip_address"`
	Time_stamp string             `bson:"time_stamp" json:"time_stamp"`
	Status     bool               `bson:"status" json:"status"`
}
