package services

import "github.com/HEEPOKE/generate-db/internals/core/interfaces"

type UtilitiesService struct {
	utilitiesRepository interfaces.UtilitiesRepository
}

func NewUtilitiesService(utilitiesRepository interfaces.UtilitiesRepository) *UtilitiesService {
	return &UtilitiesService{utilitiesRepository: utilitiesRepository}
}
