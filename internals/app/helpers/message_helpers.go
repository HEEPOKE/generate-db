package helpers

import (
	"github.com/HEEPOKE/generate-db/internals/domains/models/response"
	"github.com/labstack/echo/v4"
)

func FailResponse(c echo.Context, err error, msg string, status int) error {
	if err != nil {
		response := response.ErrorResponse{
			Message: msg,
			Error:   err.Error(),
		}
		return c.JSON(status, response)
	}
	return nil
}

func SuccessResponse(c echo.Context, ResponseMessage *response.ResponseMessage) error {
	if ResponseMessage != nil {
		return c.JSON(ResponseMessage.Code, ResponseMessage)
	}
	return nil
}
