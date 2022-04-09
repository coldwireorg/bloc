package users

import (
	"bloc/models"
	"bloc/utils"
	errors "bloc/utils/errs"
	"bloc/utils/tokens"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/gofiber/fiber/v2"
	"github.com/lithammer/shortuuid/v4"
)

func Register(c *fiber.Ctx) error {
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

	var root = models.Folder{
		Id:   shortuuid.New(),
		Name: "root",
	}

	var usr = models.User{
		Username:   request.Username,
		PrivateKey: request.PrivateKey,
		PublicKey:  request.PublicKey,
		AuthMode:   "LOCAL",
	}

	if request.Username == "" {
		return errors.Handle(c, errors.ErrBody, err)
	}

	if request.PrivateKey == "" || request.PublicKey == "" {
		return errors.Handle(c, errors.ErrAuth)
	}

	exist, err := usr.Exist()
	if err != nil {
		return errors.Handle(c, errors.ErrAuth, err)
	}

	if exist {
		return errors.Handle(c, errors.ErrAuthExist)
	}

	hash, err := argon2id.CreateHash(request.Password, argon2id.DefaultParams)
	if err != nil {
		return errors.Handle(c, errors.ErrAuth, err)
	}

	usr.Password = hash

	// Create root folder
	root.Create()
	usr.Root = root.Id // add root folder to the user for the response

	err = usr.Create()
	if err != nil {
		return errors.Handle(c, errors.ErrDatabaseCreate, err)
	}

	err = usr.SetRoot(root.Id) // Set root folder of the user
	if err != nil {
		return errors.Handle(c, errors.ErrDatabaseCreate, err)
	}

	err = root.SetOwner(usr.Username) // Set the owner of the root folder
	if err != nil {
		return errors.Handle(c, errors.ErrDatabaseCreate, err)
	}

	token := tokens.Generate(tokens.Token{
		Username:   request.Username,
		PrivateKey: request.PrivateKey,
		PublicKey:  request.PublicKey,
	}, 12*time.Hour)

	utils.SetCookie(c, "token", token, time.Now().Add(time.Hour*6))

	return errors.Handle(c, errors.Success, fiber.Map{
		"token": token,
		"user":  usr,
	})
}
