package request

type BloggerId struct {
	BlogId  uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
