package api

import (
	"github.com/Demonx24/Vblog-backend/model/database"
	"github.com/Demonx24/Vblog-backend/model/request"
	"github.com/Demonx24/Vblog-backend/model/response"
	"github.com/gin-gonic/gin"
)

type BlogApi struct{}

func (BlogApi *BlogApi) BloggerID(c *gin.Context) {

	var req request.BloggerId
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage("传参有问题", c)
		return
	}
	b := database.Blog{Id: req.BlogId}
	blog, err := blogService.BloggerId(b)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	response.OkWithData(blog, c)
}
