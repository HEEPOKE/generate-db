package handlers

import "github.com/HEEPOKE/generate-db/internals/app/services"

type UtilitiesHandler struct {
	utilitiesService services.UtilitiesService
}

func NewUtilitiesHandler(utilitiesService services.UtilitiesService) *UtilitiesHandler {
	return &UtilitiesHandler{utilitiesService: utilitiesService}
}
