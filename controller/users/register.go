package users

import "github.com/gofiber/fiber/v2"

func Register(c *fiber.Ctx) error {
	return c.SendStatus(302)
}
