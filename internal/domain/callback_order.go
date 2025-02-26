package domain

import "time"

type CallbackOrder struct {
	Id                uint   `gorm:"primaryKey"`
	ChannelId         string `gorm:"index"` // Связываем с Source.ChannelId
	CustomerID        string `json:"customer_id"`
	RepostBotUsername string `json:"repost_bot_username"`
	Webhooks          string `gorm:"type:json"`
	CreatedAt       time.Time          `json:"created_at"`

}
