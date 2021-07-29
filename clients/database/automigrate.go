package database

import (
	"log"
	"pub/entities"

	"gorm.io/gorm"
)

func CreateTable(db *gorm.DB, err error) {
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(
		&entities.SystemAgents{},
		&entities.SystemAccounts{},
		&entities.SystemBots{},
		&entities.ConsumersLineAccounts{},
		&entities.Messages{},
		&entities.ConsumersSegments{},
		&entities.MessagesText{},
		&entities.Schedules{},
		&entities.RelaScheduleMessage{},
		&entities.RelaScheduleSegment{},
		&entities.RelaSegmentLineAccount{},
		&entities.ImagemapAction{},
		&entities.LineEmoji{},
		&entities.LineEmojiProduct{},
		&entities.MessageImagemap{},
		&entities.MessageImage{},
	)
}
