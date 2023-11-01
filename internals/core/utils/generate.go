package utils

import (
	"math/rand"
	"strconv"
	"sync"

	"github.com/HEEPOKE/generate-db/internals/domains/models"
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

	source := rand.NewSource(GetTimeNowThai().UnixNano())
	random := rand.New(source)

	randomNumber := make([]byte, length)
	for i := 0; i < length; i++ {
		randomNumber[i] = validChars[random.Intn(len(validChars))]
	}

	return string(randomNumber)
}

func GenerateBool(genLean string) bool {
	checkBool, _ := strconv.ParseBool(genLean)
	return checkBool
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

func GenerateBatchData(key string, generateRequest *request.GenerateRequest) models.JsonStructure {
	results := make([]map[string]interface{}, generateRequest.Quantity)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for j := int64(0); j < generateRequest.Quantity; j++ {
		wg.Add(1)
		go func(index int64) {
			defer wg.Done()

			rowData := make(map[string]interface{})

			for columnName, columnOptions := range generateRequest.Columns {
				defaultValue := columnOptions.Default

				switch {
				case columnOptions.AutoGenerate:
					rowData[columnName] = GenerateRandomData(columnOptions.Types, columnOptions.Length)
				case defaultValue == "true" || (defaultValue == "false" && columnOptions.Types == enums.BOOL):
					rowData[columnName] = GenerateBool(defaultValue)
				case defaultValue != "" && columnOptions.AutoGenerate:
					rowData[columnName] = ConvertToColumnType(defaultValue, columnOptions.Types)
				default:
					rowData[columnName] = defaultValue
				}
			}

			mu.Lock()
			results[index] = rowData
			mu.Unlock()
		}(j)
	}

	wg.Wait()

	data := models.JsonStructure{
		Table:    generateRequest.Table,
		DataKey:  key,
		Quantity: generateRequest.Quantity,
		Datas:    results,
	}

	return data
}
