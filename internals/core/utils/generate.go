package utils

import (
	"github.com/HEEPOKE/generate-db/internals/domains/models/request"
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

func GenerateBatchData(size int64, generateRequest *request.GenerateRequest) []map[string]interface{} {
	results := make([]map[string]interface{}, size)

	for j := int64(0); j < size; j++ {
		rowData := make(map[string]interface{})

		for columnName, columnOptions := range generateRequest.Columns {
			defaultValue := columnOptions.Default
			if defaultValue == "" {
				category := columnOptions.Category
				switch category {
				case enums.NAME:
					rowData[columnName] = GenerateRandomNames()
				case enums.TEL:
					rowData[columnName] = GenerateRandomPhones()
				default:
					rowData[columnName] = nil
				}
			} else {
				rowData[columnName] = defaultValue
			}
		}

		results[j] = rowData
	}

	return results
}
