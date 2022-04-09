package folders

import (
	"bloc/models"
	"bloc/storage"
	errors "bloc/utils/errs"
	"bloc/utils/tokens"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// Files and folders to delete
var Queue []string

func toDelete(id string) error {
	f := models.Folder{Id: id}
	folders, files, shares, err := f.GetChildrens()
	if err != nil {
		return err
	}

	for _, fi := range files {
		if len(files) == 0 {
			break
		}

		Queue = append(Queue, "fi:"+fi.Id)
	}

	for _, sh := range shares {
		if len(shares) == 0 {
			break
		}

		Queue = append(Queue, "sh:"+sh.Id)
	}

	for _, fo := range folders {
		if len(folders) == 0 {
			break
		}

		Queue = append(Queue, "fo:"+fo.Id)
		err := toDelete(fo.Id)
		if err != nil {
			return err
		}
	}

	return nil
}

func cascade(f models.Folder) error {
	err := toDelete(f.Id)
	if err != nil {
		return err
	}

	for _, q := range Queue {
		t := strings.Split(q, ":")

		// if it's a file
		if t[0] == "fi" {
			// Since the file will be automatically deleted by SQL, we just need to delete it from the storage
			err := storage.Driver.Delete(t[1])
			if err != nil {
				log.Err(err).Msg(err.Error())
				return err
			}
		} else if t[0] == "fo" {
			fo := models.Folder{Id: t[1]}
			err := fo.Delete()
			if err != nil {
				log.Err(err).Msg(err.Error())
				return err
			}
		} else if t[0] == "sh" {
			sh := models.Share{Id: t[1]}
			err := sh.Revoke()
			if err != nil {
				log.Err(err).Msg(err.Error())
				return err
			}
		}
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
		return errors.Handle(c, errors.ErrAuth, err)
	}

	f, err := folder.Find()
	if err != nil {
		return errors.Handle(c, errors.ErrDatabaseNotFound, err)
	}

	if f.Owner != token.Username {
		return errors.Handle(c, errors.ErrPermission, err)
	}

	// Delete every sub-files and sub-folders
	err = cascade(f)
	if err != nil {
		return errors.Handle(c, errors.ErrUnknown, err)
	}

	// When everything is deleted: delete the folder
	err = f.Delete()
	if err != nil {
		return errors.Handle(c, errors.ErrDatabaseRemove, err)
	}

	return errors.Handle(c, errors.Success)
}
