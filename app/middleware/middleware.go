package middleware

import "github.com/gofiber/fiber/v2"

func Fake(c *fiber.Ctx) error {
	return c.Next()
}
