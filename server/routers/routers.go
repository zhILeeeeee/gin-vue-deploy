package routers

type RouterGroup struct {
	DockerRouter
	HarborRouter
}

var RouterGroupApp = new(RouterGroup)
