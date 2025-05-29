package api

import (
	"github.com/Demonx24/Vblog-backend/model/request"
	"github.com/Demonx24/Vblog-backend/model/response"
	"github.com/gin-gonic/gin"
)

type ClassifyApi struct{}

func (ClassifyApi *ClassifyApi) ClassifyList(c *gin.Context) {
	var pageInfo request.ClassifyList

	list, total, err := classifyService.ClassifyList(pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(response.PageResult{
		List:  list,
		Total: total,
	}, c)

}
