package models

type Authen struct {
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
}

type Token struct {
	Token string `bson:"token" json:"token"`
}
