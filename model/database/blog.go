package database

type Blog struct {
	Id      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  int    `json:"status"`
	Class   int    `json:"class"`
	Lang    string `json:"lang"`
}

func (Blog) TableName() string {
	return "blog"
}
