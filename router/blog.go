package router

import (
	"github.com/Demonx24/Vblog-backend/api"
	"github.com/gin-gonic/gin"
)

type BlogRouter struct{}

func (u *BlogRouter) InitBlogRouter(Router *gin.RouterGroup) {
	blogRouter := Router.Group("blog")

	blogApi := api.ApiGroupApp.BlogApi
	{
		blogRouter.GET("bloggerid", blogApi.BloggerID)
	}
}
