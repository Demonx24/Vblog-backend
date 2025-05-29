package database

type Link struct {
	LinkId int    `json:"link_id" form:"link_id"`
	Link   string `json:"link"`
	Pfp    string `json:"pfp" gorm:"type:mediumtext"`
	Title  string `json:"title"`
	Word   string `json:"word"`
}

func (Link *Link) TableName() string {
	return "link"
}
