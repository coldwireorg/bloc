package users

import "github.com/gofiber/fiber/v2"

func Logout(c *fiber.Ctx) error {
	return c.SendStatus(302)
}
