package main

import (
	"github.com/Demonx24/Vblog-backend/core"
	"github.com/Demonx24/Vblog-backend/global"
	"github.com/Demonx24/Vblog-backend/initialize"
)

func main() {
	global.Config = core.InitConfig()
	global.Log = core.InitLogger()
	initialize.OtherInit()
	global.Redis = initialize.ConnectRedis()
	global.DB = initialize.InitGorm()
	core.RunServer()

}
