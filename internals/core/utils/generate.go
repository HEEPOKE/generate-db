package utils

import (
	"math/rand"
	"time"

	"github.com/HEEPOKE/generate-db/internals/domains/models/request"
	"github.com/HEEPOKE/generate-db/pkg/config"
	"github.com/HEEPOKE/generate-db/pkg/enums"
	"github.com/brianvoe/gofakeit/v6"
)

func GenerateRandomNames() string {
	name := gofakeit.Name()

	return name
}

func GenerateRandomPhones() string {
	phone := "0" + gofakeit.PhoneFormatted()[1:]

	return phone
}

func GeneratePassword(length int) string {
	var validChars string
	validChars += config.UppercaseLetters
	validChars += config.LowercaseLetters
	validChars += config.Numbers
	validChars += config.SpecialChars

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	password := make([]byte, length)
	for i := 0; i < length; i++ {
		password[i] = validChars[random.Intn(len(validChars))]
	}

	return string(password)
}

func GenerateBool(Bool string) bool {
	return Bool == "true"
}

func GenerateBatchData(size int64, generateRequest *request.GenerateRequest) []map[string]interface{} {
	results := make([]map[string]interface{}, size)

	for j := int64(0); j < size; j++ {
		rowData := make(map[string]interface{})

		for columnName, columnOptions := range generateRequest.Columns {
			defaultValue := columnOptions.Default
			if defaultValue == "" && !columnOptions.AutoGenerate {
				category := columnOptions.Category
				switch category {
				case enums.NAME:
					rowData[columnName] = GenerateRandomNames()
				case enums.TEL:
					rowData[columnName] = GenerateRandomPhones()
				case enums.PASSWORD:
					rowData[columnName] = GeneratePassword(columnOptions.Length)
				default:
					rowData[columnName] = nil
				}
			} else if defaultValue == "true" || defaultValue == "false" {
				rowData[columnName] = GenerateBool(defaultValue)
			} else {
				rowData[columnName] = defaultValue
			}
		}

		results[j] = rowData
	}

	return results
}
