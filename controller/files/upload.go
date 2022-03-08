package files

import "github.com/gofiber/fiber/v2"

func Upload(c *fiber.Ctx) error {
	return c.SendStatus(302)
}
