package database

type Blog struct {
	Id         uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Createdate string `json:"createdate"`
	Status     int    `json:"status"`
	Class      int    `json:"class"`
	Lang       string `json:"lang"`
}

func (Blog) TableName() string {
	return "blog"
}
