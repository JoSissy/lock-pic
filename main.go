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

	// Load configuration
	err := loadConfig(ctx)
	if err != nil {
		e.Logger.Fatal("Failed to load configuration: ", err)
	}

	// Connect to the database
	err = PostgresConnect(ctx)
	if err != nil {
		e.Logger.Fatal("Failed to connect to database: ", err)
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/health", healthHandler)
	e.POST("/locks", createLockHandler)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
