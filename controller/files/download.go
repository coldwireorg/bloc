package files

import (
	"bloc/models"
	"bloc/storage"
	errors "bloc/utils/errs"
	"bloc/utils/tokens"

	"github.com/gofiber/fiber/v2"
)

func Download(c *fiber.Ctx) error {
	fileId := c.Params("id")

	file := models.File{
		Id: fileId,
	}

	token, err := tokens.Parse(c.Cookies("token")) // Parse user's JWT token
	if err != nil {
		return errors.Handle(c, errors.ErrAuth, err)
	}

	file, err = file.Find()
	if err != nil {
		return errors.Handle(c, errors.ErrDatabaseNotFound, err)
	}

	if token.Username != file.Owner {
		return errors.Handle(c, errors.ErrPermission)
	}

	c.Response().Header.Add("Content-Disposition", "attachment; filename=\""+file.Name+"\"")
	c.Response().Header.Add("Content-Type", file.Type)

	stream, err := storage.Driver.Get(fileId)
	if err != nil {
		return errors.Handle(c, errors.ErrUnknown, err)
	}

	return c.SendStream(stream)
}
