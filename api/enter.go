package api

import "github.com/Demonx24/Vblog-backend/service"

type ApiGroup struct {
	BlogApi
	TagApi
	LinkApi
	ClassifyApi
}

var ApiGroupApp = new(ApiGroup)
var blogService = service.ServiceGroupApp.BlogService
var tagService = service.ServiceGroupApp.TagService
var linkService = service.ServiceGroupApp.LinkService
var classifyService = service.ServiceGroupApp.ClassifyService
