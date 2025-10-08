package application

import (
	"go_scenario_improver/controllers"
	"go_scenario_improver/services"
	"go_scenario_improver/utils"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
)

type App struct {
	ScenarioController controllers.ScenarioController
	HealthController   controllers.HealthController
}

func NewApp() *App {
	v := validator.New()
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		panic("api key wasn't set")
	}
	networkManager, err := utils.NewNeworkManager(apiKey, &http.Client{})
	if err != nil {
		panic(err)
	}
	scenarioService := services.NewScenatrioService(*networkManager)
	scenarioController := controllers.NewScenarioController(scenarioService, v)
	healthController := controllers.NewHealthController()
	return &App{
		ScenarioController: scenarioController,
		HealthController:   healthController,
	}
}
