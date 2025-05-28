package initialize

import (
	"github.com/Demonx24/Vblog-backend/global"
	"github.com/Demonx24/Vblog-backend/router"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// 设置gin模式
	gin.SetMode(global.Config.System.Env)
	Router := gin.Default()
	//使用gin会话路由
	var store = cookie.NewStore([]byte(global.Config.System.SessionsSecret))
	Router.Use(sessions.Sessions("session", store))
	// 将指定目录下的文件提供给客户端
	// "uploads" 是URL路径前缀，http.Dir("uploads")是实际文件系统中存储文件的目录
	//Router.StaticFS(global.Config.Upload.Path, http.Dir(global.Config.Upload.Path))
	// 创建路由组

	routerGroup := router.RouterGroupApp

	publicGroup := Router.Group(global.Config.System.RouterPrefix)

	routerGroup.InitBlogRouter(publicGroup)

	return Router
}
