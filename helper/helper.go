package helper

import (
	"image/png"
	"os"
	"time"

	"math/rand"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

func IdCreate(Len int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, Len)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func QrCreator(link, id string) error {
	qrCode, err := qr.Encode(link, qr.M, qr.Auto)
	if err != nil {
		return err
	}

	qrCode, err = barcode.Scale(qrCode, 200, 200)
	if err != nil {
		return err
	}

	file, err := os.Create("assets/qr/" + id + ".png")
	if err != nil {
		return err
	}
	defer file.Close()

	err = png.Encode(file, qrCode)
	if err != nil {
		return err
	}

	return nil
}
