package main

import (
	"world/database"
	"world/decoder"
)

func main() {

	database.MigrateModels()

	world := decoder.DecodeWorld()

	database.CreateWorld(*world)
}
