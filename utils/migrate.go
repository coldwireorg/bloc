package utils

import (
	"bloc/database"
	"bloc/models"
	"time"

	"github.com/rs/zerolog/log"
)

// Migrate database tables
func Migrate() bool {
	for {
		if database.Ready {
			database.DB.Migrator().DropTable(&models.User{}, &models.File{}, &models.FileAccess{}, &models.FileKey{}, &models.Folder{}, &models.FolderAccess{})
			err := database.DB.AutoMigrate(&models.User{}, &models.File{}, &models.FileAccess{}, &models.FileKey{}, &models.Folder{}, &models.FolderAccess{})
			if err != nil {
				log.Err(err).Msg(err.Error())
				time.Sleep(5 * time.Second)
			} else {
				return true
			}
		}
	}
}
