package handlers

import (
	"net/http"
	"time"

	"github.com/HEEPOKE/generate-db/internals/app/helpers"
	"github.com/HEEPOKE/generate-db/internals/app/services"
	"github.com/HEEPOKE/generate-db/internals/core/utils"
	"github.com/HEEPOKE/generate-db/internals/domains/models"
	"github.com/HEEPOKE/generate-db/internals/domains/models/request"
	"github.com/HEEPOKE/generate-db/pkg/constants"
	"github.com/labstack/echo/v4"
)

type GenerateHandler struct {
	generateService services.GenerateService
}

func NewGenerateHandler(generateService services.GenerateService) *GenerateHandler {
	return &GenerateHandler{generateService: generateService}
}

// Get List All Generate Data
// @Summary Get List All Generate Data
// @Description Get List All Generate Data
// @Tags Generate
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /generate/get-details [get]
func (gh *GenerateHandler) GetListGenerateAll(c echo.Context) error {
	users, err := gh.generateService.GetGenerateAll()
	if err != nil {
		responseData := models.FailMessage{
			Echo:        c,
			StatusCode:  http.StatusInternalServerError,
			Code:        constants.CODE_FAILED,
			Message:     constants.MESSAGE_FAIL,
			Service:     constants.SERVICE_GENERATE,
			Description: err,
		}
		return helpers.FailResponse(responseData)
	}

	response := models.SuccessMessage{
		Echo:        c,
		StatusCode:  http.StatusOK,
		Code:        constants.CODE_SUCCESS,
		Message:     constants.MESSAGE_SUCCESS,
		Service:     constants.SERVICE_GENERATE,
		Description: constants.DESCRIPTION_GET_LIST_ALL_GENERATE_SUCCESS,
		Payload:     users,
	}

	return helpers.SuccessResponse(response)
}

// Post Auto Generate Mock up data
// @Summary Get Auto Generate Mock up data
// @Description Get Auto Generate Mock up data
// @Tags Generate
// @Accept application/json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /generate/mockup-data [post]
// @param Body body examples.GenerateExample true "GenerateRequestBody"
func (gh *GenerateHandler) MockupData(c echo.Context) error {
	var req request.GenerateRequest
	var key string

	err := c.Bind(&req)
	if err != nil {
		responseData := models.FailMessage{
			Echo:        c,
			StatusCode:  http.StatusInternalServerError,
			Code:        constants.CODE_FAILED,
			Message:     constants.MESSAGE_FAIL,
			Service:     constants.SERVICE_GENERATE,
			Description: err,
		}
		return helpers.FailResponse(responseData)
	}

	if req.DataKey == "" {
		key, err = utils.GenerateRandomKey()
		if err != nil {
			responseData := models.FailMessage{
				Echo:        c,
				StatusCode:  http.StatusInternalServerError,
				Code:        constants.CODE_FAILED,
				Message:     constants.MESSAGE_FAIL,
				Service:     constants.SERVICE_GENERATE,
				Description: err,
			}
			return helpers.FailResponse(responseData)
		}
	} else {
		key = req.DataKey
	}

	generateData := models.Generate{
		Key:         key,
		Table:       req.Table,
		Quantity:    req.Quantity,
		TimeExpired: utils.GetTimeNowThai().Add(24 * time.Hour),
	}

	err = gh.generateService.SaveDetailsGenerate(&generateData)
	if err != nil {
		responseData := models.FailMessage{
			Echo:        c,
			StatusCode:  http.StatusInternalServerError,
			Code:        constants.CODE_FAILED,
			Message:     constants.MESSAGE_FAIL,
			Service:     constants.SERVICE_GENERATE,
			Description: err,
		}
		return helpers.FailResponse(responseData)
	}

	result, err := gh.generateService.GenerateData(generateData, &req)
	if err != nil {
		responseData := models.FailMessage{
			Echo:        c,
			StatusCode:  http.StatusInternalServerError,
			Code:        constants.CODE_FAILED,
			Message:     constants.MESSAGE_FAIL,
			Service:     constants.SERVICE_GENERATE,
			Description: err,
		}
		return helpers.FailResponse(responseData)
	}

	response := models.SuccessMessage{
		Echo:        c,
		StatusCode:  http.StatusOK,
		Code:        constants.CODE_SUCCESS,
		Message:     constants.MESSAGE_SUCCESS,
		Service:     constants.SERVICE_MOCKUP_DATA,
		Description: constants.DESCRIPTION_MOCKUP_DATA_SUCCESS,
		Payload:     result,
	}

	return helpers.SuccessResponse(response)
}
