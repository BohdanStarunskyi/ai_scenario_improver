package services

import (
	"go_scenario_improver/dto"
	"go_scenario_improver/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSceario_Success(t *testing.T) {
	var networkManagerMock = mocks.NetworkManagerSuccessMock{}
	var service = NewScenatrioService(networkManagerMock)
	resp, err := service.GenerateScenario(dto.GenerateScenarioRequest{Text: "test"})
	if assert.NoError(t, err) {
		assert.Equal(t, "Test Response", resp.Scenario)
	}
}

func TestSceario_Error(t *testing.T) {
	var networkManagerMock = mocks.NetworkManagerErrorMock{}
	var service = NewScenatrioService(networkManagerMock)
	resp, err := service.GenerateScenario(dto.GenerateScenarioRequest{Text: "test"})
	assert.Nil(t, resp)
	assert.EqualError(t, err, "Test Error")
}
