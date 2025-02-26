package domain

import "time"

// Links    []Link `gorm:"many2many:source_links"`
type SourcePost struct {
	ChannelId      string `json:"channel_id"`
	PostId         string `json:"post_id"`
	IsStory        bool   `json:"is_story"`
	Subscribers    string
	IsEdited       bool `json:"is_edited"`
	IsDeleted      bool `json:"is_deleted"`
	Message        string
	MediaGroupID   string          `json:"media_group_id"`
	MediaURL       string          `gorm:"-"` // MVP skip
	Entities       string          `gorm:"type:json"`
	Buttons        string          `gorm:"type:json"`
	Reactions      string          `gorm:"type:json"`
	PostStatistics []PostStatistic `gorm:"foreignKey:PostId" json:"post_statistics"`
	SendedAt       time.Time       `json:"sended_at"`
	EditedAt       time.Time       `json:"edited_at"`
	CreatedAt       time.Time          `json:"created_at"`

}
