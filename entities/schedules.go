package entities

import (
	"time"

	"gorm.io/gorm"
)

type Schedules struct {
	gorm.Model
	CreateBy       string
	AgentId        uint
	ScheduleTitle  string
	SendingAt      time.Time
	ScheduleStatus int
}
