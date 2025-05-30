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
		blogRouter.POST("createblog", blogApi.CreateBlog)
		blogRouter.DELETE("deleteblog", blogApi.DeleteBlog)
		blogRouter.PUT("updateblog", blogApi.UpdateBlog)
		blogRouter.GET("blogByPage", blogApi.BlogByPage)
		blogRouter.GET("getCard1", blogApi.GetCard1)
		blogRouter.GET("getCard4", blogApi.GetCard4)
		blogRouter.GET("getArchive", blogApi.GetArchive)
		blogRouter.GET("getIndex", blogApi.GetIndex)
	}
}
