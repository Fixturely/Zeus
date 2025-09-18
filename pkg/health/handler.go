package health

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status":  "ok",
		"message": "Zeus API is healthy",
	})
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	e.GET("/health", h.HealthCheck)
}
