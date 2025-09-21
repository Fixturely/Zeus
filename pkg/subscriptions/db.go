package subscriptions

import (
	"fmt"
	"zeus/pkg/application"
	"zeus/pkg/models"

	"github.com/labstack/echo/v4"
)

func InsertSubscription(c echo.Context, app *application.App, subscription *models.Subscription) (*models.Subscription, error) {
	userId, ok := c.Get("user_id").(int)
	if !ok {
		return nil, fmt.Errorf("user_id not found in context")
	}
	subscription.UserID = userId

	// Validate required fields
	if subscription.UserID == 0 {
		return nil, fmt.Errorf("user_id is required")
	}
	if subscription.SportID == 0 {
		return nil, fmt.Errorf("sport_id is required")
	}

	if err := app.Database.NewInsert().
		Model(subscription).
		Returning("*").
		Scan(c.Request().Context()); err != nil {
		return nil, err
	}

	return subscription, nil
}
