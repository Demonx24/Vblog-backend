package service

type ServiceGroup struct {
	BlogService
	TagService
	LinkService
	ClassifyService
	JwtService
}

var ServiceGroupApp = new(ServiceGroup)
