package folders

import (
	"bloc/models"
	"bloc/storage"
	"bloc/utils/errs"
	"bloc/utils/tokens"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func cascade(f models.Folder) error {
	folders, files, err := f.GetChildrens()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	// find files and delete them
	for _, d := range files {
		err := storage.Driver.Delete(d.Id)
		if err != nil {
			log.Err(err).Msg(err.Error())
			return err
		}

		err = d.Delete()
		if err != nil {
			log.Err(err).Msg(err.Error())
			return err
		}
	}

	// find others subdir and delete them
	for _, v := range folders {
		go cascade(v)
	}

	return nil
}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	folder := models.Folder{
		Id: id,
	}

	token, err := tokens.Parse(c.Cookies("token"))
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.Status(500).JSON(errs.BadRequest)
	}

	f, err := folder.Find()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.JSON(errs.Internal)
	}

	if f.Owner != token.Username {
		return c.JSON(errs.Permission)
	}

	// Delete every sub-files and sub-folders
	err = cascade(f)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.JSON(errs.Internal)
	}

	err = f.Delete()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return c.JSON(errs.Internal)
	}

	return c.JSON(errs.Success)
}
