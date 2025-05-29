package database

type Tag struct {
	TagId   int    `json:"tag_id"`
	TagName string `json:"tag_name"`
}

func (Tag) TableName() string {
	return "tag"
}
