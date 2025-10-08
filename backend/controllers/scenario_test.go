package controllers

import (
	"bytes"
	"encoding/json"
	"go_scenario_improver/dto"
	"go_scenario_improver/mocks"
	"go_scenario_improver/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGenerate_Success(t *testing.T) {
	e := echo.New()
	networkManagerMock := mocks.NetworkManagerSuccessMock{}
	service := services.NewScenatrioService(networkManagerMock)
	controller := NewScenarioController(service, validator.New())
	body := dto.GenerateScenarioRequest{
		Text: "mock prompt",
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("failed to marshal body: %v", err)
	}
	req := httptest.NewRequest(http.MethodPost, "/generate", bytes.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, controller.GenerateScenario(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		expected := dto.GenerateScenarioResponse{
			Scenario: "Test Response",
		}
		var resp dto.GenerateScenarioResponse
		err := json.NewDecoder(rec.Body).Decode(&resp)
		assert.NoError(t, err)
		assert.Equal(t, expected, resp)
	}
}

func TestGenerate_BadReqest(t *testing.T) {
	e := echo.New()
	networkManagerMock := mocks.NetworkManagerSuccessMock{}
	service := services.NewScenatrioService(networkManagerMock)
	controller := NewScenarioController(service, validator.New())
	body := dto.GenerateScenarioRequest{}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("failed to marshal body: %v", err)
	}
	req := httptest.NewRequest(http.MethodPost, "/generate", bytes.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, controller.GenerateScenario(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		var resp echo.Map
		err := json.NewDecoder(rec.Body).Decode(&resp)
		assert.NoError(t, err)

		msg, _ := resp["message"].(string)
		assert.Contains(t, msg, "error")

	}
}

func TestGenerate_ServiceError(t *testing.T) {
	e := echo.New()
	networkManagerMock := mocks.NetworkManagerErrorMock{}
	service := services.NewScenatrioService(networkManagerMock)
	controller := NewScenarioController(service, validator.New())
	body := dto.GenerateScenarioRequest{
		Text: "mock prompt",
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("failed to marshal body: %v", err)
	}
	req := httptest.NewRequest(http.MethodPost, "/generate", bytes.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, controller.GenerateScenario(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		var resp echo.Map
		err := json.NewDecoder(rec.Body).Decode(&resp)
		assert.NoError(t, err)

		msg, _ := resp["message"].(string)
		assert.Contains(t, msg, "error")

	}
}
