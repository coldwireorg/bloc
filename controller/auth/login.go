package auth

import (
	"bloc/models"
	"bloc/utils"
	errors "bloc/utils/errs"
	"bloc/utils/tokens"
	"regexp"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/gofiber/fiber/v2"
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
		return errors.Handle(c, errors.ErrBody, err)
	}

	var user = models.User{
		Username: request.Username,
	}

	usernameValidation, _ := regexp.MatchString("[a-zA-Z]{3,}", request.Username)
	if !usernameValidation {
		return errors.Handle(c, errors.ErrBody, "invalid username")
	}

	user, err = user.Find()
	if err != nil {
		return errors.Handle(c, errors.ErrAuth, err)
	}

	isValid, err := argon2id.ComparePasswordAndHash(request.Password, user.Password)
	if !isValid {
		return errors.Handle(c, errors.ErrAuthPassword)
	}

	if err != nil {
		return errors.Handle(c, errors.ErrAuth, err)
	}

	token := tokens.Generate(tokens.Token{
		Username: request.Username,
	}, 12*time.Hour)

	utils.SetCookie(c, "token", token, time.Now().Add(time.Hour*6))

	return errors.Handle(c, errors.Success, fiber.Map{
		"token": token,
		"user":  user,
	})
}
