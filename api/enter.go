package api

import "github.com/Demonx24/Vblog-backend/service"

type ApiGroup struct {
	BlogApi
}

var ApiGroupApp = new(ApiGroup)
var blogService = service.ServiceGroupApp.BlogService
