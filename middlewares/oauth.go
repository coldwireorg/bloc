package middlewares

import (
	"codeberg.org/coldwire/cwauth"
	"github.com/gofiber/fiber/v2"
)

func oauth(c *fiber.Ctx) error {
	idToken := c.Cookies("token")
	accesToken := c.Cookies("access_token")

	isTokenValide := cwauth.CheckToken(idToken, accesToken)

	if isTokenValide {
		return c.Next()
	} else {
		url, err := cwauth.AuthURL()
		if err != nil {
			return c.SendStatus(fiber.ErrInternalServerError.Code)
		}

		// if token is not valid, redirect to auth the service
		return c.Redirect(url)
	}
}
