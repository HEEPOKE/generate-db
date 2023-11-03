package services

import (
	"github.com/HEEPOKE/generate-db/internals/core/interfaces"
	"github.com/HEEPOKE/generate-db/internals/domains/models"
)

type UtilitiesService struct {
	utilitiesRepository interfaces.UtilitiesRepository
}

func NewUtilitiesService(utilitiesRepository interfaces.UtilitiesRepository) *UtilitiesService {
	return &UtilitiesService{utilitiesRepository: utilitiesRepository}
}

func (u *UtilitiesService) CheckKeyData(key string) (models.JsonStructure, error) {
	return u.utilitiesRepository.CheckKeyData(key)
}

func (u *UtilitiesService) UpdatePathFileJson(table, key, path string) error {
	return u.utilitiesRepository.UpdatePathFileJson(table, key, path)
}
