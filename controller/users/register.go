package users

import (
	"bloc/models"
	"bloc/utils"
	"bloc/utils/errs"
	"bloc/utils/tokens"
	"time"

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
		AuthMode   string `json:"auth_mode"`
	}{}

	err := c.BodyParser(&request)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.JSON(errs.BadRequest)
	}

	log.Print(request)

	var root = models.Folder{
		Id:   shortuuid.New(),
		Name: "root",
	}

	var usr = models.User{
		Username:   request.Username,
		PrivateKey: request.PrivateKey,
		PublicKey:  request.PublicKey,
		AuthMode:   request.AuthMode,
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

	// Create root folder
	root.Create()

	err = usr.Create()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.JSON(errs.Internal)
	}

	usr.SetRoot(root.Id)        // Set root folder of the user
	root.SetOwner(usr.Username) // Set the owner of the root folder

	token := tokens.Generate(tokens.Token{
		Username:   request.Username,
		PrivateKey: request.PrivateKey,
		PublicKey:  request.PublicKey,
	}, 12*time.Hour)

	utils.SetCookie(c, "token", token, time.Now().Add(time.Hour*6))

	return c.JSON(utils.Reponse{
		Success: true,
		Data: fiber.Map{
			"token": token,
			"user":  usr,
		},
	})
}
