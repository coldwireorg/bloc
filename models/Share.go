package models

import (
	"bloc/database"
	"context"
	"log"

	"github.com/georgysavva/scany/pgxscan"
)

type Share struct {
	Id         string
	Name       string
	Size       int
	IsFavorite bool
	Key        string
	File       string
	Parent     string
	Owner      string
}

func (u Share) Add() error {
	_, err := database.DB.Exec(context.Background(), `INSERT INTO files(
		id,
		name,
		size,
		is_favorite,
		key,
		f_file,
		f_owner AS owner,
		f_parent AS parent
	) VALUES($1, $2, $3, $4, $6)`, u.Id, u.Name, u.Size, u.IsFavorite, u.Key, u.File, u.Owner, u.Parent)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (u Share) Revoke() error {
	_, err := database.DB.Exec(context.Background(), `DELETE FROM shares WHERE id = $1`, u.Id)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return err
}

func (u Share) Find() (Share, error) {
	var share Share
	err := pgxscan.Get(context.Background(), database.DB, &share, `SELECT
	id,
	name,
	size,
	is_favorite,
	key,
	f_file,
	f_owner AS owner,
	f_parent AS parent
		FROM files
			WHERE f_parent = $1`, u.Id)

	if err != nil {
		log.Println(err.Error())
		return share, err
	}

	return share, err
}
