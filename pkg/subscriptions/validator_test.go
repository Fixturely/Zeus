package subscriptions

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestValidateCreateSubscriptionRequest(t *testing.T) {
	tests := []struct {
		name          string
		requestBody   map[string]interface{}
		wantErr       bool
		expectedError string
	}{
		{
			name: "valid subscription request",
			requestBody: map[string]interface{}{
				"sport_id": 1,
				"team_id":  1,
			},
			wantErr:       false,
			expectedError: "",
		},
		{
			name: "missing sport_id",
			requestBody: map[string]interface{}{
				"team_id": 1,
			},
			wantErr:       true,
			expectedError: "sport_id is required",
		},
		{
			name: "validates true if team_id is missing",
			requestBody: map[string]interface{}{
				"sport_id": 1,
			},
			wantErr:       false,
			expectedError: "",
		},
		{
			name:          "invalid request - empty body",
			requestBody:   map[string]interface{}{},
			wantErr:       true,
			expectedError: "sport_id is required",
		},
		{
			name: "invalid request - zero sport_id",
			requestBody: map[string]interface{}{
				"sport_id": 0,
				"team_id":  1,
			},
			wantErr:       true,
			expectedError: "sport_id is required",
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

			// Call the validation function
			result, err := ValidateCreateSubscriptionRequest(c)

			if test.wantErr {
				// Check that we got an error
				assert.Error(t, err)
				assert.Nil(t, result)
				assert.Contains(t, err.Error(), test.expectedError)
			} else {
				// Check that we got a successful result
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, test.requestBody["sport_id"], result.SportID)
				if teamID, ok := test.requestBody["team_id"]; ok {
					assert.Equal(t, teamID, *result.TeamID)
				}
			}
		})
	}
}
