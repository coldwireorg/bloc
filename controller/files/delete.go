package files

import "github.com/gofiber/fiber/v2"

func Delete(c *fiber.Ctx) error {
	return c.SendStatus(302)
}
