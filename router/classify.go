package router

import (
	"github.com/Demonx24/Vblog-backend/api"
	"github.com/gin-gonic/gin"
)

type ClassifyRouter struct{}

func (c *ClassifyRouter) InitClassifyRouter(Router *gin.RouterGroup) {
	classifyRouter := Router.Group("/classify")
	classifyApi := api.ApiGroupApp.ClassifyApi
	{
		classifyRouter.GET("getCateTree", classifyApi.ClassifyList)
	}
}
