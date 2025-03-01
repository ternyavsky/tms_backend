package domain

import "time"

type CallbackOrder struct {
	Id                uint      `gorm:"primaryKey" json:"id"`
	ChannelId         string    `gorm:"index"` // Связываем с Source.ChannelId
	CustomerID        string    `json:"customer_id"`
	RepostBotUsername string    `json:"repost_bot_username"`
	Webhooks          string    `gorm:"type:json"`
	CreatedAt         time.Time `json:"created_at"`
	SourceId          uint      `gorm:"not null" json:"source_id"`
	Source            Source    `gorm:"foreignKey:SourceId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"source,omitempty"`
}
