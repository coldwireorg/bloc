package folders

import (
	"bloc/models"
	errors "bloc/utils/errs"
	"bloc/utils/tokens"

	"github.com/gofiber/fiber/v2"
	"github.com/lithammer/shortuuid/v4"
	"github.com/rs/zerolog/log"
)

func Create(c *fiber.Ctx) error {
	request := struct {
		Name   string `json:"name"`
		Parent string `json:"parent"`
	}{}

	err := c.BodyParser(&request)
	if err != nil {
		return errors.Handle(c, errors.ErrBody, err)
	}

	parent := models.Folder{
		Id: request.Parent,
	}

	token, err := tokens.Parse(c.Cookies("token"))
	if err != nil {
		log.Err(err).Msg(err.Error())
		return errors.Handle(c, errors.ErrAuth, err)
	}

	p, err := parent.Find()
	if err != nil {
		return errors.Handle(c, errors.ErrDatabaseNotFound, err)
	}

	if p.Owner != token.Username {
		return errors.Handle(c, errors.ErrPermission, err)
	}

	folder := models.Folder{
		Id:     shortuuid.New(),
		Name:   request.Name,
		Parent: request.Parent,
		Owner:  token.Username,
	}

	err = folder.Create()
	if err != nil {
		return errors.Handle(c, errors.ErrDatabaseCreate, err)
	}

	err = folder.SetOwner(folder.Owner)
	if err != nil {
		return errors.Handle(c, errors.ErrDatabaseCreate, err)
	}

	err = folder.SetParent(folder.Parent)
	if err != nil {
		return errors.Handle(c, errors.ErrDatabaseCreate, err)
	}

	return errors.Handle(c, errors.Success, folder)
}
