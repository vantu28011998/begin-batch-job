package entities

import "gorm.io/gorm"

type LineEmojiProduct struct {
	gorm.Model
	LineProductId string
}
