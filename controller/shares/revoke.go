package shares

import (
	"bloc/models"
	"bloc/utils/errs"
	"bloc/utils/tokens"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func Revoke(c *fiber.Ctx) error {
	id := c.Params("id")

	share := models.Share{
		Id: id,
	}

	token, err := tokens.Parse(c.Cookies("token"))
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.Status(500).JSON(errs.BadRequest)
	}

	s, err := share.Find()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.JSON(errs.Internal)
	}

	if s.Owner != token.Username {
		return c.JSON(errs.Permission)
	}

	err = s.Revoke()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.JSON(errs.Internal)
	}

	return c.JSON(errs.Success)
}
