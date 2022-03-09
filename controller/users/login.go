package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func Login(c *fiber.Ctx) error {
	// Structure of the JSON request
	request := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	err := c.BodyParser(&request)
	if err != nil {
		log.Err(err).Msg(err.Error())
	}

	return c.SendStatus(200)
}
