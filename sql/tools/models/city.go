package models

type City struct {
	Id        uint   `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	StateID   uint   `json:"state"`
}
