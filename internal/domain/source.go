package domain

import "time"

type Source struct {
	ChannelId string `gorm:"primaryKey" json:"channel_id"`
	// Links            []Link            `gorm:"many2many:link_sources;"`
	SourceUpdates    []SourceUpdate    `gorm:"foreignKey:ChannelId"`
	SourceStatistics []SourceStatistic `gorm:"foreignKey:ChannelId"`
	SourcePosts      []SourcePost      `gorm:"foreignKey:ChannelId"`
	Sessions         []Session         `gorm:"foreignKey:ChannelId"`
	Orders           []CallbackOrder   `gorm:"foreignKey:ChannelId"`
	CreatedAt       time.Time          `json:"created_at"`
}
