package models

type State struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	Name      string  `json:"name"`
	StateCode string  `json:"state_code"`
	Latitude  string  `json:"latitude"`
	Longitude string  `json:"longitude"`
	Type      *string `json:"type"`
	CountryID uint    `json:"-"`
	Cities    []City  `json:"cities"`
}
