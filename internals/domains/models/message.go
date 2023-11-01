package models

import "github.com/labstack/echo/v4"

type FailMessage struct {
	Echo        echo.Context
	StatusCode  int
	Code        string
	Message     string
	Service     string
	Description error
}

type SuccessMessage struct {
	Echo        echo.Context
	StatusCode  int
	Code        string
	Message     string
	Service     string
	Description string
	Payload     interface{}
}
