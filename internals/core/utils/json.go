package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

func CreateJSONFile(data interface{}, size int64, table, key string) error {
	currentTime := time.Now()
	formattedTime := currentTime.Format("02-01-2006")

	dirPath := fmt.Sprintf("data/%s/", formattedTime)

	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			log.Printf("Failed to create directory: %v", err)
			return err
		}
	}

	filePath := fmt.Sprintf("%s%s_%s_%d.json", dirPath, table, key, size)

	file, err := os.Create(filePath)
	if err != nil {
		log.Printf("เกิดข้อผิดพลาดในการสร้างไฟล์: %v", err)
		return err
	}
	defer file.Close()

	jsonData, err := json.MarshalIndent(data, "", "  ")
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
