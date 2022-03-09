package users

import (
	"bloc/config"

	"codeberg.org/coldwire/cwauth"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func Login(c *fiber.Ctx) error {
	if config.Conf.Oauth.Server != "" {
		url, err := cwauth.AuthURL()
		if err != nil {
			log.Err(err).Msg(err.Error())
			c.Redirect("/")
		}

		return c.Redirect(url)
	}

	return c.SendStatus(200)
}
