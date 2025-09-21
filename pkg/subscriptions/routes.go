package subscriptions

import (
	"subscritracker/pkg/application"
	"subscritracker/pkg/utils"
)

func RegisterRoutes(app *application.App) {
	// Protected routes (require authentication)
	app.Echo.POST("/v1/subscriptions", CreateSubscriptionHandler, utils.AuthMiddleware)
}
