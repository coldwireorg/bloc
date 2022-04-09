package files

import (
	"bloc/models"
	errors "bloc/utils/errs"
	"bloc/utils/tokens"

	"github.com/gofiber/fiber/v2"
)

func Favorite(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return errors.Handle(c, errors.ErrRequest)
	}

	request := struct {
		Favorite bool `json:"favorite"`
	}{}

	err := c.BodyParser(&request)
	if err != nil {
		return errors.Handle(c, errors.ErrBody, err)
	}

	token, err := tokens.Parse(c.Cookies("token")) // Parse user's JWT token
	if err != nil {
		return errors.Handle(c, errors.ErrAuth, err)
	}

	file := models.File{
		Id:         id,
		IsFavorite: request.Favorite,
	}

	f, err := file.Find()
	if err != nil {
		return errors.Handle(c, errors.ErrDatabaseNotFound, err)
	}

	if f.Owner != token.Username {
		return errors.Handle(c, errors.ErrPermission)
	}

	err = file.SetFavorite()
	if err != nil {
		return errors.Handle(c, errors.ErrDatabaseCreate, err)
	}

	return errors.Handle(c, errors.Success)
}
