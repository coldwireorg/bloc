package files

import (
	"bloc/models"
	errors "bloc/utils/errs"
	"bloc/utils/tokens"

	"github.com/gofiber/fiber/v2"
)

func List(c *fiber.Ctx) error {
	folder := c.Params("folder")
	if folder == "" {
		return errors.Handle(c, errors.ErrRequest)
	}

	token, err := tokens.Parse(c.Cookies("token")) // Parse user's JWT token
	if err != nil {
		return errors.Handle(c, errors.ErrAuth, err)
	}

	fldr := models.Folder{
		Id:    folder,
		Owner: token.Username,
	}

	fldr, err = fldr.Find()
	if err != nil {
		return errors.Handle(c, errors.ErrDatabaseNotFound, err)
	}

	if fldr.Owner != token.Username {
		return errors.Handle(c, errors.ErrPermission, err)
	}

	folders, files, shares, err := fldr.GetChildrens()
	if err != nil {
		return errors.Handle(c, errors.ErrUnknown, err)
	}

	return errors.Handle(c, errors.Success, fiber.Map{
		"files":  files,
		"folder": folders,
		"shares": shares,
	})
}
