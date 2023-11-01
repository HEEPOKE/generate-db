package examples

import "github.com/HEEPOKE/generate-db/internals/domains/models"

type SuccessCheckDataFromKeyResponse struct {
	Status  SuccessStatusMessage `json:"status"`
	PayLoad models.JsonStructure `json:"payload"`
}

type SuccessGenerateGetAllResponse struct {
	Status  SuccessStatusMessage    `json:"status"`
	PayLoad GenerateResponseExample `json:"payload"`
}
