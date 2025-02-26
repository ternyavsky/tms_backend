package domain

import "time"

type PostStatistic struct {
	Id uint `gorm:"primaryKey" json:"id"`
	PostId    string `json:"post_id"`
	Views     int
	Reactions int
	CreatedAt       time.Time          `json:"created_at"`
	
}
