package models

import (
	"bloc/database"
	"context"
	"log"

	"github.com/georgysavva/scany/pgxscan"
)

type User struct {
	Username   string `db:"username"      json:"username"`
	Password   string `db:"password"      json:"-"`
	AuthMode   string `db:"auth_mode"     json:"authMode"`
	PrivateKey string `db:"private_key"   json:"privateKey"`
	PublicKey  string `db:"public_key"    json:"publicKey"`
	Root       string `db:"root"          json:"root"`
}

func (u User) Create() error {
	_, err := database.DB.Exec(context.Background(), `INSERT INTO users(username, password, auth_mode, public_key, private_key) VALUES($1, $2, $3, $4, $5)`, u.Username, u.Password, u.AuthMode, u.PublicKey, u.PrivateKey)
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

	return user, err
}

func (u User) Exist() (bool, error) {
	var user User
	err := pgxscan.Get(context.Background(), database.DB, &user, `SELECT
	username
		FROM users
			WHERE username = $1`, u.Username)

	if err != nil {
		return false, err
	}

	if user.Username == "" {
		return false, err
	}

	return true, err
}

func (u User) SetRoot(id string) error {
	_, err := database.DB.Exec(context.Background(), `UPDATE users SET f_root_folder = $1 WHERE username = $2`, id, u.Username)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return err
}
