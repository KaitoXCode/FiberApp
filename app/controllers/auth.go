package controllers

import (
	services "github.com/KaitoXCode/fiberApp/app/services"
	"go.uber.org/zap"

	fiber "github.com/gofiber/fiber/v2"
)

// Dummy users for demo
var users = map[string]string{
	"user1": "pass1",
	"user2": "pass2",
}

func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad request",
		})
	}

	password, userExists := users[input.Username]

	if !userExists || password != input.Password {
		services.Logger.Info(
			"Unauthorized access attempt",
			zap.String("username", input.Username),
		)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	// Dummpy userID
	token, err := services.GenerateToken(1)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}
