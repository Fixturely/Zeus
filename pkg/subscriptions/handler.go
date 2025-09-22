package subscriptions

import (
	"net/http"
	"zeus/lib/mike"
	"zeus/pkg/application"

	"github.com/labstack/echo/v4"
)

func CreateSubscriptionHandler(c echo.Context) error {
	// Validate the request body
	req, err := ValidateCreateSubscriptionRequest(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Get the sport_id and team_id from the request
	sportId := req.SportID
	teamId := req.TeamID

	//Make sure sportId exists in the database
	exists, err := mike.CheckSportExists(sportId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if !exists {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Sport does not exist"})
	}

	//Make sure teamId exists in the database if teamid provided
	if teamId != nil {
		exists, err := mike.CheckTeamExists(*teamId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		if !exists {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Team does not exist"})
		}
	}

	// Insert the subscription
	inserted, err := InsertSubscription(c, c.Get("app").(*application.App), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, inserted)

}
