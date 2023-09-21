package main

import (
	"github.com/KaitoXCode/fiberApp/app/routes"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	// Create App
	app := fiber.New()

	// Setting up routes
	routes.SetupRoutes(app)

	// define port
	app.Listen(":3000")
}
