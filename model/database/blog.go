package database

import "time"

type Blog struct {
	Id         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Createdate time.Time `json:"createdate"`
	Status     int       `json:"status"`
	Class      int       `json:"class"`
	Lang       string    `json:"lang"`
}

func (Blog) TableName() string {
	return "blog"
}
