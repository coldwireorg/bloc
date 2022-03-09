package users

import (
	"bloc/models"
	"bloc/utils/errs"

	"github.com/alexedwards/argon2id"
	"github.com/gofiber/fiber/v2"
	"github.com/lithammer/shortuuid/v4"
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

	var usr = models.User{
		Name:       request.Username,
		PrivateKey: request.PrivateKey,
		PublicKey:  request.PublicKey,
		Home: models.FolderAccess{
			Id:   shortuuid.New(),
			Path: "/",
			Folder: models.Folder{
				Id:      shortuuid.New(),
				Name:    "home",
				Folders: []models.Folder{},
				Files:   []models.File{},
			},
		},
	}

	err := c.BodyParser(&request)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.JSON(errs.BadRequest)
	}

	if request.Username == "" {
		return c.JSON(errs.BadRequest)
	}

	if request.PrivateKey == "" || request.PublicKey == "" {
		return c.JSON(errs.AuthNoKeypair)
	}

	exist := usr.Exist()
	if exist {
		return c.JSON(errs.AuthNameAlreadyTaken)
	}

	hash, err := argon2id.CreateHash(request.Password, argon2id.DefaultParams)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.JSON(errs.Internal)
	}

	usr.Password = hash

	err = usr.Create()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.JSON(errs.Internal)
	}

	return c.SendStatus(404)
}
