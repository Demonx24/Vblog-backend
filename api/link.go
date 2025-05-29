package api

import (
	"github.com/Demonx24/Vblog-backend/model/response"
	"github.com/gin-gonic/gin"
)

type LinkApi struct{}

func (LinkApi *LinkApi) GetAllLinks(c *gin.Context) {
	links, err := linkService.GetAllLinks()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	response.OkWithData(links, c)
}
