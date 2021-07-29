package entities

import "gorm.io/gorm"

type RelaSegmentLineAccount struct {
	gorm.Model
	ConsumerSegmentId     uint
	ConsumerLineAccountId uint
}
