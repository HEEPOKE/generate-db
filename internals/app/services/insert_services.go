package services

import "github.com/HEEPOKE/generate-db/internals/core/interfaces"

type InsertService struct {
	insertRepository interfaces.InsertRepository
}

func NewInsertService(insertRepository interfaces.InsertRepository) *InsertService {
	return &InsertService{insertRepository: insertRepository}
}
