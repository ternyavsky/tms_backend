package domain

// Links    []Link `gorm:"many2many:source_links"`
type SourceUpdate struct {
	ChannelId  string `gorm:"primaryKey" json:"channel_id"`
	Title      string `json:"title"`
	About      string `json:"about"`
	AvatarFile uint   `jons:"avatar_file"`
}
