package folders

import (
	"bloc/models"
	"bloc/utils"
	"bloc/utils/errs"
	"bloc/utils/tokens"

	"github.com/gofiber/fiber/v2"
	"github.com/lithammer/shortuuid/v4"
	"github.com/rs/zerolog/log"
)

func Create(c *fiber.Ctx) error {
	request := struct {
		Name   string `json:"name"`
		Parent string `json:"parent"`
		Owner  string `json:"owner"`
	}{}

	parent := models.Folder{
		Id: request.Parent,
	}

	token, err := tokens.Parse(c.Cookies("token"))
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.Status(500).JSON(errs.BadRequest)
	}

	p, err := parent.Find()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.JSON(errs.Internal)
	}

	if p.Owner != token.Username {
		return c.JSON(errs.Permission)
	}

	folder := models.Folder{
		Id:     shortuuid.New(),
		Name:   request.Name,
		Parent: request.Parent,
		Owner:  request.Owner,
	}

	err = folder.Create()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.JSON(errs.Internal)
	}

	return c.JSON(utils.Reponse{
		Success: true,
		Data:    folder,
	})
}
