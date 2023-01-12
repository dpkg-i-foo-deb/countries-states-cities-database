package models

import "gorm.io/gorm"

type TimeZone struct {
	gorm.Model
	CountryID     uint
	ZoneName      string `json:"zoneName"`
	GtmOffset     uint   `json:"gtmOffset"`
	GtmOffsetName string `json:"gtmOffsetName"`
	Abbreviation  string `json:"abbreviation"`
	TzName        string `json:"tzName"`
}
