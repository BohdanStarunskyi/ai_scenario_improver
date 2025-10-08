package router

import (
	"go_scenario_improver/application"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, app *application.App) {
	e.GET("/ping", app.HealthController.Ping)
	e.POST("/generate", app.ScenarioController.GenerateScenario)
}
