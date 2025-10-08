package controllers

import "github.com/labstack/echo/v4"

type HealthController interface {
	Ping(c echo.Context) error
}

type healthControllerImpl struct{}

func NewHealthController() HealthController {
	return &healthControllerImpl{}
}

func (h *healthControllerImpl) Ping(c echo.Context) error {
	return c.JSON(200, echo.Map{"message": "pong"})
}
