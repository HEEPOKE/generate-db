package handlers

import (
	"net/http"

	"github.com/HEEPOKE/generate-db/internals/app/helpers"
	"github.com/HEEPOKE/generate-db/internals/app/services"
	"github.com/HEEPOKE/generate-db/internals/core/utils"
	"github.com/HEEPOKE/generate-db/internals/domains/models"
	"github.com/HEEPOKE/generate-db/internals/domains/models/request"
	"github.com/HEEPOKE/generate-db/internals/domains/models/response"
	"github.com/HEEPOKE/generate-db/pkg/constants"
	"github.com/labstack/echo/v4"
)

type GenerateHandler struct {
	generateService services.GenerateService
}

func NewGenerateHandler(generateService services.GenerateService) *GenerateHandler {
	return &GenerateHandler{generateService: generateService}
}

type GenerateRequest struct {
	Table    string                            `json:"table" validate:"required" example:"users"`
	Columns  map[string]models.GenerateOptions `json:"columns" validate:"required" swaggo:"true,Example=map[string]models.GenerateOptions{'name': {Default: '', Types: 'string', Length: 0, AutoGenerate: true}, 'tel': {Default: '0999999999', Types: 'string', Length: 10, AutoGenerate: false}}"`
	Quantity int64                             `json:"quantity" validate:"required" example:"1000"`
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
		return helpers.FailResponse(c, err, constants.ERR_GET_LIST_GENERATE, http.StatusInternalServerError)
	}

	status := response.StatusMessage{
		Code:        constants.STATUS_OK,
		Message:     constants.MESSAGE_SUCCESS,
		Service:     constants.SERVICE_GET_LIST_ALL_GENERATE,
		Description: constants.DESCRIPTION_GET_LIST_ALL_GENERATE_SUCCESS,
	}

	response := response.ResponseMessage{
		Status:  status,
		Payload: users,
	}

	return helpers.SuccessResponse(c, http.StatusOK, response)
}

// Post Auto Generate Mock up data
// @Summary Get Auto Generate Mock up data
// @Description Get Auto Generate Mock up data
// @Tags Generate
// @Accept application/json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /generate/mockup-data [post]
// @param Body body handlers.GenerateRequest true "GenerateRequestBody"
func (gh *GenerateHandler) MockupData(c echo.Context) error {
	var req request.GenerateRequest

	err := c.Bind(&req)
	if err != nil {
		return helpers.FailResponse(c, err, constants.ERR_DECODE_DATA, http.StatusBadRequest)
	}

	key, err := utils.GenerateRandomKey()
	if err != nil {
		return helpers.FailResponse(c, err, constants.ERR_GENERATE_KEY, http.StatusInternalServerError)
	}

	generateData := models.Generate{
		Key:      key,
		Table:    req.Table,
		Quantity: req.Quantity,
	}

	err = gh.generateService.SaveDetailsGenerate(&generateData)
	if err != nil {
		return helpers.FailResponse(c, err, constants.ERR_SAVE_DETAILS_GENERATE, http.StatusInternalServerError)
	}

	result, err := gh.generateService.GenerateData(key, &req)
	if err != nil {
		return helpers.FailResponse(c, err, constants.ERR_GENERATE_DATA, http.StatusInternalServerError)
	}

	status := response.StatusMessage{
		Code:        constants.STATUS_CREATED,
		Message:     constants.MESSAGE_SUCCESS,
		Service:     constants.SERVICE_MOCKUP_DATA,
		Description: constants.DESCRIPTION_MOCKUP_DATA_SUCCESS,
	}

	response := response.ResponseMessage{
		Status:  status,
		Payload: result,
	}

	return helpers.SuccessResponse(c, http.StatusOK, response)
}
