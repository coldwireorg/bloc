package users

import (
	"bloc/models"
	"bloc/utils"
	"bloc/utils/errs"
	"bloc/utils/tokens"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func Login(c *fiber.Ctx) error {
	// Structure of the JSON request
	request := struct {
		Username   string `json:"username"`
		Password   string `json:"password"`
		PrivateKey string `json:"private_key"`
		PublicKey  string `json:"public_key"`
	}{}

	err := c.BodyParser(&request)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.JSON(errs.Internal)
	}

	var user = models.User{
		Username: request.Username,
	}

	user, err = user.Find()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.JSON(errs.Internal)
	}

	isValid, err := argon2id.ComparePasswordAndHash(request.Password, user.Password)
	if !isValid {
		return c.JSON(errs.AuthBadPassword)
	}

	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.JSON(errs.Internal)
	}

	token := tokens.Generate(tokens.Token{
		Username:   request.Username,
		PrivateKey: request.PrivateKey,
		PublicKey:  request.PublicKey,
	}, 12*time.Hour)

	utils.SetCookie(c, "token", token, time.Now().Add(time.Hour*6))

	return c.JSON(errs.Success)
}
