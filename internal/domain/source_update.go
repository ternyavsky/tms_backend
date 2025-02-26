package domain

import "time"

// Links    []Link `gorm:"many2many:source_links"`
type SourceUpdate struct {
	Id uint `gorm:"primaryKey" json:"id"`
	ChannelId  string `json:"channel_id"`
	Title      string `json:"title"`
	About      string `json:"about"`
	AvatarFile uint   `jons:"avatar_file"`
	CreatedAt       time.Time          `json:"created_at"`
}
