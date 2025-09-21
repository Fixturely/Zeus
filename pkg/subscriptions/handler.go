package subscriptions

import (
	"net/http"
	"subscritracker/pkg/application"

	"github.com/labstack/echo/v4"
)

func CreateSubscriptionHandler(c echo.Context) error {
	// Validate the request body
	req, err := ValidateCreateSubscriptionRequest(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Insert the subscription
	inserted, err := InsertSubscription(c, c.Get("app").(*application.App), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, inserted)

}
