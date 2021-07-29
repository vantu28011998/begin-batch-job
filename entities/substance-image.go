package entities

import "gorm.io/gorm"

type MessageImage struct {
	gorm.Model
	OriginalContentUrl string
	PreviewImageUrl    string
}
