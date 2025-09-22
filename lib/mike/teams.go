package mike

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"zeus/config"
)

func CheckTeamExists(teamId int) (bool, error) {

	cfg := config.GetConfig()
	// Make a GET request to MIKE API to check if the team exisits
	url := fmt.Sprintf("%s/team/exists/%d", cfg.MikeAPIKey, teamId)

	// Make a GET request to the Mike API to check if the team exists
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close() // Always close the response body

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse the JSON response
	var response struct {
		Exists bool `json:"exists"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return false, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	return response.Exists, nil
}
