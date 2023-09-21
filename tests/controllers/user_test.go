package controllers

import (
	"net/http"
	"testing"

	"github.com/KaitoXCode/fiberApp/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func TestUsersGet(t *testing.T) {
	app := fiber.New()

	app.Get("/users", controllers.UsersGet)

	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatalf("Error creating a new /users GET request: %v", err)
	}

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error making GET request to /users: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status 200 for GET /users but got %d", resp.StatusCode)
	}
}

func TestUserGet(t *testing.T) {
	app := fiber.New()

	app.Get("/users/:id", controllers.UserGet)

	req, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		t.Fatalf("Error creating a new /users/1 GET request: %v", err)
	}

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error making GET request to /users/1: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status 200 for GET /users/1 but got %d", resp.StatusCode)
	}
}

func TestUsersCreate(t *testing.T) {
	app := fiber.New()

	app.Post("/users", controllers.UserCreate)

	req, err := http.NewRequest("POST", "/users", nil)
	if err != nil {
		t.Fatalf("Error creating a new /users POST request: %v", err)
	}

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error making POST request to /users: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status 200 for POST /users but got %d", resp.StatusCode)
	}
}

func TestUsersUpdate(t *testing.T) {
	app := fiber.New()

	app.Put("/users/:id", controllers.UserUpdate)

	req, err := http.NewRequest("PUT", "/users/1", nil)
	if err != nil {
		t.Fatalf("Error creating a new /users/1 PUT request: %v", err)
	}

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error making PUT request to /users/1: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status 200 for PUT /users/1 but got %d", resp.StatusCode)
	}
}

func TestUsersDelete(t *testing.T) {
	app := fiber.New()

	app.Delete("/users/:id", controllers.UserDelete)

	req, err := http.NewRequest("DELETE", "/users/1", nil)
	if err != nil {
		t.Fatalf("Error creating a new /users/1 DELETE request: %v", err)
	}

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error making DELETE request to /users/1: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status 200 for DELETE /users/1 but got %d", resp.StatusCode)
	}
}
