package models

import "bloc/database"

type Folder struct {
	Id      string `gorm:"not null;unique;primaryKey"`
	Name    string `gorm:"not null;"`
	Parent  string
	Folders []Folder `gorm:"foreignKey:Parent;references:Id;constraint:OnDelete:CASCADE;"`
	Files   []File   `gorm:"foreignKey:Parent;references:Id;constraint:OnDelete:CASCADE;"`
}

func (f Folder) Create() error {
	return database.DB.Create(&f).Error
}

func (f Folder) Delete() error {
	return database.DB.Delete(&f).Error
}

func (f Folder) Find() (User, error) {
	var usr User
	err := database.DB.Model(&f).Find(&usr).Error
	if err != nil {
		return User{}, err
	}

	return usr, nil
}

type FolderAccess struct {
	Id     string `gorm:"not null;unique;primaryKey"`
	Path   string `gorm:"not null;"`
	Folder Folder `gorm:"not null;foreignKey:Id"`
}

func (fa FolderAccess) Create() error {
	return database.DB.Create(&fa).Error
}

func (fa FolderAccess) Delete() error {
	return database.DB.Delete(&fa).Error
}

func (fa FolderAccess) Find() (User, error) {
	var usr User
	err := database.DB.Model(&fa).Find(&usr).Error
	if err != nil {
		return User{}, err
	}

	return usr, nil
}
