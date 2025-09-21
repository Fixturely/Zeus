package subscriptions

import (
	"zeus/pkg/application"
	"zeus/pkg/utils"
)

func RegisterRoutes(app *application.App) {
	// Protected routes (require authentication)
	app.Echo.POST("/v1/subscriptions", CreateSubscriptionHandler, utils.AuthMiddleware)
}
