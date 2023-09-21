package main

import (
	"github.com/KaitoXCode/fiberApp/app/routes"
	"github.com/KaitoXCode/fiberApp/app/services"
	"github.com/KaitoXCode/fiberApp/app/utils"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Create App
	app := fiber.New()

	// Init logger
	if err := services.InitLogger(); err != nil {
		panic("Failed to initialize logger")
	}

	// Add middleware to log all requests/responses
	zapWriter := &utils.ZapWriter{}
	app.Use(logger.New(logger.Config{
		Output: zapWriter, // Use the zap logger
	}))

	// Setting up routes
	routes.SetupRoutes(app)

	// define port
	app.Listen(":3000")
}
