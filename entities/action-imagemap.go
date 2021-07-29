package entities

import "gorm.io/gorm"

type ImagemapAction struct {
	gorm.Model
	Type              string
	LinkUri           string
	X                 int
	Y                 int
	Width             int
	Height            int
	MessageImagemapId uint
}
