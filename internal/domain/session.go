package domain

import "time"

type Session struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	ChannelId string    `json:"channel_id"`
	SessionId string    `json:"session_id"`
	IsActive  bool      `json:"is_active"`
	JoinAt    time.Time `json:"join_at"`
	CreatedAt time.Time `json:"created_at"`
	SourceId  uint      `gorm:"not null" json:"source_id"`
	Source    Source    `gorm:"foreignKey:SourceId" json:"source,omitempty"`
}
