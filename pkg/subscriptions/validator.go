package subscriptions

import (
	"fmt"
	"strconv"
	"subscritracker/pkg/models"

	"github.com/labstack/echo/v4"
)

func ValidateCreateSubscriptionRequest(c echo.Context) (*models.Subscription, error) {
	var req models.Subscription

	// First, try to bind (handles application/json and basic form names matching struct fields/tags)
	_ = c.Bind(&req)

	// For multipart/form-data, explicitly read expected form keys
	// sport_id is required
	if req.SportID == 0 {
		if v := c.FormValue("sport_id"); v != "" {
			if n, err := strconv.Atoi(v); err == nil {
				req.SportID = n
			}
		}
		if req.SportID == 0 {
			if v := c.FormValue("SportID"); v != "" {
				if n, err := strconv.Atoi(v); err == nil {
					req.SportID = n
				}
			}
		}
	}
	if req.SportID == 0 {
		return nil, fmt.Errorf("sport_id is required")
	}

	// team_id is optional; parse if present
	if v := c.FormValue("team_id"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			req.TeamID = &n
		}
	} else if v := c.FormValue("TeamID"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			req.TeamID = &n
		}
	}

	return &req, nil
}
