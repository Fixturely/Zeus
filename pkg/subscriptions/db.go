package subscriptions

import (
	"context"
	"subscritracker/pkg/application"
	"subscritracker/pkg/models"

	"github.com/labstack/echo/v4"
)

func InsertSubscription(c echo.Context, app *application.App, subscription *models.Subscription) (*models.Subscription, error) {
	userId := c.Get("user_id").(int)
	subscription.UserID = userId
	if err := app.Database.NewInsert().
		Model(subscription).
		Returning("*").
		Scan(context.Background()); err != nil {
		return nil, err
	}

	return subscription, nil
}
