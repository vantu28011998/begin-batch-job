package entities

import (
	"gorm.io/gorm"
)

type ConsumersSegments struct {
	gorm.Model
	CreateBy    string
	AgentId     uint
	SegmentName string
	Description string
}

// func (segment *ConsumersSegments) AfterCreate(tx *gorm.DB) (err error) {
// 	client := redis.GetRedisConnection()
// 	client.SAdd(strconv.Itoa(int(segment.ID)))
// 	return
// }
