package database

import (
	"log"
	"world/models"
)

func CreateWorld(world []models.Country) {

	for i := range world {
		DB.Omit("States", "Cities").Create(&world[i])
		log.Println(world[i])
	}

	for i := range world {
		for j := range world[i].States {
			world[i].States[j].CountryID = world[i].ID
			DB.Omit("Cities").Create(&world[i].States[j])
			log.Println(world[i].States[j])
		}
	}

	for i := range world {
		for j := range world[i].States {
			for k := range world[i].States[j].Cities {
				world[i].States[j].Cities[k].StateID = world[i].States[j].ID
				DB.Create(&world[i].States[j].Cities[k])
				log.Println(world[i].States[j].Cities[k])
			}
		}
	}

}
