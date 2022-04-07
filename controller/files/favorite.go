package files

import (
	"bloc/models"
	"bloc/utils/errs"
	"bloc/utils/tokens"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func Favorite(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.JSON(errs.BadRequest)
	}

	request := struct {
		Favorite bool `json:"favorite"`
	}{}

	err := c.BodyParser(&request)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.JSON(errs.Internal)
	}

	token, err := tokens.Parse(c.Cookies("token")) // Parse user's JWT token
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.Status(500).JSON(errs.BadRequest)
	}

	file := models.File{
		Id:         id,
		IsFavorite: request.Favorite,
	}

	f, err := file.Find()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.Status(500).JSON(errs.Internal)
	}

	if f.Owner != token.Username {
		return c.JSON(errs.Permission)
	}

	err = file.SetFavorite()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.Status(500).JSON(errs.Internal)
	}

	return c.JSON(errs.Success)
}
