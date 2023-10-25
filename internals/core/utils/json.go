package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/HEEPOKE/generate-db/internals/domains/models"
)

func CreateJSONFile(data models.JsonStructure, details models.Generate) error {
	currentTime := GetTimeNowThai()
	formattedTime := currentTime.Format("02-01-2006")

	dirPath := fmt.Sprintf("data/%s/", formattedTime)

	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			log.Printf("Failed to create directory: %v", err)
			return err
		}
	}

	filePath := fmt.Sprintf("%s%s_%s_%d.json", dirPath, details.Table, details.Key, details.Quantity)

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
