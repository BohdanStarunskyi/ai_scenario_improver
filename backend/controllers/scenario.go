package controllers

import (
	"go_scenario_improver/dto"
	"go_scenario_improver/services"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ScenarioController interface {
	GenerateScenario(c echo.Context) error
}

type scenarioControllerImpl struct {
	Ss services.ScenarioService
	Va *validator.Validate
}

func NewScenarioController(ss services.ScenarioService, va *validator.Validate) ScenarioController {
	return &scenarioControllerImpl{
		Ss: ss,
		Va: va,
	}
}

func (s *scenarioControllerImpl) GenerateScenario(c echo.Context) error {
	var req dto.GenerateScenarioRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "error, check request"})
	}
	if err := s.Va.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "error " + err.Error()})
	}
	resp, err := s.Ss.GenerateScenario(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "error generating scenario, try again later"})
	}

	return c.JSON(http.StatusOK, resp)
}
