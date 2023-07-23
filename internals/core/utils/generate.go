package utils

import "github.com/brianvoe/gofakeit/v6"

func GenerateRandomNames(quantity int) []string {
	var names []string

	for i := 0; i < quantity; i++ {
		name := gofakeit.Name()
		names = append(names, name)
	}

	return names
}

func GenerateRandomPhones(quantity int) []string {
	var phones []string

	for i := 0; i < quantity; i++ {
		phone := "0" + gofakeit.PhoneFormatted()[1:]
		phones = append(phones, phone)
	}

	return phones
}
