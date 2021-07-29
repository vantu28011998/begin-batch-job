package entities

import "gorm.io/gorm"

type LineEmoji struct {
	gorm.Model
	Index         int
	LineProductId string
	EmojiId       string
}
