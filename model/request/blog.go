package request

type BloggerId struct {
	BlogId uint `form:"id"`
}
type CreateBlog struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Class   int    `json:"class"`
}

type DeleteBlog struct {
	BlogId uint `json:"id"`
}
type UpdateBlog struct {
	BlogId     uint   `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Createdate string `json:"createdate"`
	Class      int    `json:"class"`
}
type BlogList struct {
	BlogId     uint   `json:"id" form:"id"`
	Title      string `json:"title" form:"title"	`
	Content    string `json:"content" form:"content"	`
	Createdate string `json:"createdate" form:"createdate"	`
	Class      int    `json:"class" form:"class"`
	PageInfo
}
