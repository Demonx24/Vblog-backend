package router

import (
	"github.com/Demonx24/Vblog-backend/api"
	"github.com/gin-gonic/gin"
)

type TagRouter struct{}

func (t *TagRouter) InitTagRouter(Router *gin.RouterGroup) {
	tagrouter := Router.Group("tag")
	tagapi := api.ApiGroupApp.TagApi
	{
		tagrouter.GET("gettagId", tagapi.GetTagId)
		tagrouter.POST("addtag", tagapi.AddTag)
		tagrouter.PUT("edittag", tagapi.EditTag)
		tagrouter.DELETE("deletetag", tagapi.DeleteTag)
		tagrouter.GET("getAllTags", tagapi.GetAllTags)
	}

}
