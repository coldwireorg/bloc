package users

import (
	"bloc/utils"
	"bloc/utils/config"

	"codeberg.org/coldwire/cwauth"
	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx) error {
	utils.DelCookie(c, "token")

	if config.Conf.Oauth.Server != "" {
		return c.Redirect(cwauth.LogoutURL())
	}

	return c.Redirect("/")
}
