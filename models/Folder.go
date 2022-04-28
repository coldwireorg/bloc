package models

import (
	"bloc/database"
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/rs/zerolog/log"
)

type Folder struct {
	Id     string `db:"id"     json:"id"`
	Name   string `db:"name"   json:"name"`
	Owner  string `db:"owner"  json:"owner"`
	Parent string `db:"parent" json:"parent"`
}

func (f Folder) Create() error {
	_, err := database.DB.Exec(context.Background(), `INSERT INTO folders(id, name) VALUES($1, $2)`, f.Id, f.Name)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	return nil
}

func (f Folder) Delete() error {
	_, err := database.DB.Exec(context.Background(), `DELETE FROM folders WHERE id = $1`, f.Id)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	return err
}

func (f Folder) Find() (Folder, error) {
	var folder Folder
	err := pgxscan.Get(context.Background(), database.DB, &folder, `SELECT
	id,
	name,
	coalesce(f_owner, '') AS owner,
	coalesce(f_parent, '') AS parent
		FROM folders
			WHERE id = $1`, f.Id)

	if err != nil {
		log.Err(err).Msg(err.Error())
		return folder, err
	}

	return folder, nil
}

// Get files and folders that are childrens of this directory
func (f Folder) Childrens() ([]Folder, error) {
	var (
		folders []Folder
		err     error
	)

	fldrRows, err := database.DB.Query(context.Background(), `SELECT
	id,
	name,
	f_owner AS owner,
	f_parent AS parent
		FROM folders
			WHERE f_parent = $1`, f.Id)

	err = pgxscan.ScanAll(&folders, fldrRows)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return folders, err
	}

	return folders, err
}

func (f Folder) SetOwner(username string) error {
	_, err := database.DB.Exec(context.Background(), `UPDATE folders SET f_owner = $1 WHERE id = $2`, username, f.Id)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	return err
}

func (f Folder) SetParent(parent string) error {
	_, err := database.DB.Exec(context.Background(), `UPDATE folders SET f_parent = $1 WHERE id = $2`, parent, f.Id)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	return err
}
