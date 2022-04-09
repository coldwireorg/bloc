package shares

import (
	"bloc/models"
	errors "bloc/utils/errs"
	"bloc/utils/tokens"

	"github.com/gofiber/fiber/v2"
	"github.com/lithammer/shortuuid/v4"
)

func Add(c *fiber.Ctx) error {
	request := struct {
		IsFile  bool   `json:"is_file"`
		ToShare string `json:"to_share"` // File to share
		ShareTo string `json:"share_to"` // User to share the file to
		Key     string `json:"key"`      // Encryption key of the file for this user
	}{}

	err := c.BodyParser(&request)
	if err != nil {
		return errors.Handle(c, errors.ErrBody, err)
	}

	token, err := tokens.Parse(c.Cookies("token"))
	if err != nil {
		return errors.Handle(c, errors.ErrAuth, err)
	}

	file := models.File{Id: request.ToShare}
	file, err = file.Find()
	if err != nil {
		return errors.Handle(c, errors.ErrDatabaseNotFound, err)
	}

	if file.Owner != token.Username {
		return errors.Handle(c, errors.ErrPermission, err)
	}

	share := models.Share{
		Id:         shortuuid.New(),
		IsFavorite: false,
		Key:        request.Key,
		Owner:      request.ShareTo,
		IsFile:     request.IsFile,
	}

	// Create share
	err = share.Add()
	if err != nil {
		return errors.Handle(c, errors.ErrDatabaseCreate, err)
	}

	if request.IsFile {
		share.File = request.ToShare

		err = share.LinkFile()
		if err != nil {
			return errors.Handle(c, errors.ErrDatabaseCreate, err)
		}

		err = share.SetKey()
		if err != nil {
			return errors.Handle(c, errors.ErrDatabaseCreate, err)
		}
	} else {
		share.Folder = request.ToShare

		err = share.LinkFolder()
		if err != nil {
			return errors.Handle(c, errors.ErrDatabaseCreate, err)
		}
	}

	return errors.Handle(c, errors.Success, fiber.Map{
		"share": share.Id,
	})
}
