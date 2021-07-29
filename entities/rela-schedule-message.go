package entities

import "gorm.io/gorm"

type RelaScheduleMessage struct {
	gorm.Model
	ScheduleId  uint
	MessageId   uint
}
