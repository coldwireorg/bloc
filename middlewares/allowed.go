package middlewares

import (
	"bloc/utils"
	"bloc/utils/config"

	"github.com/gofiber/fiber/v2"
)

func Allowed(c *fiber.Ctx) error {
	if config.Conf.Oauth.Server != "" {
		return c.JSON(utils.Reponse{
			Success: false,
			Error:   "Local authentification is disable, please use oauth",
		})
	}

	return c.Next()
}
