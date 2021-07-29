package entities

import "gorm.io/gorm"

type RelaScheduleSegment struct {
	gorm.Model
	ScheduleId uint
	SegmentId  uint
}
