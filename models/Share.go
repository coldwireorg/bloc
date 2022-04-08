package models

import (
	"bloc/database"
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/rs/zerolog/log"
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
		log.Err(err).Msg(err.Error())
		return err
	}

	return nil
}

// Link the file to the share
func (s Share) LinkFile() error {
	_, err := database.DB.Exec(context.Background(), `UPDATE shares SET f_file = $1 WHERE id = $2`, s.File, s.Id)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	return err
}

// Link the folder to the share
func (s Share) LinkFolder() error {
	_, err := database.DB.Exec(context.Background(), `UPDATE shares SET f_folder = $1 WHERE id = $2`, s.Folder, s.Id)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	return err
}

// set encryption key of the file
func (s Share) SetKey() error {
	_, err := database.DB.Exec(context.Background(), `UPDATE shares SET key = $1 WHERE id = $2`, s.Key, s.Id)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	return err
}

// Set a parent folder
func (s Share) SetParent() error {
	_, err := database.DB.Exec(context.Background(), `UPDATE shares SET f_parent = $1 WHERE id = $2`, s.Parent, s.Id)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	return err
}

func (s Share) Revoke() error {
	_, err := database.DB.Exec(context.Background(), `DELETE FROM shares WHERE id = $1`, s.Id)
	if err != nil {
		log.Err(err).Msg(err.Error())
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
	f_owner AS owner,
	coalesce(f_file, '')   AS file,
	coalesce(f_folder, '') AS folder,
	coalesce(f_parent, '') AS parent
		FROM shares
			WHERE id = $1`, s.Id)

	if err != nil {
		log.Err(err).Msg(err.Error())
		return share, err
	}

	return share, err
}

func (s Share) List() ([]Share, error) {
	var shares []Share

	shrRows, err := database.DB.Query(context.Background(), `SELECT
	id,
	is_favorite,
	key,
	is_file,
	f_owner AS owner,
	coalesce(f_file, '')   AS file,
	coalesce(f_folder, '') AS folder,
	coalesce(f_parent, '') AS parent
		FROM shares
			WHERE f_owner = $1`, s.Owner)

	err = pgxscan.ScanAll(&shares, shrRows)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return shares, err
	}

	return shares, err
}
