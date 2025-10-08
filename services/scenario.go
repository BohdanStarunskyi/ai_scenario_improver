package services

import (
	"go_scenario_improver/dto"
	"go_scenario_improver/utils"
)

type ScenarioService interface {
	GenerateScenario(req dto.GenerateScenarioRequest) (*dto.GenerateScenarioResponse, error)
}

type ScenarioServiceImpl struct {
	Req utils.Requester
}

func NewScenatrioService(req utils.Requester) ScenarioService {
	return &ScenarioServiceImpl{
		Req: req,
	}
}

func (s *ScenarioServiceImpl) GenerateScenario(req dto.GenerateScenarioRequest) (*dto.GenerateScenarioResponse, error) {
	resp, err := s.Req.SendRequest(req.Text)
	if err != nil {
		return nil, err
	}
	return &dto.GenerateScenarioResponse{Scenario: resp}, nil
}
