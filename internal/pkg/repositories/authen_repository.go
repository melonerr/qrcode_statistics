package repositories

import (
	"context"
	"log"
	"time"

	"qrcode_statistics/internal/config"
	"qrcode_statistics/internal/pkg/models"
	"qrcode_statistics/internal/pkg/service"

	"github.com/golang-jwt/jwt/v5"

	"go.mongodb.org/mongo-driver/bson"
)

func Authen(user models.Authen) (string, error) {
	collection := config.GetCollection("members")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var userData models.Members
	err := collection.FindOne(ctx, bson.M{"username": user.Username, "password": service.GenerateMD5Hash(user.Password)}).Decode(&userData)
	if err != nil {
		log.Println("Username and password do not match.")
		return "", err
	}

	claims := jwt.MapClaims{
		"id":    userData.ID,
		"name":  userData.Username,
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtKey, err := getEnv("SECRET_KEY")
	if err != nil {
		return "", err
	}

	signedToken, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
