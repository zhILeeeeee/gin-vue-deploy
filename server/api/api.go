package api

import "gin-vue-deploy/server/service"

type ApiGroup struct {
	DockerApi
	HarborApi
}

var (
	dockerService = service.ServiceGroupApp.DockerService
	harborService = service.ServiceGroupApp.HarborService
)

var ApiGroupApp = new(ApiGroup)
