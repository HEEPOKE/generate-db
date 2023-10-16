package utils

import (
	"crypto/rand"
	"encoding/base64"
	"time"
)

func GenerateRandomKey() (string, error) {
	keyLength := 8

	adjustedKeyLength := (keyLength + 3) / 4 * 4

	randomBytes := make([]byte, adjustedKeyLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	key := base64.RawURLEncoding.EncodeToString(randomBytes)

	key = key[:keyLength]

	return key, nil
}

func GetTimeNowThai() time.Time {
	bangkokOffset := 7 * 60 * 60
	currentTime := time.Now().UTC().Add(time.Second * time.Duration(bangkokOffset))

	return currentTime
}
