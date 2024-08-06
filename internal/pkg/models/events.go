package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Events struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	U_id       string             `bson:"u_id" json:"u_id"`
	Title      string             `bson:"title" json:"title"`
	Detail     string             `bson:"detail" json:"detail"`
	Date_start string             `bson:"date_start" json:"date_start"`
	Date_end   string             `bson:"date_end" json:"date_end"`
	Status     bool               `bson:"status" json:"status"`
}
