package user

import (
	"bloc/models"
	errors "bloc/utils/errs"
	"bloc/utils/tokens"

	"github.com/gofiber/fiber/v2"
)

func Info(c *fiber.Ctx) error {
	userToken := c.Cookies("token")

	token, err := tokens.Parse(userToken)
	if err != nil {
		return errors.Handle(c, errors.ErrAuth, err)
	}

	usr := models.User{
		Username: token.Username,
	}

	user, err := usr.Find()
	if err != nil {
		return errors.Handle(c, errors.ErrDatabaseGeneric)
	}

	return errors.Handle(c, errors.Success, fiber.Map{
		"username":    user.Username,
		"private_key": user.PrivateKey,
		"public_key":  user.PublicKey,
		"root":        user.Root,
	})
}
