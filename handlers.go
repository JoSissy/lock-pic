package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// healthHandler is a simple handler to check the health of the service
func healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "OK",
	})
}
