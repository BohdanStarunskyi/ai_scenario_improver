package main

import (
	"errors"
	"go_scenario_improver/application"
	"go_scenario_improver/router"
	"log/slog"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Warn("No .env file found", "error", err)
	}
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}, latency=${latency_human}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	app := application.NewApp()
	if app == nil {
		slog.Error("Failed to initialize application")
		return
	}

	router.SetupRoutes(e, app)

	slog.Info("Starting server on port 8080...")
	if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("Failed to start server", "error", err, "port", 8080)
	}
}
