package domain

import "time"

type Source struct {
	Id               uint              `gorm:"primaryKey" json:"id"`
	ChannelId        string            `json:"channel_id"`
	Links            []Link            `gorm:"many2many:link_sources;"`
	SourceUpdates    []SourceUpdate    `gorm:"foreignKey:SourceId"`
	SourceStatistics []SourceStatistic `gorm:"foreignKey:SourceId"`
	SourcePosts      []SourcePost      `gorm:"foreignKey:SourceId"`
	Sessions         []Session         `gorm:"foreignKey:SourceId"`
	Orders           []CallbackOrder   `gorm:"foreignKey:SourceId"`
	CreatedAt        time.Time         `json:"created_at"`
}
