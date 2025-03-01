package domain

import "time"

type LinkType string

const (
	Public         LinkType = "PUBLIC"
	Private        LinkType = "PRIVATE"
	InvalidValue   LinkType = "INVALID_VALUE"
	InvalidRequest LinkType = "INVALID_REQUEST"
	NotFound       LinkType = "NOT_FOUND"
)

type Link struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	LinkValue string    `json:"link_value"`
	LinkType  LinkType  `json:"link_type"`
	IsValid   bool      `json:"is_valid"`
	CreatedAt time.Time `json:"created_at"`
	Sources   []Source  `gorm:"many2many:link_sources;"`
}
