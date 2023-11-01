package helpers

import (
	"github.com/HEEPOKE/generate-db/internals/domains/models"
	"github.com/HEEPOKE/generate-db/internals/domains/models/response"
)

func FailResponse(responseData models.FailMessage) error {
	status := response.StatusMessage{
		Code:        responseData.Code,
		Message:     responseData.Message,
		Service:     responseData.Service,
		Description: responseData.Description.Error(),
	}

	response := response.ResponseMessage{
		Status: status,
	}

	return responseData.Echo.JSON(responseData.StatusCode, response)
}

func SuccessResponse(responseData models.SuccessMessage) error {
	status := response.StatusMessage{
		Code:        responseData.Code,
		Message:     responseData.Message,
		Service:     responseData.Service,
		Description: responseData.Description,
	}

	response := response.ResponseMessage{
		Status:  status,
		Payload: responseData.Payload,
	}

	return responseData.Echo.JSON(responseData.StatusCode, response)
}
