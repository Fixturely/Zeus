package subscriptions

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"zeus/pkg/application"
	"zeus/pkg/models"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateSubscriptionHandler(t *testing.T) {
	ctx := context.Background()
	app, err := application.NewApp(ctx)
	assert.NoError(t, err)

	// Create a mock HTTP server for MIKE API calls
	mikeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Mock responses based on the request path
		if r.URL.Path == "/sport/exists/1" || r.URL.Path == "/sport/exists/2" {
			_, _ = w.Write([]byte(`{"exists": true}`))
		} else if r.URL.Path == "/team/exists/1" {
			_, _ = w.Write([]byte(`{"exists": true}`))
		} else {
			_, _ = w.Write([]byte(`{"exists": false}`))
		}
	}))
	defer mikeServer.Close()

	// Override the MIKE API URL in config for testing
	// Note: This is a simplified approach. In production, you'd want to use dependency injection
	originalConfig := app.Config
	app.Config.MikeAPIKey = mikeServer.URL

	tests := []struct {
		name           string
		requestBody    map[string]interface{}
		userID         int
		expectedStatus int
		expectError    bool
	}{
		{
			name: "valid subscription request",
			requestBody: map[string]interface{}{
				"sport_id": 1,
				"team_id":  1,
			},
			userID:         1,
			expectedStatus: http.StatusCreated,
			expectError:    false,
		},
		{
			name: "valid subscription request without team_id",
			requestBody: map[string]interface{}{
				"sport_id": 2,
			},
			userID:         1,
			expectedStatus: http.StatusCreated,
			expectError:    false,
		},
		{
			name: "invalid request - missing sport_id",
			requestBody: map[string]interface{}{
				"team_id": 1,
			},
			userID:         1,
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
		{
			name: "invalid request - zero sport_id",
			requestBody: map[string]interface{}{
				"sport_id": 0,
				"team_id":  1,
			},
			userID:         1,
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
		{
			name:           "invalid request - empty body",
			requestBody:    map[string]interface{}{},
			userID:         1,
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Create a new Echo instance for each test
			e := echo.New()

			// Convert request body to JSON
			jsonBody, err := json.Marshal(test.requestBody)
			assert.NoError(t, err)

			// Create HTTP request
			req := httptest.NewRequest(http.MethodPost, "/subscriptions", bytes.NewReader(jsonBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			// Create HTTP response recorder
			rec := httptest.NewRecorder()

			// Create Echo context
			c := e.NewContext(req, rec)
			c.Set("user_id", test.userID)
			c.Set("app", app)

			// Call the handler
			handlerErr := CreateSubscriptionHandler(c)
			assert.NoError(t, handlerErr)

			// Check response status
			assert.Equal(t, test.expectedStatus, rec.Code)

			if test.expectError {
				// Check that we got an error response
				var errorResp map[string]string
				err = json.Unmarshal(rec.Body.Bytes(), &errorResp)
				assert.NoError(t, err)
				assert.Contains(t, errorResp, "error")
			} else {
				// Check that we got a successful response
				var subscription models.Subscription
				err = json.Unmarshal(rec.Body.Bytes(), &subscription)
				assert.NoError(t, err)
				assert.Equal(t, test.userID, subscription.UserID)
				assert.NotZero(t, subscription.ID)
			}
		})
	}

	// Restore original config
	app.Config = originalConfig
}
