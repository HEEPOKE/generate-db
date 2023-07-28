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

func GenerateRandomFirstName() string {
	firstName := gofakeit.FirstName()

	return firstName
}

func GenerateRandomLastName() string {
	lastName := gofakeit.LastName()

	return lastName
}

func GenerateRandomEmail() string {
	email := gofakeit.Email()

	return email
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

func GenerateRandomData(dataType enums.Category, length int) interface{} {
	switch dataType {
	case enums.NAME:
		return GenerateRandomNames()
	case enums.FIRSTNAME:
		return GenerateRandomFirstName()
	case enums.LASTNAME:
		return GenerateRandomLastName()
	case enums.EMAIL:
		return GenerateRandomEmail()
	case enums.TEL:
		return GenerateRandomPhones()
	case enums.PASSWORD:
		return GeneratePassword(length)
	case enums.WORD:
		return GenerateWord()
	case enums.SENTENCE:
		return GenerateSentence(length)
	default:
		return nil
	}
}

func GenerateBatchData(size int64, key string, generateRequest *request.GenerateRequest) []map[string]interface{} {
	results := make([]map[string]interface{}, size)

	for j := int64(0); j < size; j++ {
		rowData := make(map[string]interface{})

		for columnName, columnOptions := range generateRequest.Columns {
			defaultValue := columnOptions.Default

			if columnOptions.AutoGenerate {
				rowData[columnName] = GenerateRandomData(columnOptions.Types, columnOptions.Length)
			} else if defaultValue == "true" || defaultValue == "false" && columnOptions.Types == enums.BOOL {
				rowData[columnName] = GenerateBool(defaultValue)
			} else if defaultValue != "" && columnOptions.AutoGenerate {
				rowData[columnName] = ConvertToColumnType(defaultValue, columnOptions.Types)
			} else {
				rowData[columnName] = defaultValue
			}
		}

		results[j] = rowData
	}

	return results
}
