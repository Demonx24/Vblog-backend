package router

import (
	"github.com/Demonx24/Vblog-backend/api"
	"github.com/gin-gonic/gin"
)

type LinkRouter struct{}

func (u *LinkRouter) InitLinkRouter(Router *gin.RouterGroup) {
	linkRouter := Router.Group("link")
	linkApi := api.ApiGroupApp.LinkApi
	{
		linkRouter.GET("getalllinks", linkApi.GetAllLinks)
	}
}
