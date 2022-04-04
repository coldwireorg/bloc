package models

import (
	"bloc/database"
	"context"
	"log"

	"github.com/georgysavva/scany/pgxscan"
)

type Folder struct {
	Id     string `db:"id" json:"id"`
	Name   string `db:"name" json:"name"`
	Owner  string `db:"f_owner" json:"owner"`
	Parent string `db:"f_parent" json:"parent"`
}

func (f Folder) Create() error {
	_, err := database.DB.Exec(context.Background(), `INSERT INTO folders(id, name, f_parent) VALUES($1, $2, $3)`, f.Id, f.Name, f.Parent)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (f Folder) Delete() error {
	_, err := database.DB.Exec(context.Background(), `DELETE FROM folders WHERE id = $1`, f.Id)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return err
}

func (f Folder) Find() (Folder, error) {
	var folder Folder
	err := pgxscan.Get(context.Background(), database.DB, &folder, `SELECT
	id,
	name,
	f_owner AS owner,
	f_parent AS parent,
		FROM folders
			WHERE id = $1`, f.Id)

	if err != nil {
		log.Println(err.Error())
		return folder, err
	}

	return folder, nil
}

func (f Folder) Move(parent string) error {
	_, err := database.DB.Exec(context.Background(), `UPDATE folders SET f_parent = $1 WHERE id = $2`, parent, f.Id)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return err
}

func (f Folder) SetOwner(username string) error {
	_, err := database.DB.Exec(context.Background(), `UPDATE folders SET f_owner = $1 WHERE id = $2`, username, f.Id)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return err
}
