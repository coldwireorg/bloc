package files

import (
	"bloc/models"
	"bloc/utils"
	"bloc/utils/errs"
	"bloc/utils/tokens"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func List(c *fiber.Ctx) error {
	folder := c.Params("folder")
	if folder == "" {
		return c.JSON(errs.BadRequest)
	}

	token, err := tokens.Parse(c.Cookies("token")) // Parse user's JWT token
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.Status(500).JSON(errs.BadRequest)
	}

	fldr := models.Folder{
		Id:    folder,
		Owner: token.Username,
	}

	folders, files, err := fldr.GetChildrens()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.JSON(errs.Internal)
	}

	return c.JSON(utils.Reponse{
		Success: true,
		Data: fiber.Map{
			"files":  files,
			"folder": folders,
		},
	})
}
