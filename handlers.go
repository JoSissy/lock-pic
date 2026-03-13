package main

import (
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// healthHandler is a simple handler to check the health of the service
func healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "OK",
	})
}

func createLockHandler(c echo.Context) error {

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "file required",
		})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "cannot open file",
		})
	}
	defer src.Close()

	availableAtStr := c.FormValue("available_at")

	availableAt, err := time.Parse(time.RFC3339, availableAtStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid available_at format",
		})
	}

	data, err := io.ReadAll(src)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "cannot read file",
		})
	}

	id := uuid.New()
	idStr, err := SaveImage(c.Request().Context(), id, file.Filename, file.Header.Get("Content-Type"), file.Size, data, availableAt.Format(time.RFC3339))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "database error",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Lock created successfully",
		"id":      idStr,
	})
}
