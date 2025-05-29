package service

type ServiceGroup struct {
	BlogService
	TagService
	LinkService
	ClassifyService
}

var ServiceGroupApp = new(ServiceGroup)
