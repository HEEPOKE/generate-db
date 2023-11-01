package handlers

import (
	"net/http"

	"github.com/HEEPOKE/generate-db/internals/app/helpers"
	"github.com/HEEPOKE/generate-db/internals/app/services"
	"github.com/HEEPOKE/generate-db/internals/domains/models"
	"github.com/HEEPOKE/generate-db/pkg/constants"
	"github.com/labstack/echo/v4"
)

type UtilitiesHandler struct {
	utilitiesService services.UtilitiesService
}

func NewUtilitiesHandler(utilitiesService services.UtilitiesService) *UtilitiesHandler {
	return &UtilitiesHandler{utilitiesService: utilitiesService}
}

// Get Data From Key in Cache
// @Summary Get Data From Key in Cache
// @Description Get Data From Key in Cache
// @Tags Utilities
// @Accept application/json
// @Produce json
// @param id query string true "key" example(123)
// @Router /generate/check-data [get]
// @Success 200 {object} map[string]interface{}
func (uh *UtilitiesHandler) CheckKeyData(c echo.Context) error {
	param := c.QueryParam("key")

	data, err := uh.utilitiesService.CheckKeyData(param)
	if err != nil {
		responseData := models.FailMessage{
			Echo:        c,
			StatusCode:  http.StatusInternalServerError,
			Code:        constants.CODE_FAILED,
			Message:     constants.MESSAGE_FAIL,
			Service:     constants.SERVICE_UTILITIES,
			Description: err,
		}
		return helpers.FailResponse(responseData)
	}

	response := models.SuccessMessage{
		Echo:        c,
		StatusCode:  http.StatusOK,
		Code:        constants.CODE_SUCCESS,
		Message:     constants.MESSAGE_SUCCESS,
		Service:     constants.SERVICE_GET_DATA_KEY_FROM_CACHE,
		Description: constants.DESCRIPTION_GET_DATA_KEY_FROM_CACHE_SUCCESS,
		Payload:     data,
	}

	return helpers.SuccessResponse(response)
}
