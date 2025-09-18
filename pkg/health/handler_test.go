package health

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	// Create echo instance
	e := echo.New()

	// Create handler
	handler := NewHandler()

	// Register route
	handler.RegisterRoutes(e)

	// Create request
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	// Perform request
	e.ServeHTTP(rec, req)

	// Assertions
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "ok")
	assert.Contains(t, rec.Body.String(), "Zeus API is healthy")
}

func TestHealthCheckWithDifferentEnvironments(t *testing.T) {
	tests := []struct {
		name     string
		env      string
		expected string
	}{
		{
			name:     "development environment",
			env:      "development",
			expected: "ok",
		},
		{
			name:     "test environment",
			env:      "test",
			expected: "ok",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environment
			t.Setenv("ENV", tt.env)

			// Create echo instance
			e := echo.New()

			// Create handler
			handler := NewHandler()

			// Register route
			handler.RegisterRoutes(e)

			// Create request
			req := httptest.NewRequest(http.MethodGet, "/health", nil)
			rec := httptest.NewRecorder()

			// Perform request
			e.ServeHTTP(rec, req)

			// Assertions
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Contains(t, rec.Body.String(), tt.expected)
		})
	}
}
