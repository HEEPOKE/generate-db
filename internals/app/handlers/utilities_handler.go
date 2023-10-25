package handlers

import (
	"net/http"

	"github.com/HEEPOKE/generate-db/internals/app/helpers"
	"github.com/HEEPOKE/generate-db/internals/app/services"
	"github.com/HEEPOKE/generate-db/internals/domains/models/response"
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
// @param id query string true "key"
// @Success 200 {object} map[string]interface{}
// @Router /generate/check-data [get]
func (uh *UtilitiesHandler) CheckKeyData(c echo.Context) error {
	param := c.QueryParam("key")

	data, err := uh.utilitiesService.CheckKeyData(param)
	if err != nil {
		return helpers.FailResponse(c, err, constants.ERR_GET_DATA_FROM_CACHE, http.StatusInternalServerError)
	}

	status := response.StatusMessage{
		Code:        constants.STATUS_OK,
		Message:     constants.MESSAGE_SUCCESS,
		Service:     constants.SERVICE_GET_DATA_KEY_FROM_CACHE,
		Description: constants.DESCRIPTION_GET_DATA_KEY_FROM_CACHE_SUCCESS,
	}

	response := response.ResponseMessage{
		Status:  status,
		Payload: data,
	}

	return helpers.SuccessResponse(c, http.StatusOK, response)
}
