package shares

import "github.com/gofiber/fiber/v2"

func Revoke(c *fiber.Ctx) error {
	return c.SendStatus(302)
}
