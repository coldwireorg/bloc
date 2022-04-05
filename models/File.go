package models

import (
	"bloc/database"
	"context"
	"log"

	"github.com/georgysavva/scany/pgxscan"
)

type File struct {
	Id         string
	Name       string
	Size       int
	IsFavorite bool
	Key        string
	Parent     string
	Owner      string
}

func (f File) Create() error {
	_, err := database.DB.Exec(context.Background(), `INSERT INTO files(id, name, size, is_favorite, key, f_parent, f_owner) VALUES($1, $2, $3, $4, $5, $6, $7)`, f.Id, f.Name, f.Size, f.IsFavorite, f.Key, f.Parent, f.Owner)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (f File) Delete() error {
	_, err := database.DB.Exec(context.Background(), `DELETE FROM files WHERE id = $1`, f.Id)
	if err != nil {
		log.Println(err.Error())
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
	is_favorite,
	key,
	f_owner AS owner
		FROM files
			WHERE id = $1`, f.Id)

	if err != nil {
		log.Println(err.Error())
		return file, err
	}

	return file, nil
}

func (f File) SetOwner(username string) error {
	_, err := database.DB.Exec(context.Background(), `UPDATE folders SET f_owner = $1 WHERE id = $2`, username, f.Id)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return err
}

func (f File) SetParent(parent string) error {
	_, err := database.DB.Exec(context.Background(), `UPDATE files SET f_parent = $1 WHERE id = $2`, parent, f.Id)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return err
}
