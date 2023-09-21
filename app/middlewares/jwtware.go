package middlewares

import (
	fiber "github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
)

func Protected() fiber.Handler {
	return jwt.New(jwt.Config{
		SigningKey:  []byte("secret_key"),
		TokenLookup: "header:Authorization",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	})
}
