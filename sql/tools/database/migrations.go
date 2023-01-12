package database

import (
	"world/models"
	"world/util"
)

func MigrateModels() {
	err := DB.AutoMigrate(
		models.Country{},
		models.State{},
		models.City{},
		models.TimeZone{},
		models.Translation{},
	)
	util.HandleErrorStop(err)
}
