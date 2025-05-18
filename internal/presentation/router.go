package presentation

import (
	"net/http"

	"github.com/labstack/echo"
)

func Router(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!")
	})
}
