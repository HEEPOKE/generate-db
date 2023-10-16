package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

func CreateJSONFile(data interface{}, key string, size int64, fileNumber int) error {
	currentTime := time.Now()
	formattedTime := currentTime.Format("02-01-2006")

	dirPath := fmt.Sprintf("../../../data/%s_%s/", formattedTime, key)

	if err := os.MkdirAll(dirPath, 0755); err != nil {
		log.Printf("Failed to create directory: %v", err)
		return err
	}

	filePath := fmt.Sprintf("%s/%d_%d.json", dirPath, fileNumber, size)

	file, err := os.Create(filePath)
	if err != nil {
		log.Printf("เกิดข้อผิดพลาดในการสร้างไฟล์: %v", err)
		return err
	}
	defer file.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("เกิดข้อผิดพลาดในการแปลงข้อมูลเป็น JSON: %v", err)
		return err
	}

	_, err = file.Write(jsonData)
	if err != nil {
		log.Printf("เกิดข้อผิดพลาดในการเขียนข้อมูลลงไฟล์: %v", err)
		return err
	}

	return nil
}
