package models

import (
	"bloc/database"
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/rs/zerolog/log"
)

type File struct {
	Id         string `db:"id"          json:"id"`
	Name       string `db:"name"        json:"name"`
	Size       int    `db:"size"        json:"size"`
	Type       string `db:"type"        json:"type"`
	IsFavorite bool   `db:"is_favorite" json:"is_favorite"`
	Key        string `db:"key"         json:"key"`
	Parent     string `db:"parent"      json:"parent"`
	Owner      string `db:"owner"       json:"owner"`
}

func (f File) Create() error {
	_, err := database.DB.Exec(context.Background(), `INSERT INTO files(id, name, size, type, is_favorite, key, f_parent, f_owner) VALUES($1, $2, $3, $4, $5, $6, $7, $8)`, f.Id, f.Name, f.Size, f.Type, f.IsFavorite, f.Key, f.Parent, f.Owner)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	return nil
}

func (f File) Delete() error {
	_, err := database.DB.Exec(context.Background(), `DELETE FROM files WHERE id = $1`, f.Id)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	return err
}

func (f File) Find() (File, error) {
	var file File
	err := pgxscan.Get(context.Background(), database.DB, &file, `SELECT
	id,
	name,
	size,
	type,
	is_favorite,
	key,
	f_owner AS owner,
	f_parent AS parent
		FROM files
			WHERE id = $1`, f.Id)

	if err != nil {
		log.Err(err).Msg(err.Error())
		return file, err
	}

	return file, nil
}

func (f File) List() ([]File, error) {
	var (
		files []File
		err   error
	)

	filesRows, err := database.DB.Query(context.Background(), `SELECT
	id,
	name,
	size,
	is_favorite,
	key,
	f_owner AS owner,
	f_parent AS parent
		FROM files
			WHERE f_parent = $1`, f.Parent)

	err = pgxscan.ScanAll(&files, filesRows)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return files, err
	}

	return files, err
}

func (f File) SetOwner(username string) error {
	_, err := database.DB.Exec(context.Background(), `UPDATE folders SET f_owner = $1 WHERE id = $2`, username, f.Id)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	return err
}

func (f File) SetParent(parent string) error {
	_, err := database.DB.Exec(context.Background(), `UPDATE files SET f_parent = $1 WHERE id = $2`, parent, f.Id)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	return err
}

func (f File) SetFavorite() error {
	_, err := database.DB.Exec(context.Background(), `UPDATE files SET is_favorite = $1 WHERE id = $2`, f.IsFavorite, f.Id)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	return err
}
