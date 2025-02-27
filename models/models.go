package models

import "gorm.io/gorm"

type URL struct {
	gorm.Model
	ShortURL string `json:"short_url"`
	LongURL  string `json:"long_url"`
}
