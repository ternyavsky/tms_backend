package domain

// Links    []Link `gorm:"many2many:source_links"`
type SourceStatistic struct {
	ChannelId   string `gorm:"primaryKey" json:"channel_id"`
	Subscribers uint
}
