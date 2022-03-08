package shares

import "github.com/gofiber/fiber/v2"

func Add(c *fiber.Ctx) error {
	return c.SendStatus(302)
}
