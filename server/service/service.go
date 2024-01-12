package service

type ServiceGroup struct {
	DockerService
	HarborService
}

var ServiceGroupApp = new(ServiceGroup)
