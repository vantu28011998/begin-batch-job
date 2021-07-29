package entities

import "gorm.io/gorm"

type Messages struct {
	gorm.Model
	Title            string
	SubstanceId      uint
	Status           int
	MessageType      int
	CreateByUsername string
	CreateByUserId   uint
	AgentId          uint
}
