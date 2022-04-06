package folders

import (
	"bloc/models"
	"bloc/utils/errs"
	"bloc/utils/tokens"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func Move(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(500).JSON(errs.BadRequest)
	}

	request := struct {
		Parent string `json:"parent"`
	}{}

	err := c.BodyParser(&request)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.JSON(errs.Internal)
	}

	folder := models.Folder{
		Id: id,
	}

	token, err := tokens.Parse(c.Cookies("token"))
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.Status(500).JSON(errs.BadRequest)
	}

	f, err := folder.Find()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.JSON(errs.Internal)
	}

	if f.Owner != token.Username {
		return c.JSON(errs.Permission)
	}

	err = f.SetParent(request.Parent)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.JSON(errs.Internal)
	}

	return c.JSON(errs.Success)
}
