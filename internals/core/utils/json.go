package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func CreateJSONFile(data interface{}, key string, size int64, fileNumber int) error {
	currentTime := time.Now()
	formattedTime := currentTime.Format("02-01-2006")

	filePath := fmt.Sprintf("./data/%s_%s/%d_%d.json", formattedTime, key, fileNumber, size)

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("เกิดข้อผิดพลาดในการแปลงข้อมูลเป็น JSON: %w", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("เกิดข้อผิดพลาดในการสร้างไฟล์: %w", err)
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		return fmt.Errorf("เกิดข้อผิดพลาดในการเขียนข้อมูลลงไฟล์: %w", err)
	}

	return nil
}
