package service

import (
	"bytes"
	"encoding/base64"
	"image/png"
	"log"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func QecodeGenerator(data string) (string, error) {
	// Generate QR Code

	qrCode, err := qr.Encode(data, qr.L, qr.Auto)
	if err != nil {
		log.Fatalf("Failed to generate QR code: %v", err)
	}

	qrCode, err = barcode.Scale(qrCode, 256, 256)
	if err != nil {
		log.Fatalf("Failed to scale QR code: %v", err)
	}

	// Convert QR Code to PNG
	var pngBuffer bytes.Buffer
	err = png.Encode(&pngBuffer, qrCode)
	if err != nil {
		log.Fatalf("Failed to encode QR code to PNG: %v", err)
	}

	// Encode PNG to Base64
	base64PNG := base64.StdEncoding.EncodeToString(pngBuffer.Bytes())

	return base64PNG, err
}
