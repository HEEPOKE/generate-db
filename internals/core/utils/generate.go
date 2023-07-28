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

func GenerateWord() string {
	gofakeit.Seed(0)

	randomWord := gofakeit.Word()

	return randomWord
}

func GenerateSentence(length int) string {
	randomWord := gofakeit.Sentence(length)

	return randomWord
}

func GenerateBatchData(size int64, key string, generateRequest *request.GenerateRequest) []map[string]interface{} {
	results := make([]map[string]interface{}, size)

	for j := int64(0); j < size; j++ {
		rowData := make(map[string]interface{})

		for columnName, columnOptions := range generateRequest.Columns {
			defaultValue := columnOptions.Default
			if defaultValue == "" && columnOptions.AutoGenerate {
				category := columnOptions.Types
				switch category {
				case enums.NAME:
					rowData[columnName] = GenerateRandomNames()
				case enums.TEL:
					rowData[columnName] = GenerateRandomPhones()
				case enums.PASSWORD:
					rowData[columnName] = GeneratePassword(columnOptions.Length)
				case enums.WORD:
					rowData[columnName] = GenerateWord()
				case enums.SENTENCE:
					rowData[columnName] = GenerateSentence(columnOptions.Length)
				default:
					rowData[columnName] = nil
				}
			} else if defaultValue == "true" || defaultValue == "false" && columnOptions.Types == enums.BOOL {
				rowData[columnName] = GenerateBool(defaultValue)
			} else if defaultValue != "" && columnOptions.AutoGenerate {
				category := columnOptions.Types
				switch category {
				case enums.FLOAT32:
					rowData[columnName] = ConvertToFloat32(defaultValue)
				case enums.FLOAT64:
					rowData[columnName] = ConvertToFloat64(defaultValue)
				case enums.UINT:
					rowData[columnName] = ConvertToUint(defaultValue)
				case enums.UINT8:
					rowData[columnName] = ConvertToUint8(defaultValue)
				case enums.UINT16:
					rowData[columnName] = ConvertToUint16(defaultValue)
				case enums.UINT32:
					rowData[columnName] = ConvertToUint32(defaultValue)
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
