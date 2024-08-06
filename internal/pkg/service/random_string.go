package service

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"strings"
)

func RandomString(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("length must be greater than 0")
	}

	// Calculate the number of bytes needed for the specified length
	numBytes := (length * 3) / 4

	// Generate random bytes
	bytes := make([]byte, numBytes)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	// Encode bytes to a base64 string and trim to the desired length
	randomString := base64.URLEncoding.EncodeToString(bytes)
	return strings.TrimRight(randomString[:length], "="), nil
}
