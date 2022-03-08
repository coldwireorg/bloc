package files

import "github.com/gofiber/fiber/v2"

func List(c *fiber.Ctx) error {
	return c.SendStatus(302)
}
