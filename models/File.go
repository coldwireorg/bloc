package models

import (
	"bloc/database"
	"errors"

	"github.com/lithammer/shortuuid/v4"
)

/* File model */

type File struct {
	Id   string    `gorm:"not null;unique;primaryKey"`
	Name string    `gorm:"not null;"`
	Type string    `gorm:"not null;"`
	Size int64     `gorm:"not null;"`
	Keys []FileKey `gorm:"foreignKey:FileId;references:Id;constraint:OnDelete:CASCADE;"`
}

func (f File) Create() error {
	return database.DB.Create(&f).Error
}

func (f File) Delete() error {
	return database.DB.Delete(&f).Error
}

func (f File) Update() error {
	return database.DB.Model(&f).Updates(&f).Error
}

func (f File) Find() (User, error) {
	var usr User
	err := database.DB.Model(&f).Find(&usr).Error
	if err != nil {
		return User{}, err
	}

	return usr, nil
}

/* File key model (functions are still linked to the file model) */

type FileKey struct {
	Id     string `gorm:"not null;unique;primaryKey"`
	FileId string `gorm:"not null;"`
	Key    string `gorm:"not null;"`
	User   User   `gorm:"not null;foreignKey:Name"`
}

func (f File) GetKey(usr User) (string, error) {
	var key FileKey
	err := database.DB.Model(&f).Association("Keys").Find(&key, FileKey{
		User: usr,
	}).Error()
	if err != "" {
		return "", errors.New(err)
	}

	if key.Key == "" {
		return "", errors.New("no key")
	}

	return key.Key, nil
}

func (f File) AddKey(usr User, key string) error {
	var k = FileKey{
		Id:   shortuuid.New(),
		User: usr,
		Key:  key,
	}

	err := database.DB.Model(&f).Association("Keys").Append(&k).Error()
	if err != "" {
		return errors.New(err)
	}

	return nil
}

func (f File) DelKey(usr User) error {
	err := database.DB.Model(&f).Association("Keys").Delete(&FileKey{
		User: usr,
	}).Error()
	if err != "" {
		return errors.New(err)
	}

	return nil
}

type FileAccess struct {
	Id       string `gorm:"not null;unique;primaryKey"`
	Favorite bool   `gorm:"not null;"`
	File     File   `gorm:"not null;foreignKey:Id"`
}

func (fa FileAccess) Create() error {
	return database.DB.Create(&fa).Error
}

func (fa FileAccess) Delete() error {
	return database.DB.Delete(&fa).Error
}

func (fa FileAccess) Update() error {
	return database.DB.Model(&fa).Updates(&fa).Error
}
