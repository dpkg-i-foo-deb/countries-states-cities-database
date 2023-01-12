package models

type Country struct {
	ID             uint        `gorm:"primaryKey" json:"id"`
	Name           string      `json:"name"`
	Iso3           string      `json:"iso3"`
	Iso2           string      `json:"iso2"`
	NumericCode    string      `json:"numeric_code"`
	PhoneCode      string      `json:"phone_code"`
	Capital        string      `json:"capital"`
	Currency       string      `json:"currency"`
	CurrencyName   string      `json:"currency_name"`
	CurrencySymbol string      `json:"currency_symbol"`
	Tld            string      `json:"tld"`
	Native         string      `json:"native"`
	Region         string      `json:"region"`
	Subregion      string      `json:"subregion"`
	Latitude       string      `json:"latitude"`
	Longitude      string      `json:"longitude"`
	Emoji          string      `json:"emoji"`
	EmojiU         string      `json:"emojiU"`
	States         []State     `json:"states"`
	TimeZones      []TimeZone  `json:"timezones"`
	Translations   Translation `json:"translations"`
}
