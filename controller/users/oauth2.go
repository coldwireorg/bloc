package users

import "github.com/gofiber/fiber/v2"

func Oauth2Callback(c *fiber.Ctx) error {
	return c.SendStatus(302)
}
