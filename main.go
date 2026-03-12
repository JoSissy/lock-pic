package main

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	ctx := context.Background()

	// Create a new Echo instance
	e := echo.New()

	// Initialize Google Drive service
	err := initGoogleDrive(ctx)
	if err != nil {
		panic(err)
	}

	// Middleware
	e.Use(middleware.Logger())

	// Routes
	e.GET("/health", healthHandler)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
