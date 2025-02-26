package domain

import "time"

type Session struct {
	Id uint `gorm:"primaryKey" json:"id"`
	ChannelId string    `gorm:"primaryKey" json:"channel_id"`
	SessionId string    `gorm:"primaryKey" json:"session_id"`
	IsActive  bool      `json:"is_active"`
	JoinAt    time.Time `json:"join_at"`
	CreatedAt       time.Time          `json:"created_at"`
}
