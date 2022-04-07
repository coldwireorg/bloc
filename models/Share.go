package models

import (
	"bloc/database"
	"context"
	"log"

	"github.com/georgysavva/scany/pgxscan"
)

type Share struct {
	Id         string `db:"id"          json:"id"`
	IsFavorite bool   `db:"is_favorite" json:"is_favorite"`
	Key        string `db:"key"         json:"key"`
	Parent     string `db:"parent"      json:"parent"`
	Owner      string `db:"owner"       json:"owner"`

	IsFile bool   `db:"is_file" json:"is_file"`
	File   string `db:"file"    json:"file"`
	Folder string `db:"folder"  json:"folder"`
}

func (s Share) Add() error {
	_, err := database.DB.Exec(context.Background(), `INSERT INTO shares(id, is_favorite, key, f_owner, is_file) VALUES($1, $2, $3, $4, $5)`, s.Id, s.IsFavorite, s.Key, s.Owner, s.IsFile)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

// Link the file to the share
func (s Share) LinkFile() error {
	_, err := database.DB.Exec(context.Background(), `UPDATE shares SET f_file = $1 WHERE id = $2`, s.File, s.Id)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return err
}

// Link the folder to the share
func (s Share) LinkFolder() error {
	_, err := database.DB.Exec(context.Background(), `UPDATE shares SET f_folder = $1 WHERE id = $2`, s.Folder, s.Id)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return err
}

// set encryption key of the file
func (s Share) SetKey() error {
	_, err := database.DB.Exec(context.Background(), `UPDATE shares SET key = $1 WHERE id = $2`, s.Key, s.Id)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return err
}

func (s Share) Revoke() error {
	_, err := database.DB.Exec(context.Background(), `DELETE FROM shares WHERE id = $1`, s.Id)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return err
}

func (s Share) Find() (Share, error) {
	var share Share
	err := pgxscan.Get(context.Background(), database.DB, &share, `SELECT
	id,
	is_favorite,
	key,
	is_file,
	f_file   AS file,
	f_folder AS folder,
	f_owner  AS owner,
	f_parent AS parent
		FROM shares
			WHERE f_parent = $1`, s.Id)

	if err != nil {
		log.Println(err.Error())
		return share, err
	}

	return share, err
}
