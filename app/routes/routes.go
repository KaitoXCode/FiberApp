package routes

import (
	"github.com/KaitoXCode/fiberApp/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Root getter
	app.Get("/", controllers.RootGet)

	// Authentication
	app.Post("/login", controllers.Login)

	// Users full crud
	app.Get("/users", controllers.UsersGet)
	app.Get("/users/:id", controllers.UserGet)
	app.Post("/users", controllers.UserCreate)
	app.Put("/users/:id", controllers.UserUpdate)
	app.Delete("/users/:id", controllers.UserDelete)
}
