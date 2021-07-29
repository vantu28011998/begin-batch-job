package entities

import "gorm.io/gorm"

type MessageImagemap struct {
	gorm.Model
	BaseUrl string
	AltText string
	Width int
	Height int
}
