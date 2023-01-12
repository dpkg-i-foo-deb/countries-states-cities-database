package database

import (
	"world/util"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error

	//DB, err = gorm.Open(sqlite.Open("wold.db"), &gorm.Config{})

	DB, err = gorm.Open(postgres.Open("host=localhost user=world password=world dbname=world port=5434 sslmode=disable TimeZone=America/Bogota"))

	util.HandleErrorStop(err)

	MigrateModels()
}

func connect() {
	var err error

	DB, err = gorm.Open(postgres.Open("host=localhost user=world password=world dbname=world port=5434 sslmode=disable TimeZone=America/Bogota"))

	util.HandleErrorStop(err)
}
