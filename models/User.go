package models

import (
	"bloc/database"
	"context"
	"log"

	"github.com/georgysavva/scany/pgxscan"
)

type User struct {
	Username   string
	Password   string
	AuthMode   string
	PrivateKey string
	PublicKey  string
	Root       string
}

func (u User) Create() error {
	_, err := database.DB.Exec(context.Background(), `INSERT INTO users(username, password, auth_mode, f_root_folder, public_key, private_key) VALUES($1, $2, $3, $4, $5, $6)`, u.Username, u.Password, u.AuthMode, u.Root, u.PublicKey, u.PrivateKey)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (u User) Delete() error {
	_, err := database.DB.Exec(context.Background(), `DELETE FROM users WHERE username = $1`, u.Username)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return err
}

func (u User) Find() (User, error) {
	var user User
	err := pgxscan.Get(context.Background(), database.DB, &user, `SELECT
	username,
	password,
	auth_mode,
	f_root_folder AS root,
	private_key,
	public_key
		FROM users
			WHERE username = $1`, u.Username)

	if err != nil {
		log.Println(err.Error())
		return user, err
	}

	return user, nil
}

func (u User) Exist() bool
