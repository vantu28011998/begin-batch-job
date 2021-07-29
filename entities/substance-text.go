package entities

import (
	"gorm.io/gorm"
)

type MessagesText struct {
	gorm.Model
	Content string
}
