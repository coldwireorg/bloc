package files

import (
	"bloc/models"
	"bloc/storage"
	"bloc/utils/errs"
	"bloc/utils/tokens"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func Download(c *fiber.Ctx) error {
	fileId := c.Params("id")

	file := models.File{
		Id: fileId,
	}

	token, err := tokens.Parse(c.Cookies("token")) // Parse user's JWT token
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.Status(500).JSON(errs.BadRequest)
	}

	file, err = file.Find()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.JSON(errs.Internal)
	}

	if token.Username != file.Owner {
		log.Err(err).Msg(err.Error())
		return c.JSON(errs.Permission)
	}

	c.Response().Header.Add("Content-Disposition", "attachment; filename=\""+file.Name+"\"")
	c.Response().Header.Add("Content-Type", file.Type)

	stream, err := storage.Driver.Get(fileId)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.JSON(errs.Internal)
	}

	return c.SendStream(stream)
}
