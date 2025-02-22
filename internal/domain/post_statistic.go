package domain

type PostStatistic struct {
	PostId    string `gorm:"primaryKey" json:"post_id"`
	Views     int
	Reactions int
}
