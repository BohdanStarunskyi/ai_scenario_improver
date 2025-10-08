package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var contoller = NewHealthController()

func TestPing_Success(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, contoller.Ping(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		expected := `{"message":"pong"}`
		assert.JSONEq(t, expected, rec.Body.String())
	}
}
