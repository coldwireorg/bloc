package shares

import (
	"bloc/models"
	errors "bloc/utils/errs"
	"bloc/utils/tokens"

	"github.com/gofiber/fiber/v2"
)

func Revoke(c *fiber.Ctx) error {
	id := c.Params("id")

	share := models.Share{
		Id: id,
	}

	token, err := tokens.Parse(c.Cookies("token"))
	if err != nil {
		return errors.Handle(c, errors.ErrAuth, err)

	}

	s, err := share.Find()
	if err != nil {
		return errors.Handle(c, errors.ErrDatabaseNotFound, err)
	}

	if s.Owner != token.Username {
		return errors.Handle(c, errors.ErrPermission)
	}

	err = s.Revoke()
	if err != nil {
		return errors.Handle(c, errors.ErrDatabaseRemove, err)
	}

	return errors.Handle(c, errors.Success)
}
