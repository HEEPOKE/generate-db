package utils

import (
	"crypto/rand"
	"encoding/base64"
	"os"
	"strconv"
	"strings"
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

func FindNextFileNumber(key string) (int, error) {
	dirPath := "./data/"
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return 0, err
	}

	maxFileNumber := 0
	for _, file := range files {
		if file.IsDir() {
			name := file.Name()
			if strings.HasPrefix(name, key+"_") {
				fileNumberStr := strings.TrimPrefix(name, key+"_")
				fileNumber, err := strconv.Atoi(fileNumberStr)
				if err == nil && fileNumber > maxFileNumber {
					maxFileNumber = fileNumber
				}
			}
		}
	}

	return maxFileNumber + 1, nil
}
