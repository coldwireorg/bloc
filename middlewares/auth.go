package middlewares

import (
	"bloc/utils"
	"bloc/utils/config"
	"bloc/utils/tokens"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func IsAuthenticated(c *fiber.Ctx) error {
	// Check if oauth is configured
	if config.Conf.Oauth.Server != "" {
		return checkOauth2(c)
	}

	token := c.Cookies("token")
	t, err := tokens.Verify(token)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.SendStatus(fiber.ErrForbidden.Code)
	} else {
		if string(t.Token) != "" {
			return c.Next()
		} else {
			utils.DelCookie(c, "token")
			return c.Redirect("/")
		}
	}
}
