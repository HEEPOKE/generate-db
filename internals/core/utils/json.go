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

	filePath := fmt.Sprintf("./data/%s_%s/%d_%d.json", formattedTime, key, fileNumber, size)

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("เกิดข้อผิดพลาดในการแปลงข้อมูลเป็น JSON: %v", err)
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		log.Printf("เกิดข้อผิดพลาดในการสร้างไฟล์: %v", err)
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		log.Printf("เกิดข้อผิดพลาดในการเขียนข้อมูลลงไฟล์: %v", err)
		return err
	}

	return nil
}
