package shares

import (
	"bloc/models"
	errors "bloc/utils/errs"
	"bloc/utils/tokens"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type responseFile struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Size  int    `json:"size"`
	Type  string `json:"type"`
	Owner string `json:"owner"`
}

type responseFolder struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

type response struct {
	Id         string      `json:"id"`
	IsFavorite bool        `json:"is_favorite"`
	Key        string      `json:"key"`
	Parent     string      `json:"parent"`
	Owner      string      `json:"owner"`
	IsFile     bool        `json:"is_file"`
	File       interface{} `json:"file"`
	Folder     interface{} `json:"folder"`
}

func List(c *fiber.Ctx) error {
	var res []response
	folder := c.Params("folder")

	token, err := tokens.Parse(c.Cookies("token"))
	if err != nil {
		log.Err(err).Msg(err.Error())
		return errors.Handle(c, errors.ErrAuth, err)
	}

	share := models.Share{
		Owner: token.Username,
	}

	if folder != "" {
		share.Parent = folder
	}

	shares, err := share.List()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return errors.Handle(c, errors.ErrDatabaseNotFound, err)
	}

	for _, sh := range shares {
		r := response{
			Id:         sh.Id,
			IsFavorite: sh.IsFavorite,
			Key:        sh.Key,
			Parent:     sh.Parent,
			Owner:      sh.Owner,
			IsFile:     sh.IsFile,
		}

		if sh.IsFile {
			fi := models.File{Id: sh.File}
			fi, err := fi.Find()
			if err != nil {
				return errors.Handle(c, errors.ErrDatabaseNotFound, err)
			}

			r.File = responseFile{
				Id:    fi.Id,
				Name:  fi.Name,
				Size:  fi.Size,
				Type:  fi.Type,
				Owner: fi.Owner,
			}

			r.Folder = nil
		} else {
			fo := models.Folder{Id: sh.Folder}
			fo, err := fo.Find()
			if err != nil {
				return errors.Handle(c, errors.ErrDatabaseNotFound, err)
			}

			r.Folder = responseFolder{
				Id:    fo.Id,
				Name:  fo.Name,
				Owner: fo.Owner,
			}

			r.File = nil
		}

		res = append(res, r)
	}

	return errors.Handle(c, errors.Success, res)
}
