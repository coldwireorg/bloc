package files

import (
	"bloc/models"
	"bloc/storage"
	"bloc/utils/errs"
	"bloc/utils/tokens"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	file := models.File{
		Id: id,
	}

	token, err := tokens.Parse(c.Cookies("token"))
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.Status(500).JSON(errs.BadRequest)
	}

	f, err := file.Find()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.JSON(errs.Internal)
	}

	if f.Owner != token.Username {
		return c.JSON(errs.Permission)
	}

	err = f.Delete()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.JSON(errs.Internal)
	}

	err = storage.Driver.Delete(f.Id)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	return c.SendStatus(302)
}
