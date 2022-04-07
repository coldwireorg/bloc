package shares

import (
	"bloc/models"
	"bloc/utils"
	"bloc/utils/errs"
	"bloc/utils/tokens"

	"github.com/gofiber/fiber/v2"
	"github.com/lithammer/shortuuid/v4"
	"github.com/rs/zerolog/log"
)

func Add(c *fiber.Ctx) error {
	request := struct {
		IsFile  bool   `json:"is_file"`
		ToShare string `json:"to_share"` // File to share
		ShareTo string `json:"share_to"` // User to share the file to
		Key     string `json:"key"`      // Encryption key of the file for this user
	}{}

	err := c.BodyParser(&request)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.JSON(errs.Internal)
	}

	token, err := tokens.Parse(c.Cookies("token"))
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.Status(500).JSON(errs.BadRequest)
	}

	file := models.File{Id: request.ToShare}
	file, err = file.Find()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.JSON(errs.Internal)
	}

	if file.Owner != token.Username {
		return c.JSON(errs.Permission)
	}

	share := models.Share{
		Id:         shortuuid.New(),
		IsFavorite: false,
		Key:        request.Key,
		Owner:      request.ShareTo,
		IsFile:     request.IsFile,
	}

	// Create share
	err = share.Add()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.JSON(errs.Internal)
	}

	if request.IsFile {
		share.File = request.ToShare

		err = share.LinkFile()
		if err != nil {
			log.Err(err).Msg(err.Error())
			return c.JSON(errs.Internal)
		}

		err = share.SetKey()
		if err != nil {
			log.Err(err).Msg(err.Error())
			return c.JSON(errs.Internal)
		}
	} else {
		share.Folder = request.ToShare

		err = share.LinkFolder()
		if err != nil {
			log.Err(err).Msg(err.Error())
			return c.JSON(errs.Internal)
		}
	}

	return c.JSON(utils.Reponse{
		Success: true,
		Data: fiber.Map{
			"share": share.Id,
		},
	})
}
