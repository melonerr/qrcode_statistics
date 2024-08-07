package main

import (
	"log"
	"os"
	"qrcode_statistics/internal/app"
)

func init() {
	// Get the hostname
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("Error when getting hostname: %s", err)
	}

	// Log the hostname
	log.Println("Hostname:", hostname)

	// Log application start with additional fields
	// Ensure that `log.InfoWithFields` is from a properly initialized custom logger
	InfoWithFields(
		"start app",
		Fields{
			"Action": "start app",
		},
	)
}

// Assuming you have a custom logger with InfoWithFields method
func InfoWithFields(message string, fields Fields) {
	log.Println(message, fields)
}

type Fields map[string]interface{}

func main() {
	app.Start()
}
