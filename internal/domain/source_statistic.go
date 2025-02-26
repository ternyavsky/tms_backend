package domain

import "time"

// Links    []Link `gorm:"many2many:source_links"`
type SourceStatistic struct {
	Id uint `gorm:"primaryKey" json:"id"`
	ChannelId   string ` json:"channel_id"`
	Subscribers uint
	CreatedAt       time.Time          `json:"created_at"`
}
