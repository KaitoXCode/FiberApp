package controllers

import (
	"net/http"
	"testing"

	"github.com/KaitoXCode/fiberApp/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func TestRootGet(t *testing.T) {
	app := fiber.New()

	app.Get("/", controllers.RootGet)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Error creating a new /root request: %v", err)
	}

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error making GET request to /root: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status 200 for /root but got %d", resp.StatusCode)
	}
}
