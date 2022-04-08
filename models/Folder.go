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
func (f Folder) GetChildrens() ([]Folder, []File, []Share, error) {
	var (
		folders []Folder
		files   []File
		shares  []Share
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
		return folders, files, shares, err
	}

	filesRows, err := database.DB.Query(context.Background(), `SELECT
	id,
	name,
	size,
	is_favorite,
	key,
	f_owner AS owner,
	f_parent AS parent
		FROM files
			WHERE f_parent = $1`, f.Id)

	err = pgxscan.ScanAll(&files, filesRows)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return folders, files, shares, err
	}

	shrRows, err := database.DB.Query(context.Background(), `SELECT
	id,
	is_favorite,
	key,
	is_file,
	f_file   AS file,
	f_folder AS folder,
	f_owner  AS owner,
	f_parent AS parent
		FROM shares
			WHERE f_parent = $1`, f.Parent)

	err = pgxscan.ScanAll(&shares, shrRows)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return folders, files, shares, err
	}

	return folders, files, shares, err
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
