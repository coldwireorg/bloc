package middlewares

import (
	"bloc/utils"
	"bloc/utils/config"

	"codeberg.org/coldwire/cwauth"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func Allowed(c *fiber.Ctx) error {
	if config.Conf.Oauth.Server != "" {
		redirect, err := cwauth.AuthURL()
		if err != nil {
			log.Err(err).Msg(err.Error())
		}

		if c.Method() == "GET" {
			return c.Redirect(redirect)
		}

		return c.JSON(utils.Reponse{
			Success: false,
			Error:   "Local authentification is disable, please connect to the authorization server",
			Data: fiber.Map{
				"redirect_url": redirect,
			},
		})
	}

	return c.Next()
}
