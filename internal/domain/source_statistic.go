package domain

import "time"

// Links    []Link `gorm:"many2many:source_links"`
type SourceStatistic struct {
	Id          uint   `gorm:"primaryKey" json:"id"`
	ChannelId   string ` json:"channel_id"`
	Subscribers uint
	CreatedAt   time.Time `json:"created_at"`
	SourceId    uint      `gorm:"not null" json:"source_id"`
	Source      Source    `gorm:"foreignKey:SourceId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"source,omitempty"`
}
