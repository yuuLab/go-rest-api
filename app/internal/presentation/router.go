package presentation

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Router sets up the application routes.
func Router(e *echo.Echo) {
	// Health check endpoint
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, world!"})
	})
}
