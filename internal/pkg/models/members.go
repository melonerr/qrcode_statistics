package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Members struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Token    string             `bson:"token" json:"token"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"password"`
	Email    string             `bson:"email" json:"email"`
	Role     string             `bson:"role" json:"role"`
	Status   bool               `bson:"status" json:"status"`
}
