package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Qrcode struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	U_id     string             `bson:"u_id" json:"u_id"`
	E_id     string             `bson:"E_id" json:"E_id"`
	Title    string             `bson:"title" json:"title"`
	Target   string             `bson:"target" json:"target"`
	ShortUrl string             `bson:"shortUrl" json:"shortUrl"`
	Status   bool               `bson:"status" json:"status"`
}

type QrcodeRes struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title    string             `bson:"title" json:"title"`
	Target   string             `bson:"target" json:"target"`
	ShortUrl string             `bson:"shortUrl" json:"shortUrl"`
	Qrcode   string             `bson:"qrcode" json:"qrcode"`
	Status   bool               `bson:"status" json:"status"`
}
