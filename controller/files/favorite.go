package files

import "github.com/gofiber/fiber/v2"

func Favorite(c *fiber.Ctx) error {
	return c.SendStatus(302)
}
