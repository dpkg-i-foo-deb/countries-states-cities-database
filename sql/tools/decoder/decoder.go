package decoder

import (
	"encoding/json"
	"os"
	"world/models"
	"world/util"
)

func DecodeWorld() *[]models.Country {

	var world []models.Country

	reader, err := os.Open("../../countries+states+cities.json")

	util.HandleErrorStop(err)

	decoder := json.NewDecoder(reader)

	err = decoder.Decode(&world)

	util.HandleErrorStop(err)

	return &world

}
