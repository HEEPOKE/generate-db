package handlers

import (
	"net/http"

	"github.com/HEEPOKE/generate-db/internals/app/helpers"
	"github.com/HEEPOKE/generate-db/internals/app/services"
	"github.com/HEEPOKE/generate-db/pkg/constants"
	"github.com/labstack/echo/v4"
)

type GenerateHandler struct {
	generateService services.GenerateService
}

func NewGenerateHandler(generateService services.GenerateService) *GenerateHandler {
	return &GenerateHandler{generateService: generateService}
}

func (gh *GenerateHandler) GetListGenerateAll(c echo.Context) error {
	users, err := gh.generateService.GetGenerateAll()
	if err != nil {
		return helpers.FailResponse(c, err, constants.ERR_GET_LIST_GENERATE, http.StatusInternalServerError)
	}

	status := response.StatusMessage{
Code: http.StatusOK,
Message:"success",
Service: "",
Description: "",
	}

	response := response.ResponseMessage{
		Status: status,
		Payload:   users,
	}

	return helpers.SuccessResponse(c, response)
}
