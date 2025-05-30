package request

type TagId struct {
	TagId   int    `form:"id" json:"tag_id"`
	TagName string `json:"tag_name"`
}
type AddTag struct {
	TagId   int    `json:"tag_id" form:"tag_id"`
	TagName string `json:"tagname" form:"tagname" validate:"required"`
}
