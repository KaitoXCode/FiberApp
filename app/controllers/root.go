package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func RootGet(c *fiber.Ctx) error {
	return c.SendString("Welcome to the Fiber App!")
}
