package entities

import "gorm.io/gorm"

type SystemAccounts struct {
	gorm.Model
	AgentId  uint
	Username string
	Email    string
	Password string
	
}
