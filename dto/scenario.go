package dto

type GenerateScenarioRequest struct {
	Text string `json:"text" validate:"required"`
}

type GenerateScenarioResponse struct {
	Scenario string `json:"scenario" validate:"required"`
}
