package main

import (
	"github.com/Demonx24/Vblog-backend/core"
	"github.com/Demonx24/Vblog-backend/global"
	"github.com/Demonx24/Vblog-backend/initialize"
)

func main() {
	global.Config = core.InitConfig()
	global.DB = initialize.InitGorm()
	core.RunServer()
}
