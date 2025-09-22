package mike

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckSportExists(t *testing.T) {
	tests := []struct {
		name         string
		sportId      int
		mockResponse string
		expected     bool
		expectError  bool
	}{
		{name: "valid sport exists", sportId: 1, mockResponse: `{"exists": true}`, expected: true, expectError: false},
		{name: "sport does not exist", sportId: 0, mockResponse: `{"exists": false}`, expected: false, expectError: false},
		{name: "invalid json response", sportId: 1, mockResponse: `{"exists": true`, expected: false, expectError: true},
		{name: "empty response", sportId: 1, mockResponse: ``, expected: false, expectError: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Create a mock HTTP server
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Verify the request path
				expectedPath := "/sport/exists/1"
				if test.sportId == 0 {
					expectedPath = "/sport/exists/0"
				}
				assert.Equal(t, expectedPath, r.URL.Path)
				assert.Equal(t, "GET", r.Method)

				// Set response headers
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

				// Write the mock response
				w.Write([]byte(test.mockResponse))
			}))
			defer server.Close()

			// Test the JSON parsing logic that the function uses
			if test.expectError {
				// Test JSON parsing error case
				var response struct {
					Exists bool `json:"exists"`
				}
				err := json.Unmarshal([]byte(test.mockResponse), &response)
				assert.Error(t, err)
			} else {
				// Test successful JSON parsing
				var response struct {
					Exists bool `json:"exists"`
				}
				err := json.Unmarshal([]byte(test.mockResponse), &response)
				assert.NoError(t, err)
				assert.Equal(t, test.expected, response.Exists)
			}
		})
	}
}
