package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func Register(c *fiber.Ctx) error {
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
	}

	return c.SendStatus(404)
}
