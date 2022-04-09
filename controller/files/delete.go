package files

import (
	"bloc/models"
	"bloc/storage"
	errors "bloc/utils/errs"
	"bloc/utils/tokens"

	"github.com/gofiber/fiber/v2"
)

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	file := models.File{
		Id: id,
	}

	token, err := tokens.Parse(c.Cookies("token"))
	if err != nil {
		return errors.Handle(c, errors.ErrAuth, err)
	}

	f, err := file.Find()
	if err != nil {
		return errors.Handle(c, errors.ErrDatabaseNotFound, err)
	}

	if f.Owner != token.Username {
		return errors.Handle(c, errors.ErrPermission)
	}

	err = f.Delete()
	if err != nil {
		return errors.Handle(c, errors.ErrDatabaseRemove, err)
	}

	err = storage.Driver.Delete(f.Id)
	if err != nil {
		return errors.Handle(c, errors.ErrUnknown, err)
	}

	return errors.Handle(c, errors.Success)
}
