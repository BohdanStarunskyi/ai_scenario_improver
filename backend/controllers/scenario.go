package controllers

import (
	"go_scenario_improver/dto"
	"go_scenario_improver/services"
	"log/slog"
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
		slog.Error("Failed to bind request", "error", err, "ip", c.RealIP())
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid request format"})
	}

	if err := s.Va.Struct(req); err != nil {
		slog.Warn("Request validation failed", "error", err, "request", req, "ip", c.RealIP())
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Validation error: " + err.Error()})
	}

	resp, err := s.Ss.GenerateScenario(req)
	if err != nil {
		slog.Error("Failed to generate scenario", "error", err, "request", req, "ip", c.RealIP())
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error generating scenario, please try again later"})
	}

	return c.JSON(http.StatusOK, resp)
}
