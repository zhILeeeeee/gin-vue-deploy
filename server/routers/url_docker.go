package routers

import (
	"gin-vue-deploy/server/api"
	"github.com/gin-gonic/gin"
)

type DockerRouter struct{}

func (s *DockerRouter) InitDockerRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	dockerRouter := Router.Group("docker")

	dockerApi := api.ApiGroupApp.DockerApi
	{
		dockerRouter.POST("queryImages", dockerApi.SearchDockerImagesApi) // 查询镜像
		dockerRouter.POST("pullImage", dockerApi.PullDockerImageApi)      // 拉取镜像
		dockerRouter.POST("packImages", dockerApi.PackDockerImageApi)     // 拉取镜像
	}

	return dockerRouter
}
