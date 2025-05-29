package router

type RouterGroup struct {
	BlogRouter
	TagRouter
	LinkRouter
	ClassifyRouter
}

var RouterGroupApp = new(RouterGroup)
