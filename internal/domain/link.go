package domain

type LinkType string

const (
	Public         LinkType = "PUBLIC"
	Private        LinkType = "PRIVATE"
	InvalidValue   LinkType = "INVALID_VALUE"
	InvalidRequest LinkType = "INVALID_REQUEST"
	NotFound       LinkType = "NOT_FOUND"
)

type Link struct {
	LinkValue string
	LinkType  LinkType
	// Sources   []Source `gorm:"many2many:link_sources;"`
}
