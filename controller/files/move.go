package files

import (
	"bloc/models"
	errors "bloc/utils/errs"
	"bloc/utils/tokens"

	"github.com/gofiber/fiber/v2"
)

func Move(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return errors.Handle(c, errors.ErrRequest)
	}

	request := struct {
		Parent string `json:"parent"`
	}{}

	err := c.BodyParser(&request)
	if err != nil {
		return errors.Handle(c, errors.ErrBody, err)
	}

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
		return errors.Handle(c, errors.ErrPermission, err)
	}

	err = f.SetParent(request.Parent)
	if err != nil {
		return errors.Handle(c, errors.ErrDatabaseCreate, err)
	}

	return errors.Handle(c, errors.Success)
}
