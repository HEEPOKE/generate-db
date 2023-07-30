package handlers

import "github.com/HEEPOKE/generate-db/internals/app/services"

type InsertHandler struct {
	insertService services.InsertService
}

func NewInsertHandler(insertService services.InsertService) *InsertHandler {
	return &InsertHandler{insertService: insertService}
}
