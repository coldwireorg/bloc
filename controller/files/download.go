package files

import "github.com/gofiber/fiber/v2"

func Download(c *fiber.Ctx) error {
	return c.SendStatus(302)
}
