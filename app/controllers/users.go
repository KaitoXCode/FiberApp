package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func UsersGet(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "All Users"})
}

func UserGet(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(fiber.Map{"message": "User with ID " + id})
}

func UserCreate(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "User Created"})
}

func UserUpdate(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(fiber.Map{"message": "User with ID " + id + " updated"})
}

func UserDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(fiber.Map{"message": "User with ID " + id + " deleted"})
}
