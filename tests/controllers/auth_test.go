package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/KaitoXCode/fiberApp/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func TestLogin(t *testing.T) {
	app := fiber.New()

	app.Post("/login", controllers.Login)

	payload := map[string]string{
		"username": "user1",
		"password": "pass1",
	}

	requestBody, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", "/login", bytes.NewReader(requestBody))
	if err != nil {
		t.Fatalf("Error creating a new /login request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error making POST request to /login: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status 200 for /login but got %d", resp.StatusCode)
	}
}
