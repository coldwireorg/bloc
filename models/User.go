package models

import "bloc/database"

type User struct {
	Name       string       `gorm:"not null;unique;primaryKey"`
	Password   string       `gorm:"not null"`
	PrivateKey string       `gorm:"not null"`
	PublicKey  string       `gorm:"not null"`
	Quota      int64        `gorm:"not null"`
	Home       FolderAccess `gorm:"not null;foreignKey:Id"`
}

func (u User) Create() error {
	return database.DB.Create(&u).Error
}

func (u User) Delete() error {
	return database.DB.Delete(&u).Error
}

func (u User) Find() (User, error) {
	var usr User
	err := database.DB.Model(&u).Find(&usr).Error
	if err != nil {
		return User{}, err
	}

	return usr, nil
}

func (u User) Exist() bool {
	usr, err := u.Find()
	if err != nil {
		return false
	}

	if usr.Name != "" {
		return true
	}

	return false
}
