package application

import (
	"go_scenario_improver/controllers"
	"go_scenario_improver/services"
	"go_scenario_improver/utils"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
)

type App struct {
	ScenarioController controllers.ScenarioController
	HealthController   controllers.HealthController
}

func NewApp() *App {
	slog.Info("Initializing application dependencies...")

	v := validator.New()
	slog.Info("Validator initialized")

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		slog.Error("API_KEY environment variable is not set")
		panic("API key wasn't set - please check your .env file")
	}
	slog.Info("API key loaded from environment")

	networkManager, err := utils.NewNeworkManager(apiKey, &http.Client{})
	if err != nil {
		slog.Error("Failed to initialize network manager", "error", err)
		panic(err)
	}
	slog.Info("Network manager initialized")

	scenarioService := services.NewScenatrioService(*networkManager)
	slog.Info("Scenario service initialized")

	scenarioController := controllers.NewScenarioController(scenarioService, v)
	slog.Info("Scenario controller initialized")

	healthController := controllers.NewHealthController()
	slog.Info("Health controller initialized")

	slog.Info("Application initialized successfully")
	return &App{
		ScenarioController: scenarioController,
		HealthController:   healthController,
	}
}
